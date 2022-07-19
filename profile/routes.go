package profile

import (
	"bot"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Routes struct {
	log     *zap.SugaredLogger
	echo    *echo.Echo
	service *Service
}

func RegisterRoutes(log *zap.SugaredLogger, e *echo.Echo, profileService *Service) *Routes {
	r := &Routes{log: log, echo: e, service: profileService}
	profiles := r.echo.Group("/profiles")
	profiles.GET("", r.getAllProfiles)
	profiles.GET("/:id", r.getProfile)
	profiles.POST("", r.addProfile)
	profiles.POST("/:id", r.updateProfile)
	profiles.DELETE("/:id", r.deleteProfile)
	profiles.DELETE("", r.deleteAllProfiles) // only for development
	return r
}
func validateCC(CreditCard) {

}
func (r *Routes) addProfile(c echo.Context) error {

	ec := &bot.EchoContext{Context: c}
	r.log.Info("Adding new profile")
	var err error
	addProfile := new(Profile)

	if err = ec.Bind(addProfile); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = ec.Validate(addProfile); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if addProfile.CreditCard == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "creditCard field missing")
	}
	if err = addProfile.CreditCard.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = r.service.Add(addProfile)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	addProfile.CreditCard.Number = ""
	addProfile.CreditCard.CVV = ""
	return ec.JSON(http.StatusCreated, addProfile)
}

func (r *Routes) updateProfile(c echo.Context) error {
	var err error
	var id uint64

	ec := &bot.EchoContext{Context: c}
	if id, err = strconv.ParseUint(c.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	r.log.Infof("Updating profile %v", id)
	updatedProfile := new(Profile)
	updatedProfile.ID = bot.PK(id)
	if err = ec.Bind(updatedProfile); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = ec.Validate(updatedProfile); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if updatedProfile.CreditCard != nil {
		if err = updatedProfile.CreditCard.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	err = r.service.Update(updatedProfile)
	if bot.NotFoundError(err) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	updatedProfile.CreditCard.Number = ""
	updatedProfile.CreditCard.CVV = ""
	return ec.JSON(http.StatusOK, updatedProfile)
}

func (r *Routes) getAllProfiles(c echo.Context) error {
	ec := &bot.EchoContext{Context: c}
	deleted := ec.QueryParam("deleted") == "true"

	profiles, err := r.service.GetAll(deleted)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// r.log.Infof("All profiles %v", profiles)
	return ec.JSON(http.StatusOK, profiles)
}

func (r *Routes) getProfile(c echo.Context) error {
	ec := &bot.EchoContext{Context: c}
	var id uint64
	var err error
	if id, err = strconv.ParseUint(c.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var profile *Profile
	profile, err = r.service.Get(bot.PK(id))
	if bot.NotFoundError(err) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ec.JSON(http.StatusOK, profile)
}

func (r *Routes) deleteProfile(c echo.Context) error {
	ec := &bot.EchoContext{Context: c}
	var id uint64
	var err error
	if id, err = strconv.ParseUint(c.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = r.service.Delete(bot.PK(id))
	if bot.NotFoundError(err) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ec.String(http.StatusOK, "deleted profile")
}

func (r *Routes) deleteAllProfiles(c echo.Context) error {
	ec := &bot.EchoContext{Context: c}
	purge := ec.QueryParam("purge") == "true"
	var err error
	if purge {
		err = r.service.PurgeAll()
	} else {
		err = r.service.DeleteAll()
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ec.String(http.StatusOK, "deleted all profiles")
}

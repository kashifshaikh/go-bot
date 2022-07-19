package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Routes struct {
	Echo *echo.Echo
	Log  *zap.SugaredLogger
}

func RegisterRoutes(log *zap.SugaredLogger, e *echo.Echo) *Routes {
	r := &Routes{Log: log, Echo: e}

	user := r.Echo.Group("/user")
	user.GET("", r.getUser)
	user.PUT("", r.createUser)
	user.POST("", r.updateUser)
	user.DELETE("", r.deleteUser)

	return r
}

func (r *Routes) createUser(c echo.Context) error {
	return c.String(http.StatusCreated, "createUser success")
}

func (r *Routes) getUser(c echo.Context) error {
	return c.String(http.StatusOK, "user")
}

func (r *Routes) updateUser(c echo.Context) error {
	return c.String(http.StatusOK, "updateUser success")
}

func (r *Routes) deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "deleteUser success")
}

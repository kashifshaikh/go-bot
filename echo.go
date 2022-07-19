package bot

import (
	"errors"
	"net"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/iancoleman/strcase"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoContext struct {
	echo.Context
}

func (c *EchoContext) JSON(code int, i interface{}) error {
	r := c.Response()
	r.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	enc := json.NewEncoder(r)
	r.Status = code
	return enc.Encode(i)
}

func echoListener(cfg *Config) net.Listener {
	var err error
	var listener net.Listener
	if cfg.Sock != "" {
		if err = os.RemoveAll(cfg.Sock); err != nil {
			panic(err)
		}
		listener, err = net.Listen("unix", cfg.Sock)
		if err != nil {
			panic(err)
		}
	} else if cfg.Port != "" {
		listener, err = net.Listen("tcp", ":"+cfg.Port)
		if err != nil {
			panic(err)
		}
	} else {
		panic(errors.New("PORT or SOCK config not specified"))
	}
	return listener
}

type echoValidator struct {
	validator *validator.Validate
}

func (ev *echoValidator) Validate(i interface{}) error {

	if err := ev.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// DefaultCORSConfig = CORSConfig{
//   Skipper:      DefaultSkipper,
//   AllowOrigins: []string{"*"},
//   AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
// }

// func validateExp(fl validator.FieldLevel) bool {

// 	_, _, err := types.ExpMonthYear(fl.Field().String())
// 	if err != nil {
// 		fmt.Printf("exp: %s", err.Error())
// 		return false
// 	}
// 	return true
// }

func NewValidator() *validator.Validate {
	v := validator.New()
	// v.RegisterValidation("exp", validateExp)
	return v
}
func NewEcho(cfg *Config) *echo.Echo {

	extra.SetNamingStrategy(strcase.ToLowerCamel)
	e := echo.New()
	e.Listener = echoListener(cfg)
	e.Validator = &echoValidator{validator: NewValidator()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	return e
}

func StartEchoServer(e *echo.Echo) {
	s := new(http.Server)
	err := e.StartServer(s)
	if err != nil {
		panic(err)
	}
}

package route

import (
	"fmt"
	"net/http"
	"os"
	"testapi/core/app"
	r "testapi/response"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Option struct {
	LogFile    *os.File
	AppContext *app.Context
}

func NewWithOptions(option Option) *echo.Echo {
	router := echo.New()
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: middleware.DefaultLoggerConfig.Format,
		Output: option.LogFile,
	}))

	router.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, r.Success("Hello, Docker! <3\n"))
	})

	router.GET("/name", func(c echo.Context) error {
		return c.JSON(http.StatusOK, r.Success(
			fmt.Sprintf("hostname"),
		))
	})

	router.GET("/man", func(c echo.Context) error {
		return c.JSON(http.StatusOK, r.Success(
			fmt.Sprintf("/, /man, /healthcheck, /db, /page"),
		))
	})

	router.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, r.Success(time.Now().String()))
	})

	router.GET("/db", func(c echo.Context) error {
		return c.JSON(http.StatusInternalServerError, r.InternalServerError())
	})

	router.GET("/page", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, r.NotFound())
	})

	return router
}

func NewInsecure(i interface{}) *echo.Echo {
	router := echo.New()

	router.GET("/debug", func(c echo.Context) error {
		return c.JSON(http.StatusOK, i)
	})

	return router
}

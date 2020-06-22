package router

import (
	"web-backend-patal/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Pre(middleware.Rewrite(map[string]string{
		"/api/*": "/$1",
	}))

	e.GET("/serviceinfo", handlers.ServiceInfo)

	return e
}

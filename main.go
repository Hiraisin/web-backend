package main

import (
	"os"

	"web-backend-patal/config"

	"web-backend-patal/router"

	"github.com/labstack/echo/middleware"
)

func main() {
	defer config.App.Close()

	e := router.NewRouter()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*", "X-Accept-Charset", "X-Accept", "Content-Type", "Authorization", "Accept", "Origin", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
	}))

	e.Logger.Fatal(e.Start(":" + config.App.Port))
	os.Exit(0)

}

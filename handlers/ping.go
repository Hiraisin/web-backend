package handlers

import (
	"fmt"
	"net/http"
	"time"

	"web-backend-patal/config"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Info main type
type Info struct {
	Time string `json:"time"`
	DB   bool   `json:"database"`
}

var (
	err  error
	info Info
)

// ServiceInfo check service info
func ServiceInfo(c echo.Context) error {
	defer c.Request().Body.Close()

	info.Time = fmt.Sprintf("%v", time.Now().Format("2006-01-02T15:04:05"))
	info.DB = true

	if err = healthcheckDB(); err != nil {
		info.DB = false
	}

	return c.JSON(http.StatusOK, info)
}

func healthcheckDB() (err error) {
	dbconf := config.App.Config.GetStringMap("database")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbconf["username"].(string), dbconf["password"].(string), dbconf["host"].(string), dbconf["port"].(string), dbconf["table"].(string))

	db, err := gorm.Open("mysql", connectionString)
	defer db.Close()
	return err
}

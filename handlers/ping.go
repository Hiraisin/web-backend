package handlers

import (
	"fmt"
	"net/http"
	"time"

	"../config"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Info main type
type Info struct {
	Time  string `json:"time"`
	DB    bool   `json:"database"`
	Redis bool   `json:"redis"`
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
	info.Redis = true

	if err = healthcheckDB(); err != nil {
		info.DB = false
	}
	if err = healthcheckRedis(); err != nil {
		info.Redis = false
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

func healthcheckRedis() (err error) {
	redisconf := config.App.Config.GetStringMap(fmt.Sprintf("redis"))
	Redis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisconf["host"].(string), redisconf["port"].(string)),
		Password: redisconf["password"].(string), // no password set
		DB:       redisconf["db"].(int),          // use default DB
	})

	_, err = Redis.Ping().Result()
	return err
}

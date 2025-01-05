package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/you22fy/go_remix_todo/config"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"hello": "world"})
	})

	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil{
		panic(err)
	}

	dbGorm.Ping()

	e.Logger.Fatal(e.Start(":8080"))
}

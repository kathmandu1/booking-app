package main

import (
	faciltycontroller "dinning/app/http/controllers"
	"dinning/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()
	faciltycontroller.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}

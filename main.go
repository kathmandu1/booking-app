package main

import (
	faciltycontroller "dinning/app/http/controllers"
	"dinning/config"

	"github.com/labstack/echo/v4"
)

// @title Swagger  API Endpoints for Demo App
// @version 1.0
// @description This is a Api Endpoints server for using Swagger with Echo.
// @host localhost:8000
// @BasePath /
func main() {
	e := echo.New()
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()
	// e.GET("/swagger/*", echoSwagger.WrapHandler)
	faciltycontroller.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}

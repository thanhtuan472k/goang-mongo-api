package main

import (
	"golang-mongo-api/config"
	"golang-mongo-api/dao"
	"golang-mongo-api/database"
	"golang-mongo-api/routes"

	"github.com/labstack/echo/v4"
)

func init() {
	config.Init()
	database.Connect()
	dao.InitAdminUser()
}

func main() {
	// envVars := config.GetEnv()
	e := echo.New()
	// Routes
	routes.Route(e)

	// Start server at localhost:1323
	e.Logger.Fatal(e.Start(":1323"))

}

// RESTFUL API

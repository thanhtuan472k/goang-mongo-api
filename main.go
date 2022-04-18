package main

import (
	"github.com/labstack/echo/v4"
	"myapp/config"
	"myapp/dao"
	"myapp/database"
	"myapp/routes"
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

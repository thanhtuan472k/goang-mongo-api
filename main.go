package main

import (
	"myapp/config"
	"myapp/database"

	"github.com/labstack/echo/v4"
)

func init() {
	config.Init()
	database.Connect()
}

func main() {
	// envVars := config.GetEnv()
	e := echo.New()

	// Connect to MongoDB

	// Routes
	// routes.UserRoute(e)

	// Start server at localhost:1323
	e.Logger.Fatal(e.Start(":1323"))
}

// RESTFUL API

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/matheusgcoppi/barber-finance-api/database"
	"github.com/matheusgcoppi/barber-finance-api/middleware"
	"github.com/matheusgcoppi/barber-finance-api/repository"
	"github.com/matheusgcoppi/barber-finance-api/routes"
	"github.com/matheusgcoppi/barber-finance-api/service"
)

func main() {
	e := echo.New()

	db, err := database.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.DbRepository{Store: db}

	server := service.NewAPIServer(db, &userRepository)

	middlewaredb := middleware.NewDatabaseMiddleware(db)

	e.Use(middlewaredb.MiddlewareChain())

	routes.SetupRoutes(e, server, middlewaredb)

	e.Logger.Fatal(e.Start(":8080"))
}

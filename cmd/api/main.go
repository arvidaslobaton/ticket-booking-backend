package main

import (
	"fmt"

	"github.com/arvidaslobaton/ticket-booking-app-v1/config"
	"github.com/arvidaslobaton/ticket-booking-app-v1/db"
	"github.com/arvidaslobaton/ticket-booking-app-v1/handlers"
	"github.com/arvidaslobaton/ticket-booking-app-v1/repositories"
	"github.com/arvidaslobaton/ticket-booking-app-v1/services"
	"github.com/gofiber/fiber/v2"
)


func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName: "Ticket-Booking-App",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")

	handlers.NewAuthHandler(server.Group("/auth"), authService)

	// privateRoutes := server.Use(middleware.AuthProtected(db))

	// Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)
	handlers.NewTicketHandler(server.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
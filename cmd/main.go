package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nnrmps/golang-boilerplate/internal/app/router"
)

func main() {
	app := fiber.New()

	groupApi := app.Group("/api")

	router.HealthRouter(groupApi)

	_ = app.Listen(":8080")
}

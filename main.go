package main

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	
	app.Use(cors.New())


	configs.ConnectDB()

	routes.UserRoute(app)

	app.Listen(":6000")
}
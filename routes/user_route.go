package routes

import (
	"github.com/gofiber/fiber/v2"
	"fiber-mongo-api/controllers"
)


func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)
	app.Get("user/:userId", controllers.GetAUser)
}
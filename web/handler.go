package web

import (
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router) {
	app.Post("/createUser", CreateUser())
	app.Post("/createSession", CreateUser())
}

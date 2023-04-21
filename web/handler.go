package web

import (
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router) {
	app.Post("/createUser", CreateUser())
	app.Post("/createSession", CreateSession())
	app.Post("/approveSession", ApproveSession())

	private := fiber.New()
	private.Use(ApproveSession())
	private.Post("/createLobby", CreateLobby())

	app.Mount("/private", private)
}

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
	// todo: /lobby subrouter
	private.Post("/createLobby", CreateLobby())
	private.Post("/connectToLobby", ConnectToLobby())
	private.Post("/isLobbyMaster", IsLobbyMaster())

	app.Mount("/private", private)
}

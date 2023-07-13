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
	app.Mount("/private", private)

	lobby := fiber.New()
	private.Mount("/lobby", lobby)
	lobby.Post("/create", CreateLobby())
	lobby.Post("/connectPlayer", ConnectToLobby())
	lobby.Post("/isMaster", IsLobbyMaster())

}

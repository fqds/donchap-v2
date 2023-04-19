package web

import (
	"donchap-v2/dto"
	"donchap-v2/service"
	request2 "donchap-v2/web/request"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateLobby() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			request2.CreateLobbyRequest
		}
		req := request{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		log.Println(req)

		lobby := &dto.LobbyDto{
			Name:            req.Name,
			MasterID:        9,
			LobbyParameters: req.LobbyParameters,
		}
		if err := service.CreateLobby(lobby); err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(201).JSON("lobby created")
	}
}

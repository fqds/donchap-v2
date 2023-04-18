package web

import (
	// "donchap-v2/dto"
	// "donchap-v2/service"
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

		// user := &dto.UserDto{
		// 	Name:     req.Name,
		// 	Password: req.Password,
		// }
		// signedToken, err := service.CreateSession(user)
		// if err != nil {
		// 	return c.Status(401).JSON(err.Error())
		// }
		return c.Status(201).JSON("")
	}
}

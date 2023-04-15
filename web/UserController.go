package web

import (
	"java-to-go/entity"
	"java-to-go/service"
	request2 "java-to-go/web/request"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			request2.UserRequest
		}
		req := request{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		log.Println(req)

		user := &entity.User{
			Login:    req.Login,
			Password: req.Password,
		}
		id, err := service.CreateUser(user)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(201).JSON(id)
	}
}

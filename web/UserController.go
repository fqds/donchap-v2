package web

import (
	"java-to-go/dto"
	"java-to-go/service"
	request2 "java-to-go/web/request"

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

		user := &dto.UserDto{
			Name:     req.Name,
			Password: req.Password,
		}
		id, err := service.CreateUser(user)
		if err != nil {
			return c.Status(401).JSON(err.Error())
		}
		return c.Status(201).JSON(id)
	}
}

func CreateSession() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			request2.UserRequest
		}
		req := request{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		user := &dto.UserDto{
			Name:     req.Name,
			Password: req.Password,
		}
		id, err := service.CreateSession(user)
		if err != nil {
			return c.Status(401).JSON(err.Error())
		}
		return c.Status(201).JSON(id)
	}
}

package web

import (
	"donchap-v2/dto"
	"donchap-v2/service"
	request2 "donchap-v2/web/request"
	response2 "donchap-v2/web/response"
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

		user := &dto.UserDto{
			Name:     req.Name,
			Password: req.Password,
		}
		
		if err := service.CreateUser(user); err != nil {
			return c.Status(401).JSON(err.Error())
		}

		signedToken, err := service.CreateSession(user)
		response := &response2.AuthTokenResponse{AuthToken: signedToken}
		if err != nil {
			return c.Status(401).JSON(err.Error())
		}
		return c.Status(201).JSON(response)
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
		log.Println(req)

		user := &dto.UserDto{
			Name:     req.Name,
			Password: req.Password,
		}
		signedToken, err := service.CreateSession(user)
		response := &response2.AuthTokenResponse{AuthToken: signedToken}
		if err != nil {
			return c.Status(401).JSON(err.Error())
		}
		return c.Status(200).JSON(response)
	}
}

func ApproveSession() fiber.Handler {
	return func(c * fiber.Ctx) error {
		type request struct {
			request2.AuthTokenRequest
		}
		req := request{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		log.Println(req)
		
		user, err := service.ApproveSession(req.AuthToken)
		if err != nil {
			return c.Status(401).JSON(err.Error())
		}

		return c.Status(201).JSON(user.Name)
	}
}
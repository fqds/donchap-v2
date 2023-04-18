package main

import (
	"donchap-v2/config"
	"donchap-v2/web"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	config.SetConfig()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Cookie",
		AllowMethods:     "GET,POST",
		AllowCredentials: true,
	}))

	api := app.Group("/api")
	web.UserRouter(api)
	log.Fatal(app.Listen(config.Config.GetString("port")))
}

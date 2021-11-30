package main

import (
	"log"

	"photon/controller/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	App := fiber.New()

	App.Use(logger.New())

	authRouter := App.Group("/api")

	{
		authRouter.Post("/signin", auth.Signin)
		authRouter.Post("/signup", auth.Signup)
	}

	if err := App.Listen(":8080"); err != nil {
		log.Println(err.Error())
	}
}

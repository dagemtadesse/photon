package main

import (
	"log"
	"strings"

	"photon/controller/auth"
	"photon/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// sample comment
	App := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			if strings.HasPrefix(ctx.BaseURL(), "/api") {
				return ctx.Status(code).JSON(utils.ErrInternalServerError)
			}

			return ctx.SendStatus(code)
		},
	})

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

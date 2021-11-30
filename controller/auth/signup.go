package auth

import (
	"log"
	"photon/database/queries"
	"photon/model"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Signup(ctx *fiber.Ctx) error {
	// get user credentials from request object
	var newUser model.Credential

	if err := ctx.BodyParser(&newUser); err != nil {
		log.Println(err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	// validate new user credential data
	validate := validator.New()
	err := validate.Struct(newUser)

	if err != nil {
		validationErr := err.(validator.ValidationErrors)
		log.Println(validationErr)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	// hash user password
	passwordSlice := []byte(newUser.Password)
	hash, err := bcrypt.GenerateFromPassword(passwordSlice, bcrypt.DefaultCost)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	// change the plain password by the hashed password
	newUser.Password = string(hash)

	// insert user to the database
	userId, err := queries.CreateUserCreds(newUser)
	if err != nil {
		log.Println(err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	// set the user id the uuid generated from the database
	newUser.Id = userId

	// return the id and email
	return ctx.JSON(fiber.Map{
		"id":    newUser.Id,
		"email": newUser.Email,
	})
}

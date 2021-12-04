package auth

import (
	"net/http"
	"photon/database/queries"
	"photon/model"
	"photon/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

var Validate = validator.New()

func Signup(ctx *fiber.Ctx) error {
	// user credential struct with new user id
	var newUser = model.Credential{
		Id: uuid.New(),
	}
	// parse user cred data into newuser
	// support xml, json and urlencoded
	if err := ctx.BodyParser(&newUser); err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(&utils.ErrInvalidDataFormat)
	}

	// validate new user credential data
	err := Validate.Struct(newUser)

	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(&utils.ErrInvalidData)
	}

	err = newUser.HashPassword()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// insert user to the database
	err = queries.CreateUserCreds(&newUser)

	if err != nil {
		errCode := err.(*pq.Error).Code
		// check if the email already exists in the database
		// and causing constraint violation error
		if errCode == pq.ErrorCode("23505") {
			return ctx.
				Status(http.StatusNotAcceptable).
				JSON(&utils.ErrEmailAlreadyRegistered)
		}

		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// return the id and email
	return ctx.JSON(fiber.Map{"id": newUser.Id})
}

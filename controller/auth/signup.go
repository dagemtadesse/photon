package auth

import (
	"log"
	"net/http"
	"photon/database/queries"
	"photon/model"
	"photon/utils"

	"golang.org/x/crypto/bcrypt"

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
		log.Println(err)

		return ctx.
			Status(http.StatusBadRequest).
			JSON(&utils.ErrInvalidData)
	}

	// hash user password
	passwordSlice := []byte(newUser.Password)
	hash, err := bcrypt.GenerateFromPassword(passwordSlice, bcrypt.DefaultCost)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// change the plain password  on the user struct to
	// the hashed password
	newUser.Password = string(hash)

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

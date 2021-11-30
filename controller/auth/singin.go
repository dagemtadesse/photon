package auth

import (
	"database/sql"
	"log"

	"photon/database/queries"
	"photon/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func logError(ctx *fiber.Ctx, err error) {
	log.Println(ctx.Method(), ctx.OriginalURL(), err.Error())
}

func Signin(ctx *fiber.Ctx) error {

	var reqCred model.Credential

	//parse user credentials from request oject
	if err := ctx.BodyParser(&reqCred); err != nil {
		logError(ctx, err)
		ctx.SendStatus(fiber.StatusBadRequest)
	}

	//fetch user credentials using request creds
	existingCred, err := queries.GetUserCreds(reqCred.Email)

	if err != nil {

		if err == sql.ErrNoRows {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		log.Println(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	//hashed password and plain passwords
	hash := []byte(existingCred.Password)
	plainPassword := []byte(reqCred.Password)

	// compare the hash and the plain passwords
	err = bcrypt.CompareHashAndPassword(hash, plainPassword)

	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		log.Println(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	//create sesssion id
	//create cookie

	return nil
}

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
		log.Println(err)

		if err == sql.ErrNoRows {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	err = reqCred.CompareHashAndPassword(&existingCred)

	if err != nil {
		log.Println(err)
		// the two passwords do not match
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	//create sesssion id
	sessionId := existingCred.CreateSessionId()
	err = queries.StoreSessionId(&existingCred, sessionId)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	//create cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     existingCred.Id.String(),
		Value:    sessionId,
		Secure:   true,
		HTTPOnly: true,
	})

	return ctx.SendStatus(fiber.StatusOK)
}

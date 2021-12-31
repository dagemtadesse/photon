package auth

import (
	"net/http"
	"photon/model"
	"photon/utils"

	"database/sql"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(ctx *gin.Context) {

	var requestUser model.User

	var statusBody *utils.GenericErrorMsg
	statusCode := http.StatusBadRequest

	if err := ctx.BindJSON(&requestUser); err != nil {
		ctx.JSON(statusCode, utils.ErrUnableToParseReqBody)
		return
	}

	// handles valiation, fetching the user from db and comparing the hashed password
	// and creating session
	authErr := authenticateUser(&requestUser, ctx)

	if authErr != nil {
		switch authErr {

		case sql.ErrNoRows: // no user with specifed email exists
			statusBody = utils.ErrUserDoesNotExists
		case bcrypt.ErrMismatchedHashAndPassword: // password mismatch
			statusBody = utils.ErrMismatchedEmailAndPassword
		default:
			panic(authErr)

		}

		ctx.JSON(statusCode, statusBody)
		return
	}

	// login successful
	ctx.Status(http.StatusOK)
}

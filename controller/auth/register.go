package auth

import (
	"net/http"
	"photon/model"
	"photon/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

func RegisterHandler(ctx *gin.Context) {
	var StatusBody *utils.GenericErrorMsg
	var statusCode = http.StatusBadRequest

	var requestUser model.User

	if err := ctx.BindJSON(&requestUser); err != nil {
		ctx.JSON(statusCode, utils.ErrUnableToParseReqBody)
		return
	}

	// handles validating, hashing and storing to database
	regErr := createUser(&requestUser)

	if regErr != nil {
		switch regErr := regErr.(type) {

		case validator.ValidationErrors:
			StatusBody = utils.ErrInvalidDataForFields(regErr)
		case *pq.Error:
			// email already Exists
			if regErr.Code == pq.ErrorCode("23505") {
				StatusBody = utils.ErrEmailAlreadyExists
			}
		default:
			panic(regErr)

		}

		ctx.JSON(statusCode, StatusBody)
		return
	}

	ctx.Status(http.StatusOK)
}

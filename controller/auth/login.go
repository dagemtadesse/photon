package auth

import (
	"photon/model"

	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {

	var requestUser model.User

	if err := ctx.BindJSON(&requestUser); err != nil {
		ctx.String(400, err.Error())
	}

	authErr := authenticateUser(&requestUser)

	if authErr != nil {
		//error handling
		return
	}

	ctx.JSON(200, requestUser)
}

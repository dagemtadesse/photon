package controller

import (
	"photon/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {
	var requestUser model.User

	if err := ctx.BindJSON(&requestUser); err != nil {
		ctx.String(400, err.Error())
	}

	sesssion := sessions.Default(ctx)
	sesssion.Set("email", requestUser.Email)
	sesssion.Set("Password", requestUser.Password)

	sesssion.Save()
	// authErr := requestUser.Authenticate()

	// if authErr != nil {
	// 	//error handling
	// 	return
	// }

	ctx.JSON(200, requestUser)
}

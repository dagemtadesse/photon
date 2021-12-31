package auth

import (
	"photon/database/queries"
	"photon/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func authenticateUser(user *model.User, ctx *gin.Context) (err error) {
	if err = validate.Struct(user); err != nil {
		return err
	}

	existingUser, err := queries.GetUserCreds(user.Email)
	if err != nil {
		return nil
	}

	err = user.CompareHashAndPassword(&existingUser)
	if err != nil {
		return nil
	}

	sesssion := sessions.Default(ctx)
	sesssion.Set("id", existingUser.Id.String())

	return sesssion.Save()

}

func createUser(user *model.User) (err error) {

	if err = validate.Struct(user); err != nil {
		return err
	}

	if err = user.HashPassword(); err != nil {
		return err
	}

	if err = queries.CreateUserCreds(user); err != nil {
		return err
	}

	return
}

func AuthRequired(ctx *gin.Context) {

}

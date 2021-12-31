package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {

	// sessions.
	sesssion := sessions.Default(ctx)
	result := sesssion.Get("id")

	ctx.String(200, result.(string))
}

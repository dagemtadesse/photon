package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LogoutHandler(ctx *gin.Context) {
	// clear the cookie
	sessions.Default(ctx).Clear()

	ctx.Status(http.StatusOK)
}

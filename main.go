package main

import (
	"log"
	"os"

	"photon/controller"
	"photon/controller/auth"
	"photon/database"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var Router *gin.Engine

func main() {

	Router = gin.New()

	Router.Use(sessions.Sessions("_session", database.RedisStore()))

	Router.GET("/test", controller.Test)

	{
		Router.POST("/login", auth.LoginHandler)
		Router.POST("/register", auth.RegisterHandler)
		Router.POST("/logout", auth.LogoutHandler)
	}

	if Router.Run(os.Getenv("PORT")) != nil {
		log.Fatalln("unable to start the server ", os.Getenv("PORT"))
	}
}

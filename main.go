package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var Router *gin.Engine

func main() {

	Router = gin.New()

	Router.GET("/login", func(ctx *gin.Context) {
		ctx.String(200, "login")
	})

	if Router.Run(os.Getenv("PORT")) != nil {
		log.Fatalln("unable to start the server ", os.Getenv("PORT"))
	}
}

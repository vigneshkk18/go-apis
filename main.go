package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vigneshkk18/go-apis/initializers"
	"github.com/vigneshkk18/go-apis/middleware"
	"github.com/vigneshkk18/go-apis/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(middleware.AuthRequired())

	routes.InitializeRoutes(r)
	if err := r.Run(); err != nil {
		log.Panicf("error: %s", err)
	}
}

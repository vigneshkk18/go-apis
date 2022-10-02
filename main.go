package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vigneshkk18/go-apis/initializers"
	"github.com/vigneshkk18/go-apis/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	routes.InitializeRoutes(r)
	http.ListenAndServe(":5000", r)
}

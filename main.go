package main

import (
	"github.com/gabrielmrts/mybooks-golang-api/cmd"
	_ "github.com/gabrielmrts/mybooks-golang-api/docs"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Gin Book Service
// @version         1.0
// @description     A book management service API in Go using Gin framework.

// @host      localhost:8090
// @BasePath  /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Run()

	router := routes.SetupRouter()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	serverPort := ":8090"
	router.Run(serverPort)
}

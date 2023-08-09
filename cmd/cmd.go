package cmd

import (
	"github.com/gabrielmrts/mybooks-golang-api/config"
	"github.com/gabrielmrts/mybooks-golang-api/internal/api/routes"
	"github.com/gabrielmrts/mybooks-golang-api/internal/database"
	"github.com/gabrielmrts/mybooks-golang-api/internal/database/migrations"
)

func Run() {
	config.Init()
	database.Init()
	migrations.Migrate()

	router := routes.SetupRouter()
	serverPort := ":8090"
	router.Run(serverPort)
}

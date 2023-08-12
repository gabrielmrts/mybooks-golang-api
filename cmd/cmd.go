package cmd

import (
	"github.com/gabrielmrts/mybooks-golang-api/config"
	"github.com/gabrielmrts/mybooks-golang-api/internal/database"
	"github.com/gabrielmrts/mybooks-golang-api/internal/database/migrations"
)

func Run() {
	config.Init()
	database.Init()
	migrations.Migrate()

	// conf := config.GetConfig()
	// if conf.ENVIRONMENT == "dev" {
	// 	migrations.Seed()
	// }
}

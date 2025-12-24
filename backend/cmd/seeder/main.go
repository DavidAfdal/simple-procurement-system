package main

import (
	"github.com/DavidAfdal/purchasing-systeam/config"
	"github.com/DavidAfdal/purchasing-systeam/database/migration"
	"github.com/DavidAfdal/purchasing-systeam/database/seeders"
	"github.com/DavidAfdal/purchasing-systeam/pkg/database"
)

func main() {
	cfg, err := config.NewConfig()
	checkError(err)

	db, err := database.InitDB(&cfg.Database)
	checkError(err)

	err = migration.AutoMigrate(db)
	checkError(err)

	seeders.Seedrs(db)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

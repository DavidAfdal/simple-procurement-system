package main

import (
	"time"

	"github.com/DavidAfdal/purchasing-systeam/config"
	"github.com/DavidAfdal/purchasing-systeam/internal/builder"
	"github.com/DavidAfdal/purchasing-systeam/pkg/database"
	"github.com/DavidAfdal/purchasing-systeam/pkg/server"
	"github.com/DavidAfdal/purchasing-systeam/pkg/token"
)

func main() {
	cfg, err := config.NewConfig()
	checkError(err)

	db, err := database.InitDB(&cfg.Database)
	checkError(err)

	err = database.AutoMigrate(db)
	checkError(err)

	tokenUse := token.NewTokenUseCase(cfg.JWT.SecretKey, time.Duration(cfg.JWT.ExpiresAt)*time.Hour)

	publicRoutes := builder.BuildAppPublicRoutes(db, tokenUse)
	privateRoutes := builder.BuildAppPrivateRoutes(db, tokenUse)

	srv := server.NewServer(publicRoutes, privateRoutes, cfg.JWT.SecretKey, tokenUse)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

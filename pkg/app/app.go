package app

import (
	server "MEDODS"
	"MEDODS/configs"
	"MEDODS/pkg/handler"
	"MEDODS/pkg/postgres"
	"MEDODS/pkg/repository"
	"MEDODS/pkg/service"
	_ "github.com/lib/pq"
	"log"
)

func Run(cfg *configs.Config) {
	logger := log.Default()
	db, err := postgres.New(cfg.DBName, cfg.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.Driver, cfg.DB.SSLMode)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	repos := repository.NewRepository(db)
	services := service.NewService(repos, cfg.JwtSecret)
	handlers := handler.NewHandler(services, logger)
	srv := new(server.Server)
	logger.Println("Starting server...")
	if err = srv.Run(cfg.HTTP.Port, handlers.InitRoutes()); err != nil {
		logger.Println(err)
	}
}

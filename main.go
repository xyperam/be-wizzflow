package main

import (
	"log"
	"net/http"

	"github.com/xyperam/wizzflow/internal/config"
	"github.com/xyperam/wizzflow/internal/handler"
	"github.com/xyperam/wizzflow/internal/repository"
	"github.com/xyperam/wizzflow/internal/routes"
	"github.com/xyperam/wizzflow/internal/service"
)

func main() {
	//inisialisasi db
	// inject depedency
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	hdl := handler.NewHandler(svc)

	router := routes.SetupRoutes(hdl)

	// load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Println("Starting server on port " + cfg.Port)

	err = http.ListenAndServe(":"+cfg.Port, router)
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

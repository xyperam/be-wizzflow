package main

import (
	"log"
	"net/http"

	"github.com/xyperam/wizzflow/internal/config"
	"github.com/xyperam/wizzflow/internal/database"
	"github.com/xyperam/wizzflow/internal/handler"
	"github.com/xyperam/wizzflow/internal/repository"
	"github.com/xyperam/wizzflow/internal/routes"
	"github.com/xyperam/wizzflow/internal/service"
)

func main() {
	// load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	//inisialisasi db
	dbPool, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Gagal connect to DB: %v", err)
	}
	defer dbPool.Close()
	// inject depedency
	authRepo := repository.NewUserRepository(dbPool)
	authSvc := service.NewAuthService(authRepo, cfg)
	authHdl := handler.NewAuthHandler(authSvc)

	txnRepo := repository.NewRepository(dbPool)
	txnSvc := service.NewTransactionService(txnRepo)
	txnHdl := handler.NewTransactionHandler(txnSvc) // Pastikan nama fungsinya sesuai

	router := routes.SetupRoutes(txnHdl, authHdl, cfg)

	log.Println("Starting server on port " + cfg.Port)

	err = http.ListenAndServe(":"+cfg.Port, router)
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

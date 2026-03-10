package service

import "github.com/xyperam/wizzflow/internal/config"

type authService struct {
	repo repository.UserRepository
	cfg  *config.Config
}

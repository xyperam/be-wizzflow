package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xyperam/wizzflow/internal/config"
	"github.com/xyperam/wizzflow/internal/models"
	"github.com/xyperam/wizzflow/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, req models.AuthRequest) (models.User, error)
	Login(ctx context.Context, req models.AuthRequest) (string, error)
}
type authService struct {
	repo repository.UserRepository
	cfg  *config.Config
}

func NewAuthService(repo repository.UserRepository, cfg *config.Config) AuthService {
	return &authService{repo: repo, cfg: cfg}
}

func (s *authService) Register(ctx context.Context, req models.AuthRequest) (models.User, error) {
	//hash password

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	//save user ke database
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashed),
	}
	return s.repo.SaveUser(ctx, user)
}

func (s *authService) Login(ctx context.Context, req models.AuthRequest) (string, error) {
	// find user
	user, err := s.repo.FindUserByUsername(ctx, req.Username)
	if err != nil {
		return "", errors.New("Invalid username")
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	// Generate token dan claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	//sign dengan secret key
	return token.SignedString([]byte(s.cfg.JWTSecret))
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xyperam/wizzflow/internal/models"
	"github.com/xyperam/wizzflow/internal/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{authService: s}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.AuthRequest

	// bind json

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format Data Salah"})
		return
	}

	// call service
	user, err := h.authService.Register(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal registerasi user"})
		return
	}
	// return user
	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration Success",
		"user":    user,
	})
}
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username atau password wajib diisi"})
		return
	}
	token, err := h.authService.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

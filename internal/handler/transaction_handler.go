package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xyperam/wizzflow/internal/models"
	"github.com/xyperam/wizzflow/internal/service"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(s service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {

	//get User id
	userID := c.MustGet("user_id").(int)

	//get data froms service
	transactions, err := h.service.GetAllTransaction(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Ambil Data"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) SaveTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Invalid"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transaksi berhasil disimpan"})
}

func (h *TransactionHandler) UpdateTranscation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Invalid"})
		return
	}

	userID := c.MustGet("user_id").(int)
	updated, err := h.service.UpdateTransaction(c.Request.Context(), id, userID, transaction)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}
func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	userID := c.MustGet("user_id").(int)

	err := h.service.DeleteTransaction(c.Request.Context(), id, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan atau bukan milikmu"})
		return
	}

	c.Status(http.StatusNoContent)
}

// 5. Get Summary
func (h *TransactionHandler) GetSummary(c *gin.Context) {
	userID := c.MustGet("user_id").(int)

	summary, err := h.service.GetSummary(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, summary)
}

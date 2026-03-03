package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/xyperam/wizzflow/internal/models"
	"github.com/xyperam/wizzflow/internal/service"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

func (h *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	//get data froms service
	transactions := h.service.GetAllTransaction()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	createTransaction := h.service.CreateTransaction(transaction)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createTransaction)

}

func (h *TransactionHandler) UpdateTranscation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var transaction models.Transaction

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	updatedtransaction, err := h.service.UpdateTransaction(id, transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// kirim response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedtransaction)
}

// delete shandler
func (h *TransactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	err := h.service.DeleteTransaction(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
func (h *TransactionHandler) GetSummary(w http.ResponseWriter, r *http.Request) {
	transactions := h.service.GetSummary()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)

}

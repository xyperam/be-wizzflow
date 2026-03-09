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
	transactions, err := h.service.GetAllTransaction(r.Context())
	if err != nil {
		http.Error(w, "Gagal Ambil Data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func (h *TransactionHandler) SaveTransaction(w http.ResponseWriter, r *http.Request) {

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

	result, err := h.service.SaveTransaction(r.Context(), transaction)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

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

	updatedtransaction, err := h.service.UpdateTransaction(r.Context(), id, transaction)
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

	err := h.service.DeleteTransaction(r.Context(), id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
func (h *TransactionHandler) GetSummary(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.service.GetSummary(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)

}

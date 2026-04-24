package service

import (
	"context"
	"errors"

	"github.com/xyperam/wizzflow/internal/models"
	"github.com/xyperam/wizzflow/internal/repository"
)

type TransactionService interface {
	GetAllTransaction(ctx context.Context, userID int) ([]models.Transaction, error)
	SaveTransaction(ctx context.Context, t models.Transaction) (models.Transaction, error)
	UpdateTransaction(ctx context.Context, id int, userID int, t models.Transaction) (models.Transaction, error)
	DeleteTransaction(ctx context.Context, id int, userID int) error
	GetSummary(ctx context.Context, userID int) (models.Summary, error)
}

// Gunakan nama kecil (unexported) agar tidak bentrok dengan interface
type transactionService struct {
	repo repository.TransactionRepository
}

// Return interface-nya, bukan struct-nya
func NewTransactionService(r repository.TransactionRepository) TransactionService {
	return &transactionService{repo: r}
}

// 1. Get All - Pastikan nerima userID
func (s *transactionService) GetAllTransaction(ctx context.Context, userID int) ([]models.Transaction, error) {
	return s.repo.FindAll(ctx, userID) // Repo juga harus nerima userID ya bray!
}

// 2. Save
func (s *transactionService) SaveTransaction(ctx context.Context, t models.Transaction) (models.Transaction, error) {
	if t.Amount <= 0 {
		return models.Transaction{}, errors.New("nominal tidak boleh nol atau minus")
	}
	if t.Title == "" {
		return models.Transaction{}, errors.New("title tidak boleh kosong")
	}
	if t.Type != "income" && t.Type != "expense" {
		return models.Transaction{}, errors.New("tipe transaksi harus income atau expense")
	}

	return s.repo.SaveTransaction(ctx, t)
}

// 3. Update - Cek kepemilikan data
func (s *transactionService) UpdateTransaction(ctx context.Context, id int, userID int, t models.Transaction) (models.Transaction, error) {
	// Cari data aslinya dulu
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return models.Transaction{}, errors.New("data tidak ditemukan")
	}

	// SECURITY CHECK: Pastikan user_id di DB sama dengan user_id yang lagi login
	if existing.UserID != userID {
		return models.Transaction{}, errors.New("kamu tidak punya akses ke data ini")
	}

	return s.repo.UpdateTransaction(ctx, id, t)
}

// 4. Delete - Cek kepemilikan data
func (s *transactionService) DeleteTransaction(ctx context.Context, id int, userID int) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("data tidak ditemukan")
	}

	if existing.UserID != userID {
		return errors.New("kamu tidak punya akses untuk menghapus data ini")
	}

	return s.repo.DeleteTransaction(ctx, id)
}

// 5. Summary - Filter per User
func (s *transactionService) GetSummary(ctx context.Context, userID int) (models.Summary, error) {
	transactions, err := s.repo.FindAll(ctx, userID)
	if err != nil {
		return models.Summary{}, err
	}

	var totalIncome, totalExpense float64
	for _, t := range transactions {
		switch t.Type {
		case "income":
			totalIncome += t.Amount
		case "expense":
			totalExpense += t.Amount
		}
	}

	return models.Summary{
		TotalIncome:  totalIncome,
		TotalExpense: totalExpense,
		Balance:      totalIncome - totalExpense,
	}, nil
}

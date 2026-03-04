package service

import (
	"context"
	"errors"

	"github.com/xyperam/wizzflow/internal/models"
	"github.com/xyperam/wizzflow/internal/repository"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewService(r repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: r}

}

func (s *TransactionService) GetAllTransaction(ctx context.Context) ([]models.Transaction, error) {
	return s.repo.FindAll(ctx)
}

func (s *TransactionService) SaveTransaction(ctx context.Context, t models.Transaction) (models.Transaction, error) {
	if t.Amount <= 0 {
		return models.Transaction{}, errors.New("Nominal tidak boleh nol atau minus")
	}
	if t.Title == "" {
		return models.Transaction{}, errors.New("Title tidak boleh kosong")
	}
	return s.repo.SaveTransaction(ctx, t)
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, id int, t models.Transaction) (models.Transaction, error) {

	_, err := s.repo.FindByID(ctx, id)

	if err != nil {
		return models.Transaction{}, errors.New("data tidak ditemukan")
	}
	return s.repo.UpdateTransaction(ctx, id, t)
}

// service delete
func (s *TransactionService) DeleteTransaction(ctx context.Context, id int) error {
	return s.repo.DeleteTransaction(ctx, id)
}

func (s *TransactionService) GetSummary(ctx context.Context) (models.Summary, error) {
	transactions, err := s.repo.FindAll(ctx)

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

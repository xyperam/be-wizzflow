package service

import (
	"github.com/xyperam/wizzflow/internal/models"
	"github.com/xyperam/wizzflow/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewService(r *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: r}

}

func (s *TransactionService) GetAllTransaction() []models.Transaction {
	return s.repo.FindAll()
}

func (s *TransactionService) CreateTransaction(t models.Transaction) models.Transaction {
	createdTransaction := s.repo.SaveTransaction(t)
	return createdTransaction
}

func (s *TransactionService) UpdateTransaction(id int, t models.Transaction) (models.Transaction, error) {

	transaction, err := s.repo.UpdateTransaction(id, t)

	if err != nil {
		return models.Transaction{}, err
	}
	return transaction, nil
}

// service delete
func (s *TransactionService) DeleteTransaction(id int) error {
	err := s.repo.DeleteTransaction(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TransactionService) GetSummary() models.Summary {
	transactions := s.repo.FindAll()

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
	}
}

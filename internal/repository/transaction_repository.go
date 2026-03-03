package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/xyperam/wizzflow/internal/models"
)

type TransactionRepository struct {
	nextID       int
	transactions []models.Transaction
}

func NewRepository() *TransactionRepository {
	return &TransactionRepository{
		nextID: 2,
		transactions: []models.Transaction{
			{ID: 1, Title: "tabungan", Amount: 2000, Type: "income", Category: "Salary", CreatedAt: time.Now()},
		},
	}
}

// FindAll

func (r *TransactionRepository) FindAll() []models.Transaction {
	return r.transactions

}

// Save

func (r *TransactionRepository) SaveTransaction(transaction models.Transaction) models.Transaction {

	transaction.ID = r.nextID
	transaction.CreatedAt = time.Now()
	r.transactions = append(r.transactions, transaction)

	r.nextID++
	return transaction

}

// FindByID
func (r *TransactionRepository) FindByID(id int) (*models.Transaction, error) {
	for i := range r.transactions {
		if r.transactions[i].ID == id {
			// return pointer
			return &r.transactions[i], nil
		}
	}
	return nil, errors.New("transaction not found")
}

// Update
func (r *TransactionRepository) UpdateTransaction(id int, updateTransaction models.Transaction) (models.Transaction, error) {
	for i, transaction := range r.transactions {
		if transaction.ID == id {
			updateTransaction.ID = id
			r.transactions[i] = updateTransaction
			return r.transactions[i], nil
		}
	}
	return models.Transaction{}, fmt.Errorf("transaction with id %d not found", id)
}

// Delete
func (r *TransactionRepository) DeleteTransaction(id int) error {
	for i, transaction := range r.transactions {
		if transaction.ID == id {
			r.transactions = append(r.transactions[:i], r.transactions[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("transaction with id %d not found", id)
}

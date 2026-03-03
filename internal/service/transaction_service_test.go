package service

import (
	"testing"

	"github.com/xyperam/wizzflow/internal/models"
	"github.com/xyperam/wizzflow/internal/repository"
)

func TestGetSummary(t *testing.T) {
	repo := repository.NewRepository()
	svc := NewService(repo)

	// Kita asumsikan ada data awal 2000 (Income)
	// Mari kita tambahkan data secara spesifik
	testCases := []struct {
		name   string
		title  string
		amount float64
		tType  string // "income" atau "expense"
	}{
		{"Tambah Income", "Project A", 5000, "income"},
		{"Tambah Expense", "Sewa Kantor", 1500, "expense"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			svc.CreateTransaction(models.Transaction{
				Title:  tc.title,
				Amount: tc.amount,
				Type:   tc.tType,
			})
		})
	}

	summary := svc.GetSummary()

	// Kalkulasi ekspektasi:
	// Income: 2000 (awal) + 5000 = 7000
	// Expense: 1500
	// Balance: 5500
	expectedIncome := 7000.0
	// expectedExpense := 1500.0
	// expectedBalance := 5500.0

	if summary.TotalIncome != expectedIncome {
		t.Errorf("%s: Income salah, mau %.2f dapet %.2f", t.Name(), expectedIncome, summary.TotalIncome)
	}
}

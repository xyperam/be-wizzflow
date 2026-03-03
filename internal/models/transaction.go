package models

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

type Summary struct {
	TotalIncome  float64 `json:"total_income"`
	TotalExpense float64 `json:"total_expense"`
	Balance      float64 `json:"balance"`
}

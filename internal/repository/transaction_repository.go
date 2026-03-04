package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xyperam/wizzflow/internal/models"
)

type TransactionRepository interface {
	FindAll(ctx context.Context) ([]models.Transaction, error)
	FindByID(ctx context.Context, id int) (models.Transaction, error)
	SaveTransaction(ctx context.Context, t models.Transaction) (models.Transaction, error)
	UpdateTransaction(ctx context.Context, id int, t models.Transaction) (models.Transaction, error)
	DeleteTransaction(ctx context.Context, id int) error
}

type postgresRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) TransactionRepository {
	return &postgresRepository{db: db}
}

// FindAll

func (r *postgresRepository) FindAll(ctx context.Context) ([]models.Transaction, error) {
	query := `SELECT id,title,amount,type,category, created_at FROM transactions`
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(&t.ID, &t.Title, &t.Amount, &t.Type, &t.Category, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil

}

// Save

func (r *postgresRepository) SaveTransaction(ctx context.Context, t models.Transaction) (models.Transaction, error) {
	query := `
	INSERT INTO transactions (title,amount,type,category)
	VALUES ($1,$2,$3,$4)
	RETURNING id,created_at`

	err := r.db.QueryRow(ctx, query,
		t.Title, t.Amount, t.Type, t.Category,
	).Scan(&t.ID, &t.CreatedAt)

	if err != nil {
		return models.Transaction{}, err
	}
	return t, nil

}

// FindByID
func (r *postgresRepository) FindByID(ctx context.Context, id int) (models.Transaction, error) {
	var t models.Transaction
	query := `SELECT id,title,amount,type,category, created_at FROM transactions WHERE id =$1 `

	err := r.db.QueryRow(ctx, query, id).Scan(&t.ID, &t.Title, &t.Amount, &t.Type, &t.Category, &t.CreatedAt)

	if err != nil {
		return models.Transaction{}, err
	}
	return t, nil

}

// Update
func (r *postgresRepository) UpdateTransaction(ctx context.Context, id int, t models.Transaction) (models.Transaction, error) {

	query := `
	UPDATE transactions
	SET title = $1, amount = $2, type =$3, category = $4 WHERE id =$5
	RETURNING id, title, amount, type, category, created_at`

	err := r.db.QueryRow(ctx, query,
		t.Title, t.Amount, t.Type, t.Category, id,
	).Scan(&t.ID, &t.Title, &t.Amount, &t.Type, &t.Category, &t.CreatedAt)

	if err != nil {
		return models.Transaction{}, err
	}
	return t, nil

}

// Delete
func (r *postgresRepository) DeleteTransaction(ctx context.Context, id int) error {
	query := `DELETE FROM transactions WHERE id=$1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("transaction with id %d not found", id)
	}
	return nil
}

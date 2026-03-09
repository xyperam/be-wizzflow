package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xyperam/wizzflow/internal/models"
)

type MockRepository struct {
	mock.Mock
}

// DeleteTransaction implements repository.TransactionRepository.
func (m *MockRepository) DeleteTransaction(ctx context.Context, id int) error {
	panic("unimplemented")
}

// FindAll implements repository.TransactionRepository.
func (m *MockRepository) FindAll(ctx context.Context) ([]models.Transaction, error) {
	panic("unimplemented")
}

// FindByID implements repository.TransactionRepository.
func (m *MockRepository) FindByID(ctx context.Context, id int) (models.Transaction, error) {
	panic("unimplemented")
}

// UpdateTransaction implements repository.TransactionRepository.
func (m *MockRepository) UpdateTransaction(ctx context.Context, id int, t models.Transaction) (models.Transaction, error) {
	panic("unimplemented")
}

func (m *MockRepository) SaveTransaction(ctx context.Context, t models.Transaction) (models.Transaction, error) {
	args := m.Called(ctx, t)
	return args.Get(0).(models.Transaction), args.Error(1)
}

func TestSaveTransaction_Validation(t *testing.T) {
	repo := new(MockRepository)
	svc := NewTransactionService(repo) // inject mock repo ke service

	t.Run("Return Error if nominal minus", func(t *testing.T) {
		input := models.Transaction{
			Title:  "Beli Kopi",
			Amount: -5000,
			Type:   "expense",
		}
		result, err := svc.SaveTransaction(context.Background(), input)
		assert.Error(t, err)
		assert.Equal(t, "Nominal tidak boleh nol atau minus", err.Error())
		assert.Empty(t, result)

		repo.AssertNotCalled(t, "SaveTransaction", mock.Anything, mock.Anything)
	})
	t.Run("Succes Save Transaction", func(t *testing.T) {
		input := models.Transaction{
			Title:  "Gajian",
			Amount: 1000000,
			Type:   "income",
		}
		// expected balikin data yang sama
		repo.On("SaveTransaction", mock.Anything, input).Return(input, nil)
		result, err := svc.SaveTransaction(context.Background(), input)

		assert.NoError(t, err)
		assert.Equal(t, input.Amount, result.Amount)
		assert.Equal(t, "Gajian", result.Title)

		//verif jika method di repo beneran dipanggil
		repo.AssertExpectations(t)
	})
}

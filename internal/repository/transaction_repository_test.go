package repository

import "testing"

func TestDeleteTransaction(t *testing.T) {
	repo := NewRepository()

	err := repo.DeleteTransaction(1)
	if err != nil {
		t.Errorf("Harusnya suskse hapus ID 1, tapi dapat error %v", err)
	}

	transactions := repo.FindAll()
	if len(transactions) != 0 {
		t.Errorf("Harusnya slice kosong tapi sisa %d", len(transactions))
	}

	err = repo.DeleteTransaction(1)
	if err == nil {
		t.Errorf("Harusnya id 1 sudah dihapus tapi malah nil")
	}
}

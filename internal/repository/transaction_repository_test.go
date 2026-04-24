// package repository

// import (
// 	"context"
// 	"testing"

// 	"github.com/jackc/pgx/v5/pgxpool"
// 	"github.com/xyperam/wizzflow/internal/models"
// )

// func TestDeleteTransaction(t *testing.T) {
// 	// Gunakan DSN yang sama dengan aplikasi jika belum buat db khusus test
// 	ctx := context.Background()
// 	dsn := "postgres://user:password@localhost:5432/wizzflow"
// 	pool, err := pgxpool.New(ctx, dsn)
// 	if err != nil {
// 		t.Fatalf("failed to create pool: %v", err)
// 	}
// 	defer pool.Close()

// 	tx, err := pool.Begin(ctx)
// 	if err != nil {
// 		t.Fatalf("Failed to begin transaction %v", err)
// 	}
// 	defer tx.Rollback(ctx)
// 	repo := NewRepository(pool)

// 	// Langkah 1: Insert data dulu buat bahan hapusan
// 	dummy := models.Transaction{Title: "Test", Amount: 100, Type: "expense", Category: "test"}
// 	created, err := repo.SaveTransaction(dummy)
// 	if err != nil {
// 		t.Fatalf("Gagal bikin data dummy: %v", err)
// 	}

// 	// Langkah 2: Test Hapus yang beneran ada
// 	err = repo.DeleteTransaction(created.ID) // Gunakan ID yang baru dibuat
// 	if err != nil {
// 		t.Errorf("Harusnya sukses hapus ID %d, tapi dapat error: %v", created.ID, err)
// 	}

// 	// Langkah 3: Cek apakah benar-benar hilang
// 	transactions, err := repo.FindAll()
// 	if err != nil {
// 		t.Fatalf("Gagal FindAll: %v", err)
// 	}

// 	// Pastikan ID yang tadi dihapus sudah tidak ada di list
// 	for _, item := range transactions {
// 		if item.ID == created.ID {
// 			t.Errorf("ID %d masih ada di database!", created.ID)
// 		}
// 	}

// 	// Langkah 4: Test Hapus ID yang TIDAK ADA (Expect Error)
// 	err = repo.DeleteTransaction(999999)
// 	if err == nil {
// 		t.Error("Harusnya error karena ID 999999 tidak ada, tapi malah nil")
// 	}
// }

// func TestFindAll(t *testing.T) {
// 	ctx := context.Background()
// 	dsn := "postgres://user:password@localhost:5432/wizzflow"
// 	pool, err := pgxpool.New(ctx, dsn)
// 	if err != nil {
// 		t.Fatalf("failed to create pool: %v", err)
// 	}
// 	defer pool.Close()
// 	repo := NewRepository(pool)

// 	dummy := models.Transaction{Title: "Test", Amount: 100, Type: "expense", Category: "test"}
// 	_, err = repo.SaveTransaction(dummy)
// 	if err != nil {
// 		t.Fatalf("Gagal Membuat data,%v", err)
// 	}

// 	transactions, err := repo.FindAll()
// 	if err != nil {
// 		t.Fatalf("gagal find All %v", err)
// 	}

// 	if len(transactions) == 0 {
// 		t.Errorf("Harusnya data tidak kosong")
// 	}
// }

package routes

import (
	"net/http"

	"github.com/xyperam/wizzflow/internal/handler"
)

// SetupRoutes menerima *gin.Engine
func SetupRoutes(hdl *handler.TransactionHandler) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/transactions", hdl.GetTransactions)
	mux.HandleFunc("/transactions/create", hdl.SaveTransaction)
	mux.HandleFunc("/transactions/summary", hdl.GetSummary)
	mux.HandleFunc("/transactions/update", hdl.UpdateTranscation)
	mux.HandleFunc("/transactions/delete", hdl.DeleteTransaction)
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return mux
}

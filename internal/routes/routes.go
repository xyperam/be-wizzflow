package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xyperam/wizzflow/internal/config"
	"github.com/xyperam/wizzflow/internal/handler"
	"github.com/xyperam/wizzflow/internal/middleware"
)

// SetupRoutes menerima *gin.Engine
// func SetupRoutes(hdl *handler.TransactionHandler) *http.ServeMux {

// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/transactions", hdl.GetTransactions)
// 	mux.HandleFunc("/transactions/create", hdl.SaveTransaction)
// 	mux.HandleFunc("/transactions/summary", hdl.GetSummary)
// 	mux.HandleFunc("/transactions/update", hdl.UpdateTranscation)
// 	mux.HandleFunc("/transactions/delete", hdl.DeleteTransaction)
// 	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("pong"))
// 	})

//		return mux
//	}
func SetupRoutes(
	txHdl *handler.TransactionHandler,
	authHdl *handler.AuthHandler,
	cfg *config.Config,

) *gin.Engine {
	r := gin.Default()

	// Ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ping"})
	})
	r.POST("register", authHdl.Register)
	r.POST("login", authHdl.Login)

	// group routes
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		protected.GET("/transactions", txHdl.GetTransactions)
		protected.POST("/transactions", txHdl.SaveTransaction)
		protected.GET("/transactions/summary", txHdl.GetSummary)
		protected.PUT("/transactions/:id", txHdl.UpdateTranscation)
		protected.DELETE("/transactions/:id", txHdl.DeleteTransaction)
	}
	return r
}

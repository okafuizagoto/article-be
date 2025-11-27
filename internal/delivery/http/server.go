package http

import (
	"net/http"

	"article-be/pkg/grace"

	"github.com/rs/cors"
)

// GoldGymHandler ...
type GoldGymHandler interface {
	GetArticle(w http.ResponseWriter, r *http.Request)
	InsertArticle(w http.ResponseWriter, r *http.Request)
	DeleteArticle(w http.ResponseWriter, r *http.Request)
	UpdateArticle(w http.ResponseWriter, r *http.Request)
	// PrintSelisih(w http.ResponseWriter, r *http.Request)
	// PrintExpiredTerpajang(w http.ResponseWriter, r *http.Request)
	// PrintExpiredTerkumpul(w http.ResponseWriter, r *http.Request)

	// PrintBatch(w http.ResponseWriter, r *http.Request)
	// PrintBatchFull(w http.ResponseWriter, r *http.Request)

	// //Trans Out
	// InsertTransOut(w http.ResponseWriter, r *http.Request)
	// InsertSales(w http.ResponseWriter, r *http.Request)
	// DeleteSalesByPeriod(w http.ResponseWriter, r *http.Request)
	// RemoveSalesByOutcode(w http.ResponseWriter, r *http.Request)
	// InsertBatchData(w http.ResponseWriter, r *http.Request)
}

// AuthHandler ...
type AuthHandler interface {
	LoginUser(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	Goldgym GoldGymHandler
	// Auth    AuthHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}

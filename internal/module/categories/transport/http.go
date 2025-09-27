package transport

import (
	"encoding/json"
	"net/http"

	"github.com/yeungon/gossr/internal/module/categories/business"
	"github.com/yeungon/gossr/internal/module/categories/domain"
)

type Handler struct {
	svc *business.Service
}

func NewHandler(svc *business.Service) http.Handler {
	return &Handler{svc: svc}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var order domain.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.svc.CreateOrder(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

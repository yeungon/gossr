package transport

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yeungon/gossr/config"
	"github.com/yeungon/gossr/internal/module/categories/business"
)

type CategoryHandler struct {
	service *business.CategoryService
	config  *config.AppConfig
}

func NewCategoryHandler(svc *business.CategoryService, cf *config.AppConfig) *CategoryHandler {
	return &CategoryHandler{
		service: svc,
		config:  cf,
	}
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	category, err := h.service.GetCategory(id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	_ = json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	cats, err := h.service.ListCategories()
	if err != nil {
		http.Error(w, "failed to fetch", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(cats)
}

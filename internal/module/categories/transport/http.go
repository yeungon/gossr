package transport

import (
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"

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

	// Temporary display HTML
	tmpl := template.Must(template.New("exam").Parse(`
	<h1>{{.Title}}</h1>
	<p>This is a category.</p>
	{{.Done}}
`))

	data := struct {
		Title string
		Done  string
	}{
		Title: "You are viewing category ID: " + idStr,
		Done:  "Hello world, this is a category.",
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, data)

	return
	//In your applicaition, fetch article from database as following

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

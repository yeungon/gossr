package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yeungon/gossr/config"
	"github.com/yeungon/gossr/internal/module/articles/business"
)

type ArticleHandler struct {
	service *business.ArticleService
	config  *config.AppConfig
}

func NewArticleHandler(svc *business.ArticleService, cf *config.AppConfig) *ArticleHandler {
	return &ArticleHandler{
		service: svc,
		config:  cf,
	}
}

func (h *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id") // Extract 'id' from the URL path
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "seems invalid id", http.StatusBadRequest)
		return
	}

	test := h.config.APP_DOMAIN_URL

	fmt.Println("Config APP_DOMAIN_URL: id "+test, id)

	article, err := h.service.GetArticle(id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	// Validate the retrieved article
	if err := article.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	_ = json.NewEncoder(w).Encode(article)
}

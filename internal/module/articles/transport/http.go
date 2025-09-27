package transport

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yeungon/gossr/internal/module/articles/business"
)

type ArticleHandler struct {
	service *business.ArticleService
}

func NewArticleHandler(svc *business.ArticleService) *ArticleHandler {
	return &ArticleHandler{service: svc}
}

func (h *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	article, err := h.service.GetArticle(id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	_ = json.NewEncoder(w).Encode(article)
}

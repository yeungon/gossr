package transport

import (
	"encoding/json"
	"fmt"
	"html/template"
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
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Error: invalid id", http.StatusBadRequest)
		return
	}

	// Temporary display HTML
	tmpl := template.Must(template.New("exam").Parse(`
	<h1>{{.Title}}</h1>
	<p>Some random content.</p>
	{{.Done}}
`))

	data := struct {
		Title string
		Done  string
	}{
		Title: "You are viewing article ID: " + idStr,
		Done:  "Hello world, this is an article.",
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, data)

	return

	//In your applicaition, fetch article from database as following

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

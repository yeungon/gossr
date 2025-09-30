package transport

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *ArticleHandler) GetMostViewArticle(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "test")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Error: invalid id", http.StatusBadRequest)
		return
	}

	// Temporary display HTML
	tmpl := template.Must(template.New("exam").Parse(`
	<h1>{{.Title}}</h1>
	<p>most view article.</p>
	{{.Done}}
`))

	data := struct {
		Title string
		Done  string
	}{
		Title: "You are viewing most view article: " + idStr,
		Done:  "Hello world, this is a special article",
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

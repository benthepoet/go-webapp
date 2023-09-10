package admin

import (
	"io"
	"net/http"

	"github.com/benthepoet/go-webapp/internal/templates"
	"github.com/go-chi/chi/v5"
)

var tplMan *templates.TemplateManager = templates.New("./internal/templates/admin", ".mustache")

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", getIndex)
	r.Get("/products", getProducts)

	return r
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	h, err := tplMan.RenderInLayout("index", "layout", map[string]string{"title": "Home"})
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	} else {
		HTML(w, 200, h)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	h, err := tplMan.RenderInLayout("products", "layout", map[string]string{"title": "Products"})
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	} else {
		HTML(w, 200, h)
	}
}

func HTML(w http.ResponseWriter, c int, h string) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(c)
	io.WriteString(w, h)
}

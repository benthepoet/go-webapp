package admin

import (
	"io"
	"net/http"

	"github.com/benthepoet/go-webapp/internal/templates"
	"github.com/go-chi/chi/v5"
)

var tplMan *templates.TemplateManager = templates.New("./internal/templates/admin", ".mustache")

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", getIndex)
	r.Get("/products", getProducts)

	return r
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	h, err := tplMan.RenderInLayout("index", "layout", map[string]string{"title": "Home"})
	if err != nil {
		internalError(w)
	} else {
		HTML(w, http.StatusOK, h)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	h, err := tplMan.RenderInLayout("products", "layout", map[string]string{"title": "Products"})
	if err != nil {
		internalError(w)
	} else {
		HTML(w, http.StatusOK, h)
	}
}

func HTML(w http.ResponseWriter, c int, h string) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	w.WriteHeader(c)
	io.WriteString(w, h)
}

func internalError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

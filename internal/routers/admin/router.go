package admin

import (
	"io"
	"net/http"

	"github.com/benthepoet/go-webapp/internal/templates"
	"github.com/go-chi/chi/v5"
)

type Admin struct {
	TemplateManager *templates.TemplateManager
}

func New() *chi.Mux {
	r := chi.NewRouter()

	t := templates.New("./internal/templates/admin", ".mustache")
	a := &Admin{t}

	r.Get("/", a.index)
	r.Get("/products", a.products)

	return r
}

func (a *Admin) index(w http.ResponseWriter, r *http.Request) {
	h, err := a.TemplateManager.RenderInLayout("index", "layout", map[string]string{"title": "Home"})
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	} else {
		HTML(w, h)
	}
}

func (a *Admin) products(w http.ResponseWriter, r *http.Request) {
	h, err := a.TemplateManager.RenderInLayout("products", "layout", map[string]string{"title": "Products"})
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	} else {
		HTML(w, h)
	}
}

func HTML(w http.ResponseWriter, h string) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	io.WriteString(w, h)
}

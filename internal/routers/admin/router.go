package admin

import (
	"net/http"

	"github.com/benthepoet/go-webapp/internal/response"
	"github.com/benthepoet/go-webapp/internal/templates"
	"github.com/go-chi/chi/v5"
)

var tplMan *templates.TemplateManager = templates.New("./internal/templates/admin", ".mustache")
var tm *templates.TM = templates.NewTM("./internal/templates/admin", map[string][]string{
	"index":    {"index.html", "base.html"},
	"products": {"products.html", "base.html"},
})

func NewRouter() *chi.Mux {

	r := chi.NewRouter()
	r.Get("/", getIndex)
	r.Get("/products", getProducts)

	return r
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	html, err := tm.RenderInLayout("index", "base", map[string]string{"title": "Home"})
	if err != nil {
		response.InternalError(w)
	} else {
		response.Ok(w, html, response.ContentTypeHTML)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	html, err := tm.RenderInLayout("products", "base", map[string]string{"title": "Products"})
	if err != nil {
		response.InternalError(w)
	} else {
		response.Ok(w, html, response.ContentTypeHTML)
	}
}

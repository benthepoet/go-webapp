package admin

import (
	"io"
	"net/http"

	"github.com/benthepoet/go-webapp/internal/templates"
	"github.com/julienschmidt/httprouter"
)

const pathPrefix = "/admin"

type AdminRouter struct {
	Router          *httprouter.Router
	TemplateManager *templates.TemplateManager
}

func New() *AdminRouter {
	r := httprouter.New()
	t := templates.New("./internal/templates/admin", ".mustache")
	a := &AdminRouter{r, t}

	r.GET(pathPrefix, a.index)
	r.GET(pathPrefix+"/products", a.products)

	return a
}

func (a *AdminRouter) index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	h, err := a.TemplateManager.RenderInLayout("index", "layout", map[string]string{"title": "Home"})
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	} else {
		HTML(w, h)
	}
}

func (a *AdminRouter) products(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

package main

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/benthepoet/go-webapp/internal/routers/admin"
	"github.com/julienschmidt/httprouter"
)

type PrefixRouter struct {
	prefixes map[string]*httprouter.Router
	root     *httprouter.Router
}

func main() {
	p := &PrefixRouter{root: httprouter.New(), prefixes: map[string]*httprouter.Router{}}

	p.root.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		io.WriteString(w, "Hello")
	})

	a := admin.New()

	p.prefixes["/admin"] = a.Router

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      p,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	srv.ListenAndServe()
}

func (p *PrefixRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	secureHeaders(w)

	for k, v := range p.prefixes {
		if strings.HasPrefix(r.URL.Path, k) {
			v.ServeHTTP(w, r)
			return
		}
	}

	p.root.ServeHTTP(w, r)
}

func secureHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
}

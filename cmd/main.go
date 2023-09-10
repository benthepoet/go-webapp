package main

import (
	"net/http"

	"github.com/benthepoet/go-webapp/internal/middlewares"
	"github.com/benthepoet/go-webapp/internal/routers/admin"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middlewares.Helmet)
	r.Use(middlewares.BasicAuth)

	r.Mount("/admin", admin.NewRouter())
	r.Mount("/debug", middleware.Profiler())

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	srv.ListenAndServe()
}

package main

import (
	"net/http"
	"time"

	"github.com/benthepoet/go-webapp/internal/middleware"
	"github.com/benthepoet/go-webapp/internal/routers/admin"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Helmet)

	r.Mount("/admin", admin.New())

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	srv.ListenAndServe()
}

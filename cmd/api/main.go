package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// Middlewares opcionais úteis para começar
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Rota simples GET /
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá, Go Chi!"))
	})

	// Exemplo de rota com parâmetro dinâmico
	r.Get("/ola/{nome}", func(w http.ResponseWriter, r *http.Request) {
		nome := chi.URLParam(r, "nome")
		w.Write([]byte("Olá, " + nome + "!"))
	})

	http.ListenAndServe(":8080", r)
}

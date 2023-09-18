package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		d, _ := json.Marshal(map[string]string{"hello": "world"})
		w.Write(d)
	})

	r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
		d, _ := json.Marshal(map[string]string{"hello": "admin"})
		w.Write(d)
	})

	r.Get("/general", func(w http.ResponseWriter, r *http.Request) {
		d, _ := json.Marshal(map[string]string{"hello": "general"})
		w.Write(d)
	})

	http.ListenAndServe(":5001", r)
}

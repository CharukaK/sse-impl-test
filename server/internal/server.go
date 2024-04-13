package internal

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func StartServer() {
	r := chi.NewRouter()
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })
	http.ListenAndServe(":8080", r)
}

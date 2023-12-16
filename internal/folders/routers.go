package folders

import (
	"database/sql"

	"github.com/go-chi/chi"
)

type handler struct {
	db *sql.DB
}

func SetRouters(r chi.Router, db *sql.DB) {
	h := handler{db}
	r.Post("/", h.Create)
}

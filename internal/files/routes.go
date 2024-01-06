package files

import (
	"GoDriver/internal/auth"
	"GoDriver/internal/bucket"
	"GoDriver/internal/queue"
	"database/sql"

	"github.com/go-chi/chi"
)

type handler struct {
	db     *sql.DB
	bucket *bucket.Bucket
	queue  *queue.Queue
}

func SetRoutes(r chi.Router, db *sql.DB, b *bucket.Bucket, q *queue.Queue) {

	r.Group(func(r chi.Router) {
		r.Use(auth.Validate)
		h := handler{db, b, q}
		r.Put("/{id}", h.Modify)
		r.Post("/", h.Create)
		r.Delete("/{id}", h.Delete)
	})

}

package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *handler) GetByID(rw http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err := Get(h.db, id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(u)
}

func Get(db *sql.DB, id int) (*User, error) {
	stmt := `select * from "users" WHERE "id" = $1`
	row := db.QueryRow(stmt, id)
	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Username, &u.Password, &u.CreatedAt, &u.ModifiedAt, &u.Deleted, &u.LastLogin)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

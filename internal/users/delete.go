package users

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (h *handler) Delete(rw http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Delete(h.db, id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
	rw.Header().Add("Content-Type", "application/json")
}

func Delete(db *sql.DB, id int) error {

	stmt := `update "users" set "deleted"= true, "modified_at"=$1 WHERE "id" = $2`
	_, err := db.Exec(stmt, time.Now(), id)
	return err

}

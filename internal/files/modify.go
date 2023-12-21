package files

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (h *handler) Modify(rw http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := Get(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	err = file.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	err = Update(h.db, int64(id), file)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(file)

}

func Update(db *sql.DB, id int64, f *File) error {
	f.ModifiedAt = time.Now()
	stmt := `update "files" set "name"=$1, "modified_at"=$2, "deleted"=$3 WHERE id=$4`
	_, err := db.Exec(stmt, f.Name, f.ModifiedAt, f.Deleted, id)
	return err
}

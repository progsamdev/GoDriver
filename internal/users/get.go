package users

import (
	"database/sql"
	"net/http"
)

func (h *handler) GetByID(rw http.ResponseWriter, r *http.Request) {

}

func Get(db *sql.DB, id int) (*User, error) {
	stmt := `select * from "users" WHERE id = $3`
	row := db.QueryRow(stmt, id)
	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Username, &u.Password, &u.CreatedAt, &u.ModifiedAt, &u.Deleted, &u.LastLogin)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

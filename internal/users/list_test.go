package users

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestList(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	rows := []string{"id", "name", "username", "password", "created_at", "modified_at", "deleted", "last_login"}

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "users" where "deleted" = false`)).
		WillReturnRows(sqlmock.NewRows(rows).
			AddRow(1, "John Doe", "johndoe", "password123", time.Now(), time.Now(), false, time.Now()).
			AddRow(2, "John Doe", "johndoe", "password123", time.Now(), time.Now(), false, time.Now()))

	_, err = SelectAll(db)

	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

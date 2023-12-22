package users

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsert(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	u, err := New("Tiago", "samuel@silva.com.br", "123456")
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec(`insert into "users" ("name", "username", "password", "modified_at")*`).
		WithArgs("Tiago", "samuel@silva.com.br", u.Password, u.ModifiedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = Insert(db, u)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

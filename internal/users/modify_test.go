package users

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestModify(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	id := 1

	mock.ExpectExec(regexp.QuoteMeta(`update "users" set "name"=$1, "modified_at"=$2 WHERE "id"= $3`)).
		WithArgs("Samuel", time.Now(), id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = Update(db, id, &User{Name: "Samuel"})
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

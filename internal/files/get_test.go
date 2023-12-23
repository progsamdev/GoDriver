package files

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGet(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	rows := []string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "files" where "id" = $1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(rows).
			AddRow(1, 2, 3, "Imagen-River.png", ".png", "/", time.Now(), time.Now(), false))

	_, err = Get(db, 1)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

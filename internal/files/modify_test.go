package files

import (
	utl "GoDriver/internal"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUpdate(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	f, err := New(1, "Imagens.goTest", ".png", "/imagens")
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec(regexp.QuoteMeta(`update "files" set "name"=$1, "modified_at"=$2, "deleted"=$3 WHERE "id"=$4`)).
		WithArgs(f.Name, utl.AnyTime{}, f.Deleted, f.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = Update(db, f.ID, f)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

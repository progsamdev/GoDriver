package folders

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetFolder(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	rows := []string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}

	mock.ExpectQuery(regexp.QuoteMeta(`select *from "folders" where "id"=$1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(rows).
			AddRow(1, 2, "Documentos", time.Now(), time.Now(), false))

	_, err = GetFolder(db, 1)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

func TestGetSubFolder(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	rows := []string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}

	mock.ExpectQuery(regexp.QuoteMeta(`select *from "folders" where "parent_id"=$1 and "deleted" = false`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(rows).
			AddRow(2, 3, "Projetos", time.Now(), time.Now(), false).
			AddRow(6, 3, "Projetos Pessoais", time.Now(), time.Now(), false))

	_, err = GetSubFolders(db, 1)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

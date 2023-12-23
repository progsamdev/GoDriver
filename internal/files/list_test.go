package files

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

	columns := []string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "files" where "folder_id" = $1 and "deleted" = false`)).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, 1, 1, "file1.txt", "txt", "/path/to/file1", time.Now(), time.Now(), false).
			AddRow(2, 1, 1, "file2.txt", "txt", "/path/to/file2", time.Now(), time.Now(), false))

	_, err = List(db, 1)

	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}

func TestListRoot(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	columns := []string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "files" where "folder_id" = is null and "deleted" = false`)).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, 1, 1, "file1.txt", "txt", "/path/to/file1", time.Now(), time.Now(), false).
			AddRow(2, 1, 1, "file2.txt", "txt", "/path/to/file2", time.Now(), time.Now(), false))

	_, err = ListRoot(db)

	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}

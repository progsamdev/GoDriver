package folders

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetRootSubFolders(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	rows := []string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}

	mock.ExpectQuery(regexp.QuoteMeta(`select *from "folders" where "parent_id"= is null and "deleted" = false`)).
		WillReturnRows(sqlmock.NewRows(rows).
			AddRow(2, 0, "Docs", time.Now(), time.Now(), false).
			AddRow(6, 0, "Imagens", time.Now(), time.Now(), false))

	_, err = getRootSubFolders(db)

	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

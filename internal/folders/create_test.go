package folders

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

/*

	stmt := `insert into "folders" {"parent_id", "name", "modified"} values ($1, $2, $3)`
	result, err := db.Exec(stmt, f.ParentID, f.Name, f.ModifiedAt)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
*/

func TestCreate(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	f, err := New("fotos", 0)

	if err != nil {
		t.Error(err)
	}
	f.ModifiedAt = time.Now()

	mock.ExpectExec(regexp.QuoteMeta(`insert into "folders" {"parent_id", "name", "modified_at"} VALUES ($1, $2, $3)`)).
		WithArgs(0, f.Name, f.ModifiedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = Insert(db, f)

	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

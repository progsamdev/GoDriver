package users

import (
	utl "GoDriver/internal"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
)

func TestModify(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	h := handler{db}

	u := User{
		ID:         1,
		Name:       "Samuel",
		ModifiedAt: time.Now(),
	}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&u)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/{id}", &b)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	mock.ExpectExec(regexp.QuoteMeta(`update "users" set "name"=$1, "modified_at"=$2 WHERE "id"= $3`)).
		WithArgs(u.Name, utl.AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	rows := []string{"id", "name", "username", "password", "created_at", "modified_at", "deleted", "last_login"}

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "users" WHERE "id" = $1`)).
		WithArgs(u.ID).
		WillReturnRows(sqlmock.NewRows(rows).
			AddRow(1, "John Doe", "johndoe", "password123", time.Now(), time.Now(), false, time.Now()))

	h.Modify(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Error: %v", rr)
	}
}

func TestUpdate(t *testing.T) {

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

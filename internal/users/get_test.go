package users

import (
	"context"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
)

func TestGetById(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	h := handler{db}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/{id}", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	rows := []string{"id", "name", "username", "password", "created_at", "modified_at", "deleted", "last_login"}

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "users" WHERE "id" = $1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(rows).
			AddRow(1, "John Doe", "johndoe", "password123", time.Now(), time.Now(), false, time.Now()))

	h.GetByID(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	rows := []string{"id", "name", "username", "password", "created_at", "modified_at", "deleted", "last_login"}

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "users" WHERE "id" = $1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(rows).
			AddRow(1, "John Doe", "johndoe", "password123", time.Now(), time.Now(), false, time.Now()))

	_, err = Get(db, 1)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

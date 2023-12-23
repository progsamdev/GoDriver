package users

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreate(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	h := handler{db}

	u := User{
		Name:     "Samuel",
		Username: "samuel@silva.com.br",
		Password: "1234567",
	}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&u)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", &b)

	mock.ExpectExec(regexp.QuoteMeta(`insert into "users" ("name", "username", "password", "modified_at") VALUES($1, $2, $3, $4)`)).
		WithArgs(u.Name, u.Username, fmt.Sprintf("%x", (md5.Sum([]byte(u.Password)))), u.ModifiedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	h.Create(rr, req)
	if rr.Code != http.StatusCreated {
		t.Errorf("Error: %v", rr)
	}
}

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

	mock.ExpectExec(regexp.QuoteMeta(`insert into "users" ("name", "username", "password", "modified_at") VALUES($1, $2, $3, $4)`)).
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

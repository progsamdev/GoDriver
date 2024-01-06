package users

import (
	utl "GoDriver/internal"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestCreate() {

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(ts.entity)
	assert.NoError(ts.T(), err)

	ts.entity.SetPassword(ts.entity.Password)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", &b)
	setMock(ts.mock, ts.entity)

	ts.handler.Create(rr, req)
	assert.Equal(ts.T(), http.StatusCreated, rr.Code)

}

func (ts *TransactionSuite) TestInsert() {

	setMock(ts.mock, ts.entity)

	_, err := Insert(ts.conn, ts.entity)
	assert.NoError(ts.T(), err)
}

func setMock(mock sqlmock.Sqlmock, entity *User) {
	mock.ExpectExec(regexp.QuoteMeta(`insert into "users" ("name", "username", "password", "modified_at") VALUES($1, $2, $3, $4)`)).
		WithArgs(entity.Name, entity.Username, entity.Password, utl.AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

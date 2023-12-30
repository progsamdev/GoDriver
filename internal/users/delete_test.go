package users

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestDeleteHTTP() {

	tcs := []struct {
		Id               string
		WithMock         bool
		MockID           int64
		MockWithErr      bool
		ExpectStatusCode int
	}{
		{"1", true, 1, false, http.StatusNoContent},
		{"A", false, -1, true, http.StatusInternalServerError},
		{"25", true, 25, true, http.StatusInternalServerError},
	}

	for _, tc := range tcs {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/{id}", nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("id", tc.Id)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		if tc.WithMock {

			setMockDelete(ts.mock, tc.MockID, tc.MockWithErr)
		}
		ts.handler.Delete(rr, req)
		assert.Equal(ts.T(), tc.ExpectStatusCode, rr.Code)
	}

}

func (ts *TransactionSuite) TestDelete() {

	setMockDelete(ts.mock, 1, false)
	err := Delete(ts.conn, 1)
	assert.NoError(ts.T(), err)
}

func setMockDelete(mock sqlmock.Sqlmock, id int64, err bool) {
	exp := mock.ExpectExec(regexp.QuoteMeta(`update "users" set "deleted"= true, "modified_at"=$1 WHERE "id" = $2`)).
		WithArgs(time.Now(), id)
	if err {
		exp.WillReturnError(sql.ErrNoRows)
	} else {
		exp.WillReturnResult(sqlmock.NewResult(1, 1))
	}

}

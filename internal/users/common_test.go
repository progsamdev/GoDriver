package users

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionSuite struct {
	suite.Suite
	conn *sql.DB
	mock sqlmock.Sqlmock

	handler handler

	entity *User
}

func (ts *TransactionSuite) SetupTest() {
	var err error

	ts.conn, ts.mock, err = sqlmock.New()
	assert.NoError(ts.T(), err)

	ts.handler = handler{ts.conn}

	ts.entity = &User{Name: "Samuel",
		Username: "samuel@silva.com.br",
		Password: "123456"}
}

func (ts *TransactionSuite) AfterTest(_, _ string) {
	assert.NoError(ts.T(), ts.mock.ExpectationsWereMet())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TransactionSuite))
}

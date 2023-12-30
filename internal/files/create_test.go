package files

import (
	utl "GoDriver/internal"
	"GoDriver/internal/bucket"
	"GoDriver/internal/queue"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
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

	b, err := bucket.New(bucket.MockProvider, nil)
	if err != nil {
		t.Error(err)
	}

	q, err := queue.New(queue.Mock, nil)
	if err != nil {
		t.Error(err)
	}

	h := handler{db, b, q}

	//Start UPLOAD
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)

	file, err := os.Open("./testdata/test.jpg")
	if err != nil {
		t.Error(err)
	}

	w, err := mw.CreateFormFile("file", "test.jpg")
	if err != nil {
		t.Error(err)
	}

	_, err = io.Copy(w, file)
	if err != nil {
		t.Error(err)
	}

	mw.Close()

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Add("Content-type", mw.FormDataContentType())

	mock.ExpectExec(regexp.QuoteMeta(`insert into "files" ("folder_id", "owner_id", "name", "type", "path", "modified_at") VALUES ($1, $2,$3,$4,$5,$6)`)).
		WithArgs(0, 1, "test.jpg", "application/octet-stream", "/test.jpg", utl.AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	h.Create(rr, req)
	if rr.Code != http.StatusCreated {
		t.Errorf("Error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}

func TestInsert(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	f, err := New(1, "Imagens.goTest", ".png", "/imagens")
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec(regexp.QuoteMeta(`insert into "files" ("folder_id", "owner_id", "name", "type", "path", "modified_at") VALUES ($1, $2,$3,$4,$5,$6)`)).
		WithArgs(f.FolderID, f.OwnerID, f.Name, f.Type, f.Path, f.ModifiedAt).
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

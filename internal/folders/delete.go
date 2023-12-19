package folders

import (
	"GoDriver/internal/files"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (h *handler) Delete(rw http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = deleteFolderContent(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = deleteSubfolders(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Delete(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")

}

func deleteFolderContent(db *sql.DB, folderID int64) error {

	err := deleteFiles(db, folderID)
	if err != nil {
		return err
	}
	err = deleteSubfolders(db, folderID)
	return err
}

func deleteSubfolders(db *sql.DB, folderID int64) error {

	subFolders, err := getSubFolders(db, folderID)
	if err != nil {
		return err
	}

	removedFolders := make([]Folder, 0, len(subFolders))
	for _, sf := range subFolders {
		err = Delete(db, sf.ID)
		if err != nil {
			break
		}

		err = deleteFolderContent(db, sf.ID)
		if err != nil {
			Update(db, sf.ID, &sf)
			break
		}
		removedFolders = append(removedFolders, sf)
	}

	if len(subFolders) != len(removedFolders) {
		for _, rf := range removedFolders {
			err = Update(db, rf.ID, &rf)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func deleteFiles(db *sql.DB, folderID int64) error {

	f, err := files.List(db, int64(folderID))
	if err != nil {
		return err
	}

	removedFiles := make([]files.File, 0, len(f))
	for _, file := range f {
		file.Deleted = true
		err = files.Update(db, file.ID, &file)
		if err != nil {
			return err
		}
		removedFiles = append(removedFiles, file)
	}

	if len(f) != len(removedFiles) {
		for _, file := range f {
			file.Deleted = false
			err = files.Update(db, file.ID, &file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Delete(db *sql.DB, id int64) error {

	stmt := `update "folders" set "modified_at=$1, "deleted"=true where id = $2`
	_, err := db.Exec(stmt, time.Now(), id)
	return err
}

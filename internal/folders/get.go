package folders

import (
	"GoDriver/internal/files"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *handler) Get(rw http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	f, err := GetFolder(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	c, err := GetFolderContent(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	fc := FolderContent{
		Folder:  *f,
		Content: c,
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(fc)
}

func GetFolder(db *sql.DB, folderID int64) (*Folder, error) {

	stml := `select *from "folders" where id=$1`
	row := db.QueryRow(stml, folderID)

	var f Folder
	err := row.Scan(&f.ID, &f.ParentID, &f.Name, &f.CreatedAt, &f.ModifiedAt, &f.Deleted)

	if err != nil {
		return nil, err
	}

	return &f, nil
}

func getSubFolders(db *sql.DB, folderID int64) ([]Folder, error) {

	stml := `select *from "folders" where parent_id=$1`
	rows, err := db.Query(stml, folderID)
	if err != nil {
		return nil, err
	}

	fs := make([]Folder, 0)
	for rows.Next() {
		var f Folder
		err = rows.Scan(&f.ID, &f.ParentID, &f.Name, &f.CreatedAt, &f.ModifiedAt, &f.Deleted)

		if err != nil {
			return nil, err
		}
		fs = append(fs, f)
	}

	return fs, nil
}

func GetFolderContent(db *sql.DB, folderID int64) ([]FolderResource, error) {
	fs, err := getSubFolders(db, folderID)
	if err != nil {
		return nil, err
	}

	fr := make([]FolderResource, 0, len(fs))
	for _, f := range fs {
		r := FolderResource{
			ID:         f.ID,
			Name:       f.Name,
			Type:       "directory",
			CreatedAt:  f.CreatedAt,
			ModifiedAt: f.ModifiedAt,
		}
		fr = append(fr, r)
	}

	folderFiles, err := files.List(db, folderID)
	if err != nil {
		return nil, err
	}

	for _, ff := range folderFiles {
		r := FolderResource{
			ID:         ff.ID,
			Name:       ff.Name,
			Type:       ff.Type,
			CreatedAt:  ff.CreatedAt,
			ModifiedAt: ff.ModifiedAt,
		}
		fr = append(fr, r)
	}

	return fr, nil
}

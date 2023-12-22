package folders

import (
	"GoDriver/internal/files"
	"database/sql"
	"encoding/json"
	"net/http"
)

func (h *handler) List(rw http.ResponseWriter, r *http.Request) {

	c, err := GetRootFolderContent(h.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	fc := FolderContent{
		Folder:  Folder{Name: "root"},
		Content: c,
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(fc)
}

func getRootSubFolders(db *sql.DB) ([]Folder, error) {

	stml := `select *from "folders" where "parent_id"= is null and "deleted" = false `
	rows, err := db.Query(stml)
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

func GetRootFolderContent(db *sql.DB) ([]FolderResource, error) {
	fs, err := getRootSubFolders(db)
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

	folderFiles, err := files.ListRoot(db)
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

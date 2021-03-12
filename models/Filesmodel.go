package models

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"transfer/db"
	"transfer/log"

	"github.com/joho/godotenv"
)

type Filesmodel struct {
	ID         int    `db:"id"`
	NameFile   string `db:"name_file"`
	DateCreate string `db:"date_create"`
	DateDelete string `db:"date_delete"`
	PathFile   string `db:"path_file"`
	IDUser     string `db:"id_user"`
	IdTarget   string `db:"id_target"`
	Uploaded   bool   `db:"uploaded"`
}

func FilesUpload(b *http.Request, iduser string, idtarget string) bool {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	b.ParseMultipartForm(32 << 20)

	m := b.MultipartForm

	//текущая дата
	currentTime := time.Now()
	tm := fmt.Sprintf(currentTime.Format("01.02.2006"))

	for _, v := range m.File {
		for _, f := range v {

			file, err := f.Open()
			var outfile *os.File
			p := os.Getenv("pathupload")
			if outfile, err = os.Create(p + f.Filename); nil != err {
				//	status = http.StatusInternalServerError
				log.WriteError(err)
			}

			var written int64
			if written, err = io.Copy(outfile, file); nil != err {
				//return
			} else {
				sql = `INSERT INTO tr_files (name_file,  date_create,  date_delete, path_file, id_user, id_target) 
				VALUES (:name_file,  :date_create,  :date_delete, :path_file, :id_user, :id_target) returning id;`

				upl := Filesmodel{NameFile: f.Filename, DateCreate: tm, DateDelete: tm, PathFile: p + f.Filename, IDUser: iduser, IdTarget: idtarget}
				r, err := db.Mgr.GetDB().NamedQuery(sql, upl)

				for r.Next() {
					r.Scan(&upl.ID)
				}

				if err != nil {
					log.WriteError(err)
					return false
				}

				fmt.Println(written)

			}

			if err != nil {
				log.WriteError(err)
				return false
			}

			defer file.Close()

		}

	}
	return true
}

func FilesUpdate() bool {
	return true
}

func FilesInsert() {

}

func FilesCreate() {

}

func FilesSelect(id_user string, uploaded bool) string {

	arg := map[string]interface{}{
		"id_user":  id_user,
		"uploaded": uploaded,
	}

	var fls []*Filesmodel
	//var files []Filesmodel
	var result = ""
	rows, err := db.Mgr.GetDB().NamedQuery(`SELECT * FROM tr_files where id_user=:id_user and uploaded=:uploaded`, arg)

	log.WriteError(err)

	for rows.Next() {
		var fl = &Filesmodel{}
		err := rows.StructScan(&fl)

		if err != nil {
			log.WriteError(err)
		}
		fls = append(fls, fl)
	}
	result = Model.ConvertJS(Model{}, fls)

	return result
}

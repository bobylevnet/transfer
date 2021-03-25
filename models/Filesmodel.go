package models

import (
	"archive/zip"
	"encoding/json"
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
	IDUser     int    `db:"id_user"`
	IDTarget   int    `db:"id_target"`
	Uploaded   bool   `db:"uploaded"`
}

type Usersender []struct {
	IDuser int `json:"IDuser"`
}
type Userfiles struct {
	Filesid int `db:"files_id"`
	Userid  int `db:"user_id"`
}

type Username struct {
	Nameuser string `db:"name_user"`
}

type Targetname struct {
	Targetname string `db:"target_name"`
}

func FilesUpload(b *http.Request, iduser int, idtarget int) bool {

	var usrsndr Usersender
	//путь сохранения файлов
	p := createFolder(iduser, idtarget)
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	b.ParseMultipartForm(32 << 20)

	m := b.MultipartForm
	//
	usersender := []byte(m.Value["usersender"][0])

	json.Unmarshal(usersender, &usrsndr)
	//текущая дата
	currentTime := time.Now()
	tm := fmt.Sprintf(currentTime.Format("01.02.2006"))

	for _, v := range m.File {
		for _, f := range v {

			file, err := f.Open()
			//p := os.Getenv("pathupload")
			//	zipFiles(p + f.Filename)
			var outfile *os.File
			var pthfile = p + f.Filename + ".zip"
			if outfile, err = os.Create(pthfile); nil != err {
				log.WriteError(err)
			}

			zipWriter := zip.NewWriter(outfile)
			log.WriteError(err)

			info, err := outfile.Stat()
			header, err := zip.FileInfoHeader(info)

			header.Name = f.Filename

			header.Method = zip.Deflate
			writer, err := zipWriter.CreateHeader(header)

			var written int64
			if written, err = io.Copy(writer, file); nil != err {
				//return
			} else {
				sql = `INSERT INTO tr_files (name_file,  date_create,  date_delete, path_file, id_user, id_target) 
				VALUES (:name_file,  :date_create,  :date_delete, :path_file, :id_user, :id_target) returning id;`

				upl := Filesmodel{NameFile: f.Filename, DateCreate: tm, DateDelete: tm, PathFile: pthfile, IDUser: iduser, IDTarget: idtarget}
				r, err := db.Mgr.GetDB().NamedQuery(sql, upl)
				log.WriteError(err)
				for r.Next() {
					r.Scan(&upl.ID)
				}

				/*зполняем структуру приндлежности файла пользователям*/

				for _, v := range usrsndr {
					oneuserfiles := Userfiles{
						Userid:  v.IDuser,
						Filesid: upl.ID,
					}
					_, err = db.Mgr.GetDB().NamedExec(`INSERT INTO tr_user_files (files_id,  user_id) 
				VALUES (:files_id,  :user_id);`, oneuserfiles)
				}

				log.WriteError(err)

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

func createFolder(iduser int, idtarget int) string {
	log.WriteError(err)

	trgname := Targetname{}
	usr := Username{}
	err = db.Mgr.GetDB().Get(&usr, `SELECT name_user FROM tr_users WHERE id=$1`, iduser)
	err = db.Mgr.GetDB().Get(&trgname, `SELECT target_name FROM tr_target WHERE id_target=$1`, idtarget)
	p := os.Getenv("pathupload")
	os.MkdirAll(p+`/`+usr.Nameuser+`/`+trgname.Targetname, os.ModePerm)
	path := p + `/` + usr.Nameuser + `/` + trgname.Targetname + `/`
	return path

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

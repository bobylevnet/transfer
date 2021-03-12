package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	s "strings"
	"time"
	"transfer/db"
	ej "transfer/errorj"
	"transfer/intrf"
	"transfer/st"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Basemodel struct {
	TableDB    string
	Idpk       int
	Writeerror intrf.Iwriteresponse
	Request    *http.Request
}

var dynamycmodel map[string]interface{}
var m st.Abstracttbl
var err error
var sql string

//var mtd []st.Metadatatable
var clmn1 string
var clmn2 string
var clmn3 string

var bin []byte
var chker ej.Errormessage

func loadmodel(dataJS map[string]string) map[string]interface{} {
	dynamyc := make(map[string]interface{})
	for key, value := range dataJS {
		dynamyc[key] = value
	}
	return dynamyc
}

//generatemeta - выборка метаданных таблицы
func generatmeta(table string, dyn *map[string]interface{}) string {
	clmn1 = ""
	clmn2 = ""
	clmn3 = ""
	var mtd []st.Metadatatable

	d := make(map[string]interface{})
	//dyn = &d
	d = *dyn
	sql = fmt.Sprintf(`select column_name from INFORMATION_SCHEMA.COLUMNS where table_name = '%s';`, table)
	err = db.Mgr.GetDB().Select(&mtd, sql)

	for _, value := range mtd {
		if value.Columnname != "id" {
			clmn1 += ` :` + value.Columnname + `,`
			clmn2 += ` "` + value.Columnname + `",`

			if d[value.Columnname] != nil {
				clmn3 += ` ` + value.Columnname + `=:` + value.Columnname + `,`

			} else {
				d[value.Columnname] = ""
			}

		}
	}
	clmn1 = s.TrimSuffix(clmn1, ",")
	clmn2 = s.TrimSuffix(clmn2, ",")
	clmn3 = s.TrimSuffix(clmn3, ",")
	return ""
}

//Update - обновление записи
func (b Basemodel) Update(modelUpdate st.Abstracttbl, result *string) bool {
	//генерируем динамическую модель
	dyn := make(map[string]interface{})
	dyn = loadmodel(modelUpdate.Dt)
	//строки для запроса
	generatmeta(b.Gettable(), &dyn)
	sql = fmt.Sprintf(`UPDATE %s SET %s  WHERE id=%d`, b.Gettable(), clmn3, b.Getid())
	_, err = db.Mgr.GetDB().NamedExec(sql, dyn)
	chker.Checkerror(b.Writeerror, "Update", err)
	return true
}

func (b Basemodel) Delete() bool {

	sql = fmt.Sprintf(`DELETE FROM %s WHERE id=%d`, b.Gettable(), b.Getid())
	_, err := db.Mgr.GetDB().DB.Exec(sql)
	chker.Checkerror(b.Writeerror, "Update", err)
	return true
}

func (b Basemodel) Insert(modelInsert st.Abstracttbl) int {
	dyn := make(map[string]interface{})
	dyn = loadmodel(modelInsert.Dt)
	//строки для запроса
	generatmeta(b.Gettable(), &dyn)
	sql = fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, b.Gettable(), clmn2, clmn1)
	_, err = db.Mgr.GetDB().NamedExec(sql, dyn)
	chker.Checkerror(b.Writeerror, "INSERT", err)
	return 0

}

func (b Basemodel) Getone(modelFind st.Abstracttbl, result *string) {
	var rows *sqlx.Rows
	dyn := make(map[string]interface{})
	dyn = loadmodel(modelFind.Dt)
	generatmeta(b.Gettable(), &dyn)
	sql = fmt.Sprintf(`SELECT * FROM %s where id=%d `, b.Gettable(), b.Getid())
	rows, err = db.Mgr.GetDB().Queryx(sql)
	chker.Checkerror(b.Writeerror, "select one", err)
	selectdata(result, rows, dyn)
	return
}

func (b Basemodel) Find(modelFind st.Abstracttbl, result *string) {
	var filter string
	var rows *sqlx.Rows
	dyn := make(map[string]interface{})
	dyn = loadmodel(modelFind.Dt)
	loadmodel(modelFind.Dt)
	generatmeta(b.Gettable(), &dyn)
	//фильтр по одному полю
	for key, value := range modelFind.Dt {
		filter += key + " LIKE '" + value + "%'"
	}
	//сделать обработчик sql иньекций
	//sqlandwhere := "and where "
	sql = fmt.Sprintf(`SELECT * FROM %s where  %s`, b.Gettable(), filter)
	rows, err = db.Mgr.GetDB().Queryx(sql)
	chker.Checkerror(b.Writeerror, "FIND", err)
	selectdata(result, rows, dyn)
	return
}

//findWhere - выибраем по   нескольким полям
func (b Basemodel) FindWhere(modelFind st.Abstracttbl, result *string) {
	var filter string
	var rows *sqlx.Rows
	var i int
	i = 0

	dyn := make(map[string]interface{})
	dyn = loadmodel(modelFind.Dt)
	loadmodel(modelFind.Dt)
	generatmeta(b.Gettable(), &dyn)
	//фильтр по одному полю
	for key, value := range modelFind.Dt {
		if i == 0 {
			filter = key + "=" + value
		} else {
			filter = filter + " AND " + key + "=" + value
		}
		i++
	}

	//сделать обработчик sql иньекций
	sql = fmt.Sprintf(`SELECT * FROM %s where  %s`, b.Gettable(), filter)
	rows, err = db.Mgr.GetDB().Queryx(sql)
	chker.Checkerror(b.Writeerror, "FINDWHERE", err)
	selectdata(result, rows, dyn)
	return
}

func (b Basemodel) Getall(modelFind st.Abstracttbl, result *string) {
	var rows *sqlx.Rows
	dyn := make(map[string]interface{})
	dyn = loadmodel(modelFind.Dt)

	generatmeta(b.Gettable(), &dyn)
	sql = fmt.Sprintf(`SELECT * FROM %s`, b.Gettable())
	rows, err = db.Mgr.GetDB().Queryx(sql)
	chker.Checkerror(b.Writeerror, "SELECT ALL", err)
	selectdata(result, rows, dyn)
	//fmt.Println("result- all" + *result)

	return
}

func (b Basemodel) Gettablemeta(model st.Abstracttbl, result *string) {

}

func (b Basemodel) Upload(model st.Abstracttbl, result *string) {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	b.Request.ParseMultipartForm(32 << 20)

	m := b.Request.MultipartForm

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
				return
			}

			var written int64
			if written, err = io.Copy(outfile, file); nil != err {
				//return
			} else {
				sql = fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, b.Gettable(),
					"name_file,  date_create,  date_delete, path_file, id_user, id_target",
					":name_file,  :date_create,  :date_delete, :path_file, :id_user, :id_target")
				upl := st.Upload{NameFile: f.Filename, DateCreate: tm, DateDelete: tm, PathFile: p + f.Filename, IDUser: 0, Idtarget: 0}
				_, err = db.Mgr.GetDB().NamedExec(sql, upl)
				chker.Checkerror(b.Writeerror, "Upload", err)

				//find()
				fmt.Println(written)

			}

			if err != nil {
				fmt.Println(err)
				return
			}

			defer file.Close()

		}
	}

	*result = "Upload"
}

//выборка данных из бд
func selectdata(result *string, rows *sqlx.Rows, dynamyc map[string]interface{}) {

	var i int

	/*dmodel := make(map[string]interface{})
	//копируем кату dynnamycmodel в dmodel
	for key, value := range dyn {
		dmodel[key] = value
	}  */

	*result = ""
	tmp := make(map[int]interface{})
	if len(dynamyc) > 0 {

		for rows.Next() {
			i++
			t := make(map[string]interface{})
			err := rows.MapScan(t)
			tmp[i] = t
			if err != nil {
				fmt.Println(err)
			}
		}

		if len(tmp) > 0 {
			bin, _ := json.Marshal(tmp)
			*result = string(bin)

			fmt.Println("select" + *result)
			//*result = ""
		} else {
			//	dynamycmodel := make(map[int]interface{})
			bin, _ := json.Marshal(dynamyc)
			*result = string(bin)
			//fmt.Println("select" + *result)
			//*result = ""
		}
	}
	if err != nil {
		fmt.Println("error:", err)
	}

}

///сделать обработку ошибок если поле не становленно

func (b Basemodel) Gettable() string {
	return b.TableDB
}

func (b Basemodel) Getid() int {
	return b.Idpk
}

package models

import (
	"encoding/json"
	"fmt"
	s "strings"
	"transfer/db"
	ej "transfer/errorj"
	"transfer/intrf"
	"transfer/st"

	"github.com/jmoiron/sqlx"
)

type Basemodel struct {
	TableDB    string
	Idpk       int
	Writeerror intrf.Iwriteresponse
}

var dynamycmodel map[string]interface{}
var m st.Abstracttbl
var err error
var sql string
var mtd []st.Metadatatable
var clmn1 string
var clmn2 string
var clmn3 string
var rows *sqlx.Rows
var bin []byte
var chker ej.Errormessage

func loadmodel(dataJS map[string]string) *map[string]interface{} {
	dynamycmodel = make(map[string]interface{})
	for key, value := range dataJS {
		dynamycmodel[key] = value
	}
	return &dynamycmodel
}

//generatemeta - выборка метаданных таблицы
func generatmeta(table string) string {
	sql = fmt.Sprintf(`select column_name, data_type
	from INFORMATION_SCHEMA.COLUMNS where table_name = '%s';`, table)
	err = db.Mgr.GetDB().Select(&mtd, sql)

	for _, value := range mtd {
		if value.Columnname != "id" {
			clmn1 += ` :` + value.Columnname + `,`
			clmn2 += ` ` + value.Columnname + `,`

			if dynamycmodel[value.Columnname] != nil {
				clmn3 += ` ` + value.Columnname + `=:` + value.Columnname + `,`
			}

		}
	}
	clmn1 = s.TrimSuffix(clmn1, ",")
	clmn2 = s.TrimSuffix(clmn2, ",")
	clmn3 = s.TrimSuffix(clmn3, ",")
	return ""
}

//Update - обновление записи
func (b Basemodel) Update(modelUpdate *st.Abstracttbl, result *string) bool {
	//генерируем динамическую модель
	loadmodel(modelUpdate.Dt)
	//строки для запроса
	generatmeta(b.Gettable())

	sql = fmt.Sprintf(`UPDATE %s SET %s  WHERE id=%d`, b.Gettable(), clmn3, b.Getid())
	_, err = db.Mgr.GetDB().NamedExec(sql, dynamycmodel)
	chker.Checkerror(b.Writeerror, "Update", err)
	return true
}

func (b Basemodel) Delete() bool {

	sql = fmt.Sprintf(`DELETE FROM %s WHERE id=%d`, b.Gettable(), b.Getid())
	_, err := db.Mgr.GetDB().DB.Exec(sql)
	chker.Checkerror(b.Writeerror, "Update", err)
	return true
}

func (b Basemodel) Insert(modelInsert *st.Abstracttbl) int {
	//генерируем динамическую модель
	loadmodel(modelInsert.Dt)
	//строки для запроса
	generatmeta(b.Gettable())
	sql = fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, b.Gettable(), clmn2, clmn1)
	_, err = db.Mgr.GetDB().NamedExec(sql, dynamycmodel)
	chker.Checkerror(b.Writeerror, "INSERT", err)
	return 0

}

func (b Basemodel) Getone(modelFind *st.Abstracttbl, result *string) {
	loadmodel(modelFind.Dt)
	sql = fmt.Sprintf(`SELECT * FROM %s where id=%d `, b.Gettable(), b.Getid())
	rows, err = db.Mgr.GetDB().Queryx(sql)
	chker.Checkerror(b.Writeerror, "Update", err)
	selectdata(result)
	return
}

func (b Basemodel) Find(modelFind *st.Abstracttbl, result *string) {
	var filter string
	loadmodel(modelFind.Dt)
	//фильтр по одному полю

	for key, value := range modelFind.Dt {
		filter += key + " LIKE '" + value + "%'"
	}
	//сделать обработчик sql иньекций
	sql = fmt.Sprintf(`SELECT * FROM %s where  %s`, b.Gettable(), filter)
	rows, err = db.Mgr.GetDB().Queryx(sql)
	chker.Checkerror(b.Writeerror, "FIND", err)
	selectdata(result)

	return
}

func (b Basemodel) Getall(modelFind *st.Abstracttbl, result *string) {
	//геерируем динамическую модель
	loadmodel(modelFind.Dt)
	sql = fmt.Sprintf(`SELECT * FROM %s`, b.Gettable())
	rows, err = db.Mgr.GetDB().Queryx(sql)
	chker.Checkerror(b.Writeerror, "SELECT ALL", err)
	selectdata(result)

	return
}

//выборка данных из бд
func selectdata(result *string) {

	var i int
	tmp := make(map[int]interface{})
	for rows.Next() {
		i++
		err := rows.MapScan(dynamycmodel)
		tmp[i] = dynamycmodel
		if err != nil {
			fmt.Println(err)
		}
	}

	bin, err := json.Marshal(tmp)

	*result = string(bin)

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

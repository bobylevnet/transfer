package models

import (
	"transfer/db"
	"transfer/log"
)

var table = "tr_files"

type Targetmodel struct {
	UserTargetID int    `db:"user_target_id"`
	TargetName   string `db:"target_name"`
	IDuser       string `db:"id_user"`
	FullNameUser string `db:"full_name_user"`
}

func TargetUpdate() bool {
	return true
}

func TargetInsert() {

}

func TargetCreate() {

}

/*func TargetSelectLikeAction(id_user string, filter string) string {

	var fls []*Targetmodel
	//var files []Filesmodel
	var result = ""
	err := db.Mgr.GetDB().Select(&fls, `SELECT * FROM tr_vuser_target where id_user=$1 AND  target_name LIKE '%$2%'`, id_user, filter)

	log.WriteError(err)

	result = Model.ConvertJS(Model{}, fls)

	return result
}*/

func TargetSelect(id_user string) string {

	arg := map[string]interface{}{
		"id_user": id_user,
	}

	var fls []*Targetmodel
	//var files []Filesmodel
	var result = ""
	rows, err := db.Mgr.GetDB().NamedQuery(`SELECT * FROM tr_vuser_target where id_user=:id_user`, arg)

	log.WriteError(err)

	for rows.Next() {
		var fl = &Targetmodel{}
		err := rows.StructScan(&fl)

		if err != nil {
			log.WriteError(err)
		}
		fls = append(fls, fl)
	}
	result = Model.ConvertJS(Model{}, fls)

	return result
}

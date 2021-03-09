package models

import (
	"fmt"
	"transfer/db"
	l "transfer/log"
)

type Target struct {
	Usertarget int `db:"user_target_id"`
}

type Users struct {
	IDuser       int    `db:"id"`
	Nameuser     string `db:"name_user"`
	Fullnameuser string `db:"full_name_user"`
	Accesstoken  string
	Target       []Target
}

var tbl string = "tr_users"
var tbl_rtarget string = "tr_user_target"

func Chekuser(user string, AccessTo string) Users {

	us := Users{
		Accesstoken: AccessTo,
	}

	//извлекаем пользователя если он есть в нашей БД
	var sql string = fmt.Sprintf("SELECT * FROM %s where name_user='%s' LIMIT 1 ", tbl, user)
	err := db.Mgr.GetDB().Get(&us, sql)
	l.WriteError(err)

	var sqltarget string = fmt.Sprintf("SELECT user_target_id FROM %s where id_user=%d", tbl_rtarget, us.IDuser)
	target := []Target{}
	err = db.Mgr.GetDB().Select(&target, sqltarget)

	//зписываем в основную структуру цели пользователя
	us.Target = target
	l.WriteError(err)

	return us

}

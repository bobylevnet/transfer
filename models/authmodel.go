package models

import (
	"fmt"
	"transfer/db"
	l "transfer/log"
)

type Users struct {
	ID           int    `db:"id"`
	NameUser     string `db:"name_user"`
	FullNameUser string `db:"full_name_user"`
	AccessToken  string
}

var tbl string = "tr_users"

func Chekuser(user string, AccessTo string) Users {

	us := Users{
		AccessToken: AccessTo,
	}
	var sql string = fmt.Sprintf("SELECT * FROM %s where name_user='%s' LIMIT 1 ", tbl, user)

	err := db.Mgr.GetDB().Get(&us, sql)
	l.WriteError(err)

	return us

}

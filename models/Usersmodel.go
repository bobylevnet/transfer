package models

import (
	"transfer/db"
	"transfer/log"
)

type Usersv struct {
	IDuser      int    `db:"id_user"`
	FullNameUser string `db:"full_name_user"`

}

func UsersSelect(targetsid string) string {

	var users []*Usersv

	var result = ""
	arg := map[string]interface{}{
		"target_id": targetsid,
	}

	rows, err := db.Mgr.GetDB().NamedQuery(`
	select id_user, full_name_user from tr_vuser_target tvt  
	where user_target_id=:target_id`, arg)

	log.WriteError(err)

	for rows.Next() {
		var us = &Usersv{}
		err := rows.StructScan(&us)

		if err != nil {
			log.WriteError(err)
		}
		users = append(users, us)
	}
	result = Model.ConvertJS(Model{}, users)

	return result
}

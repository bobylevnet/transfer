package controllers

import (
	"transfer/intrf"
	"transfer/st"
)

var result string

//ActionController действие с таблицами
func ActionController(model intrf.Model, action map[string]string, t intrf.Iwriteresponse, dataJSON map[string]string) bool {
	//	Dt := *dataJSON
	tbl := st.Abstracttbl{Idpk: 0, Dt: dataJSON}

	switch action["action"] {
	case "delete":
		model.Delete()
		t.Writeresponse(&result)

	case "update":
		model.Update(tbl, &result)
	//	t.Writeresponse()

	case "insert":
		model.Insert(tbl)
	case "findone":
		model.Getone(tbl, &result)
		t.Writeresponse(&result)

	case "find":
		model.Find(tbl, &result)
		t.Writeresponse(&result)

	case "findall":
		model.Getall(tbl, &result)
		t.Writeresponse(&result)

	case "findwhere":
		model.FindWhere(tbl, &result)
		t.Writeresponse(&result)

	case "table":
		model.Gettablemeta(tbl, &result)
		t.Writeresponse(&result)
	case "upload":
		model.Upload(tbl, &result)

	default:
		//t.Writeresponse("Нет такого метода")
	}

	return true
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	a "transfer/auth"
	cnt "transfer/controllers"
	"transfer/dbmodel"
	l "transfer/log"
	"transfer/models"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/auth", auth)

	router.HandleFunc("/route", getRoute)
	//добавляем корневой путь для статики
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
	http.ListenAndServe(":89", router)
}

func getRoute(w http.ResponseWriter, r *http.Request) {

	actionDT := map[string]string{}
	dataDT := map[string]string{}
	//filterDT := map[string]string{}

	actionRQ, _ := r.URL.Query()["rest"]
	dataRQ, _ := r.URL.Query()["datajs"]
	idpkRQ, _ := r.URL.Query()["idpk"]
	//filterRQ, _ := r.URL.Query()["filter"]

	//действие
	convertJS(&actionDT, actionRQ[0])
	//данные
	convertJS(&dataDT, dataRQ[0])
	//фильтр данных
	//convertJS(&filterDT, filterRQ[0])

	//bm := models.Basemodel{}

	//базовая модель
	model := dbmodel.Model{}

	bm := dbmodel.Basesql{
		SQL: model,
	}

	bm.TableDB = "tr_" + actionDT["model"]

	if idpkRQ != nil {
		idpk, err := strconv.Atoi(idpkRQ[0])
		bm.Idpk = idpk
		if err != nil {
			// handle error
			fmt.Println(err)
		}
	}

	wr := cnt.Writecontroller{}
	wr.T = w

	//предаем адрес  записи
	//неправильно нужно пределать обработчик ошибок должен быть реализован везде

	bm.Writeerror = &wr
	if r.Method == "POST" {

		bm.Request = r

		cnt.ActionController(bm, actionDT, wr, dataDT)
	} else {
		cnt.ActionController(bm, actionDT, wr, dataDT)
	}

}

/* func do(w http.ResponseWriter, r *http.Request) {
	go getRoute(w, r)
} */

func convertJS(stst *map[string]string, reqJSON string) {

	//var v []interface{}
	//var tmp [string]string
	//tmp := make(map[string]string)
	b := []byte(reqJSON)

	//	s := st.Filtered{}
	//st["id_target"].(string)
	json.Unmarshal([]byte(b), &stst)

	/* switch v.(type) {

	case map[string]interface{}: */

	//for _, element := range v {

	//	fmt.Print(element)
	//tmp[key.(string)] = element.(string)
	/*for _, subelement := range element.([]interface{}) {
		if subelement != "" {
			tmp[key] = subelement.(string)
		} else {
			tmp[key] = element.(string)
		}

	}*/

}

//switch element.(type) { */
/* 	case []interface{}:

	for key, selement := range element.(map(interface{})) {
		tmp[key] = selement.(string)
	}

default:
	l.WriteError(err)

}
tmp[key] = element.(string) */
//	}

/* default:
	l.WriteError(err)
} */

//}

func auth(w http.ResponseWriter, r *http.Request) {
	//id, _ := r.URL.Query()["id"]
	//запрос на авторизацию моего приложения
	au := a.Auth()

	//проверяем есть данный пользователь в БД
	t := models.Chekuser(au.UserName, au.AccessToken)

	if t.Nameuser == "" {
		fmt.Fprintf(w, "Доступ запрещен!")
		//	os.Exit(1)
	} else {
		res, err := json.Marshal(t)
		l.WriteError(err)
		fmt.Printf(string(res))
		fmt.Fprintf(w, string(res))
	}

}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	cnt "transfer/controllers"
	"transfer/models"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/route", getRoute)
	//добавляем корневой путь для статики
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
	http.ListenAndServe(":89", router)
}

func getRoute(w http.ResponseWriter, r *http.Request) {

	actionDT := map[string]string{}
	dataDT := map[string]string{}

	actionRQ, _ := r.URL.Query()["rest"]
	dataRQ, _ := r.URL.Query()["datajs"]
	idpkRQ, _ := r.URL.Query()["idpk"]

	convertJS(&actionDT, actionRQ[0])
	convertJS(&dataDT, dataRQ[0])

	bm := models.Basemodel{}

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
	b := []byte(reqJSON)
	err := json.Unmarshal([]byte(b), &stst)
	if err != nil {
		panic(err)
	}

}

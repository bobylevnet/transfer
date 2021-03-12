package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	a "transfer/auth"
	cnt "transfer/controllers"
	l "transfer/log"
	"transfer/models"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/auth", auth)

	//роуты файлов
	router.HandleFunc("/files/upload/{iduser:[0-9]+}/{idtarget:[0-9]+}", cnt.FilesUploadAction)
	router.HandleFunc("/files/select/{iduser:[0-9]+}", cnt.FilesSelectAction)

	//роуты целей
	router.HandleFunc("/target/select/{iduser:[0-9]+}", cnt.TargetSelectAction)

	//выборка всех пользователей кому можно отправлять 
	router.HandleFunc("/users/targets/{idtarget:[0-9]+}", cnt.UsersSelectAction)



	//router.HandleFunc("/target", getRoute)
	//добавляем корневой путь для статики
	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))

	http.ListenAndServe(":89", router)
}

func getRoute(w http.ResponseWriter, r *http.Request) {

}

/* func do(w http.ResponseWriter, r *http.Request) {
	go getRoute(w, r)
} */

func convertJS(stst *map[string]string, reqJSON string) {

	//var v []interface{}
	//var tmp [string]string
	//tmp := make(map[string]string)
	//b := []byte(reqJSON)

	//	s := st.Filtered{}
	//st["id_target"].(string)
	//json.Unmarshal([]byte(b), &stst)

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

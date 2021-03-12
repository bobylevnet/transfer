package controllers

import (
	"fmt"
	"net/http"
	md "transfer/models"

	"github.com/gorilla/mux"
)


func UsersSelectAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	result := md.UsersSelect(vars["idtarget"])
	fmt.Fprintf(w, result)

}

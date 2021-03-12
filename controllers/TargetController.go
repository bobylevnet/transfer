package controllers

import (
	"fmt"
	"net/http"
	md "transfer/models"

	"github.com/gorilla/mux"
)

func TargetUploadAction(w http.ResponseWriter, r *http.Request) {

}

func TargetDeleteAction(w http.ResponseWriter, r *http.Request) {

}

func TargetSaveAction(w http.ResponseWriter, r *http.Request) {

}

func TargetSelectAction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	result := md.TargetSelect(vars["iduser"])
	fmt.Fprintf(w, result)

}



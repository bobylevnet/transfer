package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"transfer/log"
	md "transfer/models"

	"github.com/gorilla/mux"
)

func FilesUploadAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	iduser, err := strconv.Atoi(vars["iduser"])
	log.WriteError(err)

	idtarget, err := strconv.Atoi(vars["idtarget"])
	log.WriteError(err)

	if md.FilesUpload(r, iduser, idtarget) {
		fmt.Fprintf(w, "true")
	}

}

func FilesDeleteAction(w http.ResponseWriter, r *http.Request) {

}

func FilesSaveAction(w http.ResponseWriter, r *http.Request) {

}

func FilesSelectAction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	result := md.FilesSelect(vars["iduser"], false)

	fmt.Fprintf(w, result)

}

package controllers

import (
	"fmt"
	"net/http"
	md "transfer/models"

	"github.com/gorilla/mux"
)

func FilesUploadAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if md.FilesUpload(r, vars["IDuser"], vars["IDtarget"]) {

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

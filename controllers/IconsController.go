package controllers

import (
	"fmt"
	"net/http"
	md "transfer/models"
)

//возвращаем иконку
func IconsGetAction(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	result := md.GetIcons(r)
	fmt.Fprintf(w, result)

}

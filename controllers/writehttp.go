package controllers

import (
	"fmt"
	"net/http"
)

type Writecontroller struct {
	T http.ResponseWriter
}

func (w Writecontroller) Writeresponse(result string) {
	fmt.Fprintf(w.T, result)
}

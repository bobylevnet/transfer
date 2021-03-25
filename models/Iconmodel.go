package models

import (
	//"transfer/log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"transfer/log"
	//"fmt"
	//"encoding/json"
)

type Iconmodel []struct {
	Nameicon string
	Svgicon  string
}

func GetIcons(b *http.Request) string {
	var t Iconmodel
	var icons Iconmodel
	body, err := ioutil.ReadAll(b.Body)

	log.WriteError(err)
	json.Unmarshal(body, &t)

	//var icons []Iconmodel

	for _, result := range t {
		content, err := ioutil.ReadFile("./web/assets/zondicons/" + result.Nameicon + ".svg")
		result.Svgicon = string(content)
		icons = append(icons, result)
		log.WriteError(err)
	}

	resultjson := Model.ConvertJS(Model{}, icons)
	return resultjson

}

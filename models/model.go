package models

import (
	"encoding/json"
	"transfer/log"
)

type Model struct {
	data interface{}
}

func (m Model)  ConvertJS(data interface{}) string {
	bin, err := json.Marshal(data)
	log.WriteError(err)
	return string(bin)

}

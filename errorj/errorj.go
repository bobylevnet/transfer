package errorj

import (
	"encoding/json"
	"transfer/intrf"
)

type Errormessage struct {
	Error   bool
	Message string
}

func (m Errormessage) Checkerror(w intrf.Iwriteresponse, mess string, err error) {

	if err != nil {
		var ms string
		m.Error = true
		m.Message = mess + " : " + err.Error()
		bin, _ := json.Marshal(m)
		ms = string(bin)
		w.Writeresponse(&ms)

	}

}

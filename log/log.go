package log

import (
	"encoding/json"

	"go.uber.org/zap"
)

var log *zap.Logger

func init() {

	//if erstr != "" {
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "./tmp/error.log"],
	  "errorOutputPaths": ["stderr"],	
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase",
	    "timeKey": "ts",
	    "timeEncoder": "ISO8601" 
	  }
	}`)

	var cfg zap.Config

	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	log = logger
	//}

}

func WriteError(err error) {
	if err != nil {
		log.Error(err.Error())
	}
}

/* import
 (
	"log"
	"os"
)

var Logg *log.Logger

func init() {
	f, err := os.OpenFile("./logfile/error.log", os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	Logg = log.New(f, "Error\t", log.Ldate|log.Ltime)

}
*/

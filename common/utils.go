package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type configuration struct {
	Server, MongoDBHost, DBUser, DBPwd, Database string
}

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}

	errorResource struct {
		Data appError `json:"data"`
	}
)

var AppConfig configuration

//initialize AppConfig

func initConfig() {
	loadAppConfig()
}

//read config.json and decode into AppConfig

func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("loadConfig %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("loadConfig %s\n", err)
	}
}

func DisplayAppError(w http.ResponseWriter, handleError error, message string, code int) {
	errObj := appError{
		Error:      handleError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Print("AppError %s\n", handleError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}

}

package api

import (
	"encoding/json"
	"net/http"
)

type ChosenAthleteParam struct {
	Username string
}


type Athlete struct {
	Id string //`json:"id"`
    Name string	//`json:"name"`
    Age  int //`json:"age"`
	Sport string //`json:"sport"`	
	Position string	//`json:"position"`
}


type Error struct {
	
	Code int
	
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error {
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, "An Unexpected Error Occurred", http.StatusInternalServerError)
	}
)

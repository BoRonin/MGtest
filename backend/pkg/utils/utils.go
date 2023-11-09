package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	message, err := json.Marshal(v)
	if err != nil {
		w.Write([]byte("Could not marshal. Moscow calling"))
		return
	}
	w.Write([]byte(message))
}

func ErrorJSON(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	errorMap := make(map[string]string)
	errorMap["error"] = err.Error()
	message, err := json.Marshal(errorMap)
	w.WriteHeader(code)
	if err != nil {
		w.Write([]byte("Could not marshal. Moscow calling"))
		return
	}
	w.Write([]byte(message))
}

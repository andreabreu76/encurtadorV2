package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(json)
}

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	resp := map[string]interface{}{"message": message}
	respJSON, _ := json.Marshal(resp)
	ResponseWithJSON(w, respJSON, code)
}

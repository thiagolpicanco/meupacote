package handlers

import (
	"encoding/json"
	"net/http"
)

func respondeJSON(w http.ResponseWriter, status int, corpo interface{}) {
	response, err := json.Marshal(corpo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondeErro(w http.ResponseWriter, code int, msg string) {
	respondeJSON(w, code, map[string]string{"erro": msg})
}

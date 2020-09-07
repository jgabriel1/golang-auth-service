package routes

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseJSON struct {
	Message string `json:"error"`
}

func jsonError(err error) []byte {
	errorMessage := err.Error()

	response := ErrorResponseJSON{
		Message: errorMessage,
	}

	js, _ := json.Marshal(response)
	return js
}

func JSONErrorResponse(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(code)

	js := jsonError(err)

	w.Write(js)
}

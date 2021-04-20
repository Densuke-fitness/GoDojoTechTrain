package view

import (
	"net/http"
)

func SuccessView(resp http.ResponseWriter, result []byte) {
	resp.Header().Set("Content-type", "application/json")

	resp.WriteHeader(http.StatusOK)
	resp.Write(result) //nolint
}

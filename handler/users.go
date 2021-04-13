package handler

import (
	"fmt"
	"net/http"
)

func CreateUser() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprint(resp, "test")
	}
}

package controller

import (
	"net/http"

	"github.com/Densuke-fitness/GoDojoTechTrain/model/users"
)

func CreateUser() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		users.CreateUser(resp, req)
	}
}

func GetUser() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		users.GetUser(resp, req)
	}
}

func UpdateUser() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		users.UpdateUser(resp, req)
	}
}

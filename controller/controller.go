package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Densuke-fitness/GoDojoTechTrain/model/users"
	"github.com/Densuke-fitness/GoDojoTechTrain/view"
)

func CreateUser() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		//Structure to be stored when a request is received from a user
		var request = struct {
			Name string `json:"name"`
		}{}

		err := json.NewDecoder(req.Body).Decode(&request)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusBadRequest,
				Message:    "Error unmarshalling the request",
			}
			view.ErrorView(resp, params)
			return
		}

		// Passing values to the model and executing the process
		token, err := users.CreateUser(request.Name)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error executing model process",
			}
			view.ErrorView(resp, params)
			return
		}

		response := struct {
			Token string `json:"token"`
		}{Token: token}

		result, err := json.Marshal(&response)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error marshalling data",
			}
			view.ErrorView(resp, params)
			return
		}

		view.SuccessView(resp, result)
	}
}

func GetUser() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {

		token := req.Header.Get("X-Auth-Token")
		name, err := users.GetUser(token)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error executing model process",
			}
			view.ErrorView(resp, params)
			return
		}

		response := struct {
			Name string `json:"name"`
		}{Name: name}

		result, err := json.Marshal(&response)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error marshalling data",
			}
			view.ErrorView(resp, params)
			return
		}

		view.SuccessView(resp, result)
	}
}

func UpdateUser() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {

		var request = struct {
			Name string `json:"name"`
		}{}
		err := json.NewDecoder(req.Body).Decode(&request)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error unmarshalling the request",
			}
			view.ErrorView(resp, params)
			return
		}

		//fetch token and extract userid
		token := req.Header.Get("X-Auth-Token")

		err = users.UpdateUser(request.Name, token)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error executing model process",
			}
			view.ErrorView(resp, params)
			return
		}
		view.SuccessView(resp, nil)
	}
}

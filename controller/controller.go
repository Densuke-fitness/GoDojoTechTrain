package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Densuke-fitness/GoDojoTechTrain/model/users"
	"github.com/Densuke-fitness/GoDojoTechTrain/view"
)

func CreateUser() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		//Structure to be stored when a reqParams is received from a user
		var reqParams = struct {
			Name string `json:"name"`
		}{}

		err := json.NewDecoder(req.Body).Decode(&reqParams)
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
		token, err := users.CreateUser(reqParams.Name)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error executing model process",
			}
			view.ErrorView(resp, params)
			return
		}

		result, err := json.Marshal(
			&struct {
				Token string `json:"token"`
			}{
				Token: token,
			},
		)
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

		result, err := json.Marshal(
			&struct {
				Name string `json:"name"`
			}{
				Name: name,
			},
		)
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

		var reqParams = struct {
			Name string `json:"name"`
		}{}
		err := json.NewDecoder(req.Body).Decode(&reqParams)
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

		err = users.UpdateUser(reqParams.Name, token)
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

//TODO: not implemented
// func DrawGacha() http.HandlerFunc {
// 	return func(resp http.ResponseWriter, req *http.Request) {

// 		var reqParams = struct {
// 			Times int `json:"times"`
// 		}{}
// 		err := json.NewDecoder(req.Body).Decode(&reqParams)
// 		if err != nil {
// 			params := view.ErrorViewParams{
// 				Error:      err,
// 				StatusCode: http.StatusInternalServerError,
// 				Message:    "Error unmarshalling the request",
// 			}
// 			view.ErrorView(resp, params)
// 			return
// 		}

// 		//fetch token and extract userid
// 		token := req.Header.Get("X-Auth-Token")

// 		results, err := gacha.DrawGacha(reqParams.Times, token)
// 		if err != nil {
// 			params := view.ErrorViewParams{
// 				Error:      err,
// 				StatusCode: http.StatusInternalServerError,
// 				Message:    "Error executing model process",
// 			}
// 			view.ErrorView(resp, params)
// 			return
// 		}

// 		gachaResults, err := json.Marshal(&results)
// 		if err != nil {
// 			params := view.ErrorViewParams{
// 				Error:      err,
// 				StatusCode: http.StatusInternalServerError,
// 				Message:    "Error marshalling data",
// 			}
// 			view.ErrorView(resp, params)
// 			return
// 		}

// 		view.SuccessView(resp, gachaResults)
// 	}
// }

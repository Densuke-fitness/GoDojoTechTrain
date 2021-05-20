package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Densuke-fitness/GoDojoTechTrain/service/character"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/gacha"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/tokenService"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/tokenService/auth"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/users"
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
		userId, err := users.CreateUser(reqParams.Name)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error executing model process",
			}
			view.ErrorView(resp, params)
			return
		}

		token, err := auth.CreateToken(userId)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error creating token",
			}
			view.ErrorView(resp, params)
			return
		}

		resParams := view.CreateUserRes{Token: token}

		result, err := json.Marshal(resParams)
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
		decodedtoken, err := auth.DecodeToken(token)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error decoding token",
			}
			view.ErrorView(resp, params)
			return
		}

		userId := tokenService.ExtractFieldFromToken(tokenService.USER_ID, decodedtoken)
		if userId == nil {
			params := view.ErrorViewParams{
				Error:      nil,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error Invaild token",
			}
			view.ErrorView(resp, params)
			return
		}

		name, err := users.GetUser(userId.(int))
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error executing model process",
			}
			view.ErrorView(resp, params)
			return
		}
		resParams := view.GetUserRes{Name: name}

		result, err := json.Marshal(resParams)
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

		//fetch token and extract userid
		token := req.Header.Get("X-Auth-Token")
		decodedtoken, err := auth.DecodeToken(token)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error decoding token",
			}
			view.ErrorView(resp, params)
			return
		}

		var reqParams = struct {
			Name string `json:"name"`
		}{}
		err = json.NewDecoder(req.Body).Decode(&reqParams)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error unmarshalling the request",
			}
			view.ErrorView(resp, params)
			return
		}

		userId := tokenService.ExtractFieldFromToken(tokenService.USER_ID, decodedtoken)
		if userId == nil {
			params := view.ErrorViewParams{
				Error:      nil,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error Invaild token",
			}
			view.ErrorView(resp, params)
			return
		}

		err = users.UpdateUser(reqParams.Name, userId.(int))
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

func DrawGacha() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {

		var reqParams = struct {
			Times int `json:"times"`
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

		if reqParams.Times <= 0 {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusBadRequest,
				Message:    "The value of times must be at least 1.",
			}
			view.ErrorView(resp, params)
			return
		}

		//fetch token and extract userid
		token := req.Header.Get("X-Auth-Token")
		decodedtoken, err := auth.DecodeToken(token)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error decoding token",
			}
			view.ErrorView(resp, params)
			return
		}

		userId := tokenService.ExtractFieldFromToken(tokenService.USER_ID, decodedtoken)
		if userId == nil {
			params := view.ErrorViewParams{
				Error:      nil,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error Invaild token",
			}
			view.ErrorView(resp, params)
			return
		}

		gachaResults, err := gacha.DrawGacha(reqParams.Times, userId.(int))
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error executing model process",
			}
			view.ErrorView(resp, params)
			return
		}

		paramsToViewList := []view.DrawGachaRes{}

		for _, val := range gachaResults {
			resParams := view.DrawGachaRes{
				CharacterId: strconv.Itoa(val.Id),
				Name:        val.Name,
			}

			paramsToViewList = append(paramsToViewList, resParams)
		}

		result, err := json.Marshal(&paramsToViewList)
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

func GetCharacterList() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {

		token := req.Header.Get("X-Auth-Token")
		decodedtoken, err := auth.DecodeToken(token)
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error decoding token",
			}
			view.ErrorView(resp, params)
			return
		}

		userId := tokenService.ExtractFieldFromToken(tokenService.USER_ID, decodedtoken)
		if userId == nil {
			params := view.ErrorViewParams{
				Error:      nil,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error Invaild token",
			}
			view.ErrorView(resp, params)
			return
		}

		characters, err := character.GetCharacterList(userId.(int))
		if err != nil {
			params := view.ErrorViewParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "Error executing model process",
			}
			view.ErrorView(resp, params)
			return
		}

		paramsToViewList := []view.GetCharacterListRes{}

		for _, val := range characters {
			resParams := view.GetCharacterListRes{
				CharacterSeq: strconv.Itoa(val.CharacterSeq),
				CharacterId:  strconv.Itoa(val.Id),
				Name:         val.Name,
			}

			paramsToViewList = append(paramsToViewList, resParams)
		}

		result, err := json.Marshal(&paramsToViewList)
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

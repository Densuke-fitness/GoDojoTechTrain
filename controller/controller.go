package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Densuke-fitness/GoDojoTechTrain/service/character"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/gacha"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/tokenService/auth"
	"github.com/Densuke-fitness/GoDojoTechTrain/service/users"
	"github.com/Densuke-fitness/GoDojoTechTrain/middleware"
)

func CreateUser() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		//Structure to be stored when a reqParams is received from a user

		reqParams := ReqParamsCreateUser{}

		err := json.NewDecoder(req.Body).Decode(&reqParams)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusBadRequest,
			}
			middleware.Error(resp, params)
			return
		}

		err = ValidateForReqParams(reqParams)

		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusBadRequest,
			}
			middleware.Error(resp, params)
			return
		}

		// Passing values to the model and executing the process
		userId, err := users.CreateUser(reqParams.Name)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}

		token, err := auth.CreateToken(userId)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}

		resParams := middleware.CreateUserRes{Token: token}

		result, err := json.Marshal(resParams)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}

		middleware.Success(resp, result)
	}
}

func GetUser() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {

		token := req.Header.Get("X-Auth-Token")
		userClaims, err := auth.DecodeToken(token)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}

		name, err := users.GetUser(userClaims.UserId)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}
		resParams := middleware.GetUserRes{Name: name}

		result, err := json.Marshal(resParams)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}

		middleware.Success(resp, result)
	}
}

func UpdateUser() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {

		//fetch token and extract userid
		token := req.Header.Get("X-Auth-Token")
		userClaims, err := auth.DecodeToken(token)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}

		reqParams := ReqParamsUpdateUser{}
		err = json.NewDecoder(req.Body).Decode(&reqParams)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}

		err = ValidateForReqParams(reqParams)

		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusBadRequest,
			}
			middleware.Error(resp, params)
			return
		}

		err = users.UpdateUser(reqParams.Name, userClaims.UserId)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}
		middleware.Success(resp, nil)
	}
}

func DrawGacha() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {

		reqParams := ReqParamsDrawGacha{}

		err := json.NewDecoder(req.Body).Decode(&reqParams)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}

		err = ValidateForReqParams(reqParams)

		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusBadRequest,
			}
			middleware.Error(resp, params)
			return
		}

		//fetch token and extract userId
		token := req.Header.Get("X-Auth-Token")
		userClaims, err := auth.DecodeToken(token)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}

		gachaResults, err := gacha.DrawGacha(reqParams.Times, userClaims.UserId)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}

		paramsToList := []middleware.DrawGachaRes{}

		for _, val := range gachaResults {
			resParams := middleware.DrawGachaRes{
				CharacterId: strconv.Itoa(val.Id),
				Name:        val.Name,
			}

			paramsToList = append(paramsToList, resParams)
		}

		result, err := json.Marshal(&paramsToList)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}

		middleware.Success(resp, result)
	}
}

func GetCharacterList() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {

		token := req.Header.Get("X-Auth-Token")
		userClaims, err := auth.DecodeToken(token)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}

		characters, err := character.GetCharacterList(userClaims.UserId)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}

		paramsToList := []middleware.GetCharacterListRes{}

		for _, val := range characters {
			resParams := middleware.GetCharacterListRes{
				CharacterSeq: strconv.Itoa(val.CharacterSeq),
				CharacterId:  strconv.Itoa(val.Id),
				Name:         val.Name,
			}

			paramsToList = append(paramsToList, resParams)
		}

		result, err := json.Marshal(&paramsToList)
		if err != nil {
			params := middleware.ErrorParams{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
			middleware.Error(resp, params)
			return
		}

		middleware.Success(resp, result)
	}
}

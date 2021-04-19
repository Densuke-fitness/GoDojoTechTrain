package users

import (
	"encoding/json"
	"net/http"

	"github.com/Densuke-fitness/GoDojoTechTrain/model/auth"
	logger "github.com/sirupsen/logrus"
)

func UpdateUser(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	//Structure to be stored when a request is received from a user
	var request = struct {
		Name string `json:"name"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		logger.Warnf("Error json.NewDecoder(req.Body).Decode: %s", err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error™: "Error unmarshalling the request"`)) //nolint
		return
	}

	//fetch token and extract userid
	token := req.Header.Get("X-Auth-Token")
	decodedtoken, err := auth.DecodeToken(token)
	if err != nil {
		logger.Errorf("Error auth.DecodeToken: %s", err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error decoding token"`)) //nolint
		return
	}

	var userId int

	switch decodedtoken["user_id"] {
	case nil:
		logger.Warnf("Error decodedtoken['user_id']: %s", err)
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`"error": "Not found user_id"`)) //nolint
		return
	default:
		// extract userid
		userId = int(decodedtoken["user_id"].(float64))
	}

	//Update
	_, err = UpdateNameById(request.Name, userId)
	if err != nil {
		logger.Errorf("Error UpdateNameById: %s", err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error updating name"`)) //nolint
		return
	}

	resp.WriteHeader(http.StatusOK)
}

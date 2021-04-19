package users

import (
	"encoding/json"
	"net/http"

	"github.com/Densuke-fitness/GoDojoTechTrain/model/auth"
	logger "github.com/sirupsen/logrus"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	//decode token
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

	//search name by using id
	name, err := SelectNameById(userId)
	if err != nil {
		logger.Errorf("Error SelectNameById: %s", err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error scanning name"`)) //nolint
		return
	}

	response := struct {
		Name string `json:"name"`
	}{Name: name}

	result, err := json.Marshal(&response)
	if err != nil {
		logger.Errorf("Error json.Marshal: %s", err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error marshalling data"`)) //nolint
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result) //nolint
}

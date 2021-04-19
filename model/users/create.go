package users

import (
	"encoding/json"
	"net/http"

	"github.com/Densuke-fitness/GoDojoTechTrain/model/auth"
	logger "github.com/sirupsen/logrus"
)

func CreateUser(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	//Structure to be stored when a request is received from a user
	var request = struct {
		Name string `json:"name"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		logger.Warnf("Error json.NewDecoder(req.Body).Decode: %s", err)
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`"error": "Error unmarshalling the request"`)) //nolint
		return
	}

	//database prosess
	id, err := Insert(request.Name)
	if err != nil {
		logger.Errorf("Error Insert: %s", err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error inserting the name"`)) //nolint
		return
	}

	//jwt prosess
	token, err := auth.CreateToken(id)
	if err != nil {
		logger.Errorf("Error auth.CreateToken: %s", err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error creating the token"`)) //nolint
		return
	}

	//Structure to be stored when sending a response to a user
	response := struct {
		Token string `json:"token"`
	}{Token: token}
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

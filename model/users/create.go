package users

import (
	"encoding/json"
	"net/http"

	"github.com/Densuke-fitness/GoDojoTechTrain/model/auth"
)

func CreateUser(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	//Structure to be stored when a request is received from a user
	var request = struct {
		Name string `json:"name"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error unmarshalling the request"`)) //nolint
	}

	//database prosess
	id, err := Insert(request.Name)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error inserting the name"`)) //nolint
		return
	}

	//jwt prosess
	token, err := auth.CreateToken(int(id))
	if err != nil {
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
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error marshalling data"`)) //nolint
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(result) //nolint
	return
}

package users

import (
	"encoding/json"
	"net/http"

	"github.com/Densuke-fitness/GoDojoTechTrain/model/auth"
)

func UpdateUser(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	//Structure to be stored when a request is received from a user
	var request = struct {
		Name string `json:"name"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error™: "Error unmarshalling the request"`)) //nolint
	}

	//fetch token and extract userid
	token := req.Header.Get("X-Auth-Token")
	decodedtoken, err := auth.DecodeToken(token)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error decoding token"`)) //nolint
	}

	if decodedtoken["user_id"] == nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`"error": "Not found user_id"`)) //nolint
	}
	// extract userid
	userId := int(decodedtoken["user_id"].(float64))

	//Update
	_, err = UpdateNameById(request.Name, userId)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error updating name"`)) //nolint
	}

	resp.WriteHeader(http.StatusOK)
}

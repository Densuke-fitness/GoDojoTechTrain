package users

import (
	"encoding/json"
	"net/http"

	"github.com/Densuke-fitness/GoDojoTechTrain/model/auth"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	//decode token
	token := req.Header.Get("X-Auth-Token")
	decodedtoken, err := auth.DecodeToken(token)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error decoding token"`)) //nolint
	}

	// extract userid
	id := int(decodedtoken["id"].(float64))

	//search name by using id
	name, err := SelectNameById(id)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error scanning name"`)) //nolint
	}

	response := struct {
		Name string `json:"name"`
	}{Name: name}

	result, err := json.Marshal(&response)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error marshalling data"`)) //nolint
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result) //nolint
}

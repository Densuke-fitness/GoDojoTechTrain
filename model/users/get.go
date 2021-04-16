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
		//TODO: The return value is listed because it was flagged as an error by golangci-lint, but I expect there is a better way.
		_, _ = resp.Write([]byte(`"error": "Error decoding token"`))
	}

	// extract userid
	id := int(decodedtoken["id"].(float64))

	//search name by using id
	name, err := SelectNameById(id)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		_, _ = resp.Write([]byte(`"error": "Error scanning name"`))
	}

	response := struct {
		Name string `json:"name"`
	}{Name: name}

	result, err := json.Marshal(&response)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		_, _ = resp.Write([]byte(`"error": "Error marshalling data"`))
	}
	resp.WriteHeader(http.StatusOK)
	_, _ = resp.Write(result)
}

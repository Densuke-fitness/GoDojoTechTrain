package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Densuke-fitness/GoDojoTechTrain/auth"
	"github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"
)

type RequestCreate struct {
	Name string `json:"name"`
}

type ResponseCreate struct {
	Token string `json:"token"`
}

func CreateUser() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-type", "application/json")
		//Structure to be stored when a request is received from a user
		var request RequestCreate
		err := json.NewDecoder(req.Body).Decode(&request)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(`"error: "Error unmarshalling the request"`))
		}
		//Insert a name into a user table using sql
		dbConn := dbConnection.GetInstance()
		defer dbConn.Close()

		db := dbConn.GetConnection()

		const sql = "INSERT INTO users(name) VALUES (?)"
		//Save the name data (id is automatically generated)
		r, err := db.Exec(sql, request.Name)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(`"error: "Error inserting name"`))
		}

		id, err := r.LastInsertId()
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(`"error: "Error reading auto increment id"`))
		}

		token, err := auth.CreateToken(int(id))
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(`"error: "Error creating the token"`))
		}

		//Structure to be stored when sending a response to a user
		response := ResponseCreate{Token: token}
		result, err := json.Marshal(&response)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(`"error: "Error creating the token"`))
		}
		resp.WriteHeader(http.StatusOK)
		resp.Write(result)
	}
}

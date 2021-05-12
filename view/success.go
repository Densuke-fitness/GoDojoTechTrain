package view

import (
	"net/http"
)

func SuccessView(resp http.ResponseWriter, result []byte) {
	resp.Header().Set("Content-type", "application/json")

	resp.WriteHeader(http.StatusOK)
	resp.Write(result) //nolint
}

type CreateUserRes struct {
	Token string `json:"token"`
}

type GetUserRes struct {
	Name string `json:"name"`
}

type DrawGachaRes struct {
	CharacterId string `json:"characterId"`
	Name        string `json:"name"`
}

type GetCharacterListRes struct {
	CharacterSeq string `json:"userCharacterID"`
	CharacterId  string `json:"characterId"`
	Name         string `json:"name"`
}

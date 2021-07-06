package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Densuke-fitness/GoDojoTechTrain/controller"
	"github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"
	"github.com/gorilla/mux"
)

func TestE2E(t *testing.T) {

	var router = mux.NewRouter()
	router.HandleFunc("/user/create", controller.CreateUser()).Methods("POST")
	router.HandleFunc("/user/get", controller.GetUser()).Methods("GET")
	router.HandleFunc("/user/update", controller.UpdateUser()).Methods("PUT")
	router.HandleFunc("/gacha/draw", controller.DrawGacha()).Methods("POST")
	router.HandleFunc("/character/list", controller.GetCharacterList()).Methods("GET")

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	defer dbConnection.GetInstance().Close()

	//Test CreateUser
	var jsonStr = []byte(`{"name":"TestUser"}`)
	req, err := http.NewRequest(http.MethodPost, testServer.URL+"/user/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		if err != nil {
			t.Errorf("req error: %s", err.Error())
		}
	}
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		if err != nil {
			t.Errorf("resp error: %s", err.Error())
		}
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))

	tmp := struct {
		Token string `json:"token"`
	}{}
	if err := json.Unmarshal(respBody, &tmp); err != nil {
		fmt.Println(err)
		return
	}
	//Test GetUser
	req, err = http.NewRequest(http.MethodGet, testServer.URL+"/user/get", nil)
	if err != nil {
		if err != nil {
			t.Errorf("req error: %s", err.Error())
		}
	}
	req.Header.Set("X-Auth-Token", tmp.Token)

	client = new(http.Client)
	resp, err = client.Do(req)
	if err != nil {
		if err != nil {
			t.Errorf("resp error: %s", err.Error())
		}
	}
	respBody, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
}

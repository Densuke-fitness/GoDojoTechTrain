package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// // testUserData
const (
	name = "TestUser"
)

// POST /user/create に対するtest
// 正常なパラメータでリクエストを行う
func TestContoller(t *testing.T) {

	//testing CreateUser
	testCreateRequest := struct {
		Name string `json:"name"`
	}{Name: name}

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(&testCreateRequest)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRequest(http.MethodPost, "/", &body)
	w := httptest.NewRecorder()
	CreateUser()(w, r)

	rw := w.Result()

	defer r.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}

	_, err = ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	//TODO: testing GetUser

	//TODO: testing UpdateUser

}

package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/user/create", nil)
	CreateUser()(w, r)
	rw := w.Result()
	defer r.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	const expected = "test"
	if s := string(b); s != expected {
		t.Fatalf("unexpected response: %s", s)
	}

}

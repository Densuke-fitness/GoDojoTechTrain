package view

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessView(t *testing.T) {

	//testing ErrorView
	b := []byte("test")
	w := httptest.NewRecorder()
	SuccessView(w, b)

	rw := w.Result()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}

	got, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	const expected = "test"
	if string(got) != expected {
		t.Errorf(`Error DecodeToken: %v but want %q`, got, expected)
	}

}

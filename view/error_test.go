package view

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorView(t *testing.T) {

	testParams := ErrorViewParams{
		nil,
		http.StatusInternalServerError,
		"test",
	}

	w := httptest.NewRecorder()
	ErrorView(w, testParams)

	rw := w.Result()

	if rw.StatusCode != testParams.StatusCode {
		t.Fatal("unexpected status code")
	}

	got, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	expected := fmt.Sprintf(`"error": "%s"`, testParams.Message)
	if string(got) != expected {
		t.Errorf(`Error DecodeToken: %v but want %q`, got, expected)
	}

}

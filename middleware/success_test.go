package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccess(t *testing.T) {

	//errorPattern
	tests := []struct {
		description string
		responseVal []byte
		wanted      []byte
	}{{
		description: "Test StatusOK",
		responseVal: []byte(`"test"`),
		wanted:      []byte(`"test"`),
	}}

	for id, tt := range tests {

		testCaseName := fmt.Sprintf("%v: %v", id+1, tt.description)

		t.Run(testCaseName, func(t *testing.T) {

			//TODO: テストケースをそれぞれのAPIの成功ケースについて実装するかどうか議論する
			w := httptest.NewRecorder()
			Success(w, tt.responseVal)

			rw := w.Result()

			if rw.StatusCode != http.StatusOK {
				t.Fatal("unexpected status code")
			}

			got, err := ioutil.ReadAll(rw.Body)
			if err != nil {
				t.Fatal("unexpected error")
			}

			if string(got) != string(tt.wanted) {
				t.Errorf(`Error Success: %v but want %q`, got, tt.wanted)
			}

		})
	}
}

package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestError(t *testing.T) {

	//errorPattern
	tests := []struct {
		description string
		statusCode  int
		wantedError []byte
	}{{
		description: "Test IntervalServerError",
		statusCode:  http.StatusInternalServerError,
		wantedError: []byte(`"error": "IntervalServerError"`),
	}, {
		description: "Test IntervalServerError",
		statusCode:  http.StatusBadRequest,
		wantedError: []byte(`"error": "StatusBadRequest"`),
	}, {
		description: "Test GeneralError",
		statusCode:  http.StatusBadGateway,
		wantedError: []byte(`"error": "GeneralError"`),
	}}

	for id, tt := range tests {
		//ParallelTest
		tt := tt

		//Define an anonymous helper function
		testParams := func(StatusCode int) ErrorParams {
			return ErrorParams{nil, StatusCode}
		}(tt.statusCode)

		testCaseName := fmt.Sprintf("%v: %v", id+1, tt.description)

		t.Run(testCaseName, func(t *testing.T) {

			//ParallelTest
			t.Parallel()

			//test :Error
			w := httptest.NewRecorder()
			Error(w, testParams)

			rw := w.Result()

			if rw.StatusCode != testParams.StatusCode {
				t.Fatal("unexpected status code")
			}

			got, err := ioutil.ReadAll(rw.Body)
			if err != nil {
				t.Fatal("unexpected error")
			}

			//Converted to string because byte array is not comparable
			if string(got) != string(tt.wantedError) {
				t.Error()
			}
		})
	}
}

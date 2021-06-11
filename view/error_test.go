package view

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorView(t *testing.T) {

	//errorPattern
	tests := []struct {
		description    string
		testStatusCode int
		wantedError    []byte
	}{{
		description:    "Test IntervalServerError",
		testStatusCode: http.StatusInternalServerError,
		wantedError:    []byte(`"error": "IntervalServerError"`),
	}, {
		description:    "Test IntervalServerError",
		testStatusCode: http.StatusBadRequest,
		wantedError:    []byte(`"error": "StatusBadRequest"`),
	}, {
		description:    "Test GeneralError",
		testStatusCode: http.StatusBadGateway,
		wantedError:    []byte(`"error": "GeneralError"`),
	}}

	for id, tt := range tests {
		//ParallelTest
		tt := tt

		//Define an anonymous helper function
		testParams := func(StatusCode int) ErrorViewParams {
			return ErrorViewParams{nil, StatusCode}
		}(tt.testStatusCode)

		testCaseName := fmt.Sprintf("%v: %v", id+1, tt.description)

		t.Run(testCaseName, func(t *testing.T) {

			//ParallelTest
			t.Parallel()

			//test :ErrorView
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

			//Converted to string because byte array is not comparable
			if string(got) != string(tt.wantedError) {
				t.Error()
			}
		})
	}
}

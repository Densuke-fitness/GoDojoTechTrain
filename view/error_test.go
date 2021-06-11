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
		id             int
		testStatusCode int
		wantedError    []byte
	}{{
		id:             1,
		testStatusCode: http.StatusInternalServerError,
		wantedError:    []byte(`"error": "IntervalServerError"`),
	}, {
		id:             2,
		testStatusCode: http.StatusBadRequest,
		wantedError:    []byte(`"error": "StatusBadRequest"`),
	}, {
		id:             3,
		testStatusCode: http.StatusBadGateway,
		wantedError:    []byte(`"error": "GeneralError"`),
	}}

	for _, tt := range tests {
		//ParallelTest
		tt := tt

		//Define an anonymous helper function
		testParams := func(StatusCode int) ErrorViewParams {
			return ErrorViewParams{nil, StatusCode}
		}(tt.testStatusCode)

		testName := fmt.Sprintf("number:%v", tt.id)

		t.Run(testName, func(t *testing.T) {

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

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
		id          int
		testParams  ErrorViewParams
		wantedError []byte
	}{{
		id:          1,
		testParams:  ErrorViewParams{nil, http.StatusInternalServerError},
		wantedError: []byte(`"error": "IntervalServerError"`),
	}, {
		id:          2,
		testParams:  ErrorViewParams{nil, http.StatusBadRequest},
		wantedError: []byte(`"error": "StatusBadRequest"`),
	}, {
		id:          3,
		testParams:  ErrorViewParams{nil, http.StatusBadGateway},
		wantedError: []byte(`"error": "GeneralError"`),
	}}

	for _, tt := range tests {
		//ParallelTest
		tt := tt

		testName := fmt.Sprintf("number:%v", tt.id)

		t.Run(testName, func(t *testing.T) {

			//ParallelTest
			t.Parallel()

			//test :ErrorView
			w := httptest.NewRecorder()
			ErrorView(w, tt.testParams)

			rw := w.Result()

			if rw.StatusCode != tt.testParams.StatusCode {
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

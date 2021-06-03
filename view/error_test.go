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
		id              int
		testParams      ErrorViewParams
		wantedErrorCode string
	}{{
		id:              1,
		testParams:      ErrorViewParams{nil, http.StatusInternalServerError},
		wantedErrorCode: "IntervalServerError",
	}, {
		id:              2,
		testParams:      ErrorViewParams{nil, http.StatusBadRequest},
		wantedErrorCode: "StatusBadRequest",
	}, {
		id:              3,
		testParams:      ErrorViewParams{nil, http.StatusBadGateway},
		wantedErrorCode: "GeneralError",
	}}

	for _, tt := range tests {
		testName := fmt.Sprintf("number:%v", tt.id)

		t.Run(testName, func(t *testing.T) {
			//test CreateUser
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

			expected := fmt.Sprintf(`"error": "%s"`, tt.wantedErrorCode)
			if string(got) != expected {
				t.Error()
			}
		})
	}
}

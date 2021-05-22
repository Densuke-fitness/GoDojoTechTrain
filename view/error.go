package view

import (
	"fmt"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func ErrorView(resp http.ResponseWriter, params ErrorViewParams) {
	resp.Header().Set("Content-type", "application/json")

	errMsg := generateError(params.StatusCode)
	logger.Errorf("%s : %s", errMsg, params.Error)
	resp.WriteHeader(params.StatusCode)
	resp.Write([]byte(fmt.Sprintf(`"error": "%s"`, errMsg))) //nolint
}

func generateError(StatusCode int) (errMsg string) {

	switch StatusCode {
	case http.StatusInternalServerError:
		errMsg = "IntervalServerError"
	case http.StatusBadRequest:
		errMsg = "StatusBadRequest"
	default:
		errMsg = "GeneralError"
	}

	return
}

type ErrorViewParams struct {
	Error      error
	StatusCode int
	Message    string
}

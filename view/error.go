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

func generateError(StatusCode int) string {

	switch StatusCode {
	case http.StatusInternalServerError:
		return "IntervalServerError"
	case http.StatusBadRequest:
		return "StatusBadRequest"
	default:
		return "GeneralError"
	}
}

type ErrorViewParams struct {
	Error      error
	StatusCode int
}

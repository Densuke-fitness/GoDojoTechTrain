package view

import (
	"fmt"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

type ErrorViewParams struct {
	Error      error
	StatusCode int
	Message    string
}

func ErrorView(resp http.ResponseWriter, params ErrorViewParams) {
	resp.Header().Set("Content-type", "application/json")

	logger.Errorf("Error : %s", params.Error)
	resp.WriteHeader(params.StatusCode)
	resp.Write([]byte(fmt.Sprintf(`"error": "%s"`, params.Message))) //nolint

}

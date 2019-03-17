package helper

import (
	"encoding/json"
	"gctest/errors"
	"gctest/logservice"
	"net/http"
)

type ErrorData struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}

func getErrorCode(err error) int {

	if err == errors.BadRequest {
		return http.StatusBadRequest
	} else if err == errors.ServerError {
		return http.StatusInternalServerError
	} else if err == errors.NotFound {
		return http.StatusNotFound
	} else if err == errors.GroupStageAlreadyFinished {
		return http.StatusConflict
	}

	return 500
}

func HandleError(endpoint string, err error, w http.ResponseWriter) {
	errorData := ErrorData{err.Error(), getErrorCode(err)}

	logservice.NewLogService(endpoint).Error(errorData)

	w.WriteHeader(errorData.Status)
	json.NewEncoder(w).Encode(&errorData)
}

package helper

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
)

func GenerateResponseError(w http.ResponseWriter, r *http.Request, code int, message string) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(model.ErrorResponse{Error: message})

	if err != nil {
		return err
	}

	return nil
}

func GenerateResponseSuccess(w http.ResponseWriter, r *http.Request, code int, successResponse model.SuccessResponse) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(successResponse)

	if err != nil {
		return err
	}

	return nil
}

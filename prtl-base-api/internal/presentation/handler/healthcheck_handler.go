package handler

import (
	"net/http"
	res "prtl-base-api/pkg/response"
)

type healthCheckResponse struct {
	Status string `json:"status"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	body := healthCheckResponse{
		Status: "OK",
	}
	res.SendJSONResponse(w, http.StatusOK, body)
}

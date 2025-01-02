package handler

import (
	"net/http"
	"encoding/json"
	"log/slog"
	res "prtl-base-api/pkg/response"
	AppService "prtl-base-api/internal/application/service"
	"prtl-base-api/internal/application/dto"
)

type AuthHandler struct {
	authService *AppService.AuthService
}

func NewAuthHandler(authService *AppService.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginReq dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		slog.Error("Invalid request")
		return
	}

	token, err := h.authService.AuthenticateUser(loginReq.Username, loginReq.Password)

	if err != nil {
		slog.Error(err.Error())
	}

	body := dto.LoginResponse{
		Token: token,
	}

	res.SendJSONResponse(w, http.StatusOK, body)
}

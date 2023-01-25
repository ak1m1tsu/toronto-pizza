package handlers

import (
	"net/http"

	"github.com/romankravchuk/toronto-pizza/internal/service"
)

type AuthHandler struct {
	svc service.IAuthService
}

func NewAuthHandler(svc service.IAuthService) *AuthHandler {
	return &AuthHandler{svc}
}

func (h *AuthHandler) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, map[string]any{"status": 200, "body": "Sign in"})
}
func (h *AuthHandler) HandleLogOut(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, map[string]any{"status": 200, "body": "Log out"})
}
func (h *AuthHandler) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, map[string]any{"status": 200, "body": "Sign up"})
}
func (h *AuthHandler) HandleRefreshToken(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, map[string]any{"status": 200, "body": "Refresh Token"})
}

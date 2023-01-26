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
	resp := NewApiResponse(http.StatusForbidden, map[string]any{})

	creds, err := NewCredetials(r.Body)
	if err != nil {
		resp.SetError(err)
		JSON(w, resp.Status, resp)
		return
	}

	if err := h.svc.ValidatePassword(creds.Phone, creds.Password); err != nil {
		resp.SetError(err)
		JSON(w, resp.Status, resp)
	}

	resp.Status = http.StatusInternalServerError

	access_token, err := h.svc.CreateToken(h.accessToken.ExpiresIn, user.ID, h.accessToken.PrivateKey)
	if err != nil {
		resp.SetError(err)
		JSON(writer, resp.Status, resp)
		return
	}

	refresh_token, err := h.svc.CreateToken(h.refreshToken.ExpiresIn, user.ID, h.refreshToken.PrivateKey)
	if err != nil {
		resp.SetError(err)
		JSON(writer, resp.Status, resp)
		return
	}
	
	resp.Status = http.StatusOK
	resp.Body = map[string]any{}
	JSON(w, resp.Status, resp)
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

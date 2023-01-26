package handlers

import (
	"net/http"

	"github.com/romankravchuk/toronto-pizza/internal/config"
	"github.com/romankravchuk/toronto-pizza/internal/router/handlers/models"
	"github.com/romankravchuk/toronto-pizza/internal/service"
)

type AuthHandler struct {
	svc          service.IAuthService
	accessToken  config.Token
	refreshToken config.Token
}

func NewAuthHandler(svc service.IAuthService, access, refresh config.Token) *AuthHandler {
	return &AuthHandler{svc: svc, accessToken: access, refreshToken: refresh}
}

func (h *AuthHandler) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	resp := NewApiResponse(http.StatusForbidden, map[string]any{})

	creds, err := models.NewCredetials(r.Body)
	if err != nil {
		resp.SetError(err)
		JSON(w, resp.Status, resp)
		return
	}

	if err := h.svc.ValidatePassword(creds.Phone, creds.Password); err != nil {
		resp.SetError(err)
		JSON(w, resp.Status, resp)
		return
	}

	resp.Status = http.StatusInternalServerError

	accessToken, err := h.svc.CreateToken(h.accessToken.ExpiresIn, creds.Phone, h.accessToken.PrivateKey)
	if err != nil {
		resp.SetError(err)
		JSON(w, resp.Status, resp)
		return
	}

	refreshToken, err := h.svc.CreateToken(h.refreshToken.ExpiresIn, creds.Phone, h.refreshToken.PrivateKey)
	if err != nil {
		resp.SetError(err)
		JSON(w, resp.Status, resp)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     AccessTokenHeader,
		Value:    accessToken,
		Domain:   "localhost",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   h.accessToken.MaxAge * 60,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     RefreshTokenHeader,
		Value:    refreshToken,
		Domain:   "localhost",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   h.refreshToken.MaxAge * 60,
	})

	resp = NewApiResponse(http.StatusOK, map[string]any{"access_token": accessToken})
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

package handlers

import (
	"fmt"
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
	creds, err := models.NewCredetials(r.Body)
	if err != nil {
		JSON(w, http.StatusForbidden, nil, "", err)
		return
	}

	if err := h.svc.ValidatePassword(creds.Phone, creds.Password); err != nil {
		JSON(w, http.StatusUnauthorized, nil, "", err)
		return
	}

	accessToken, err := h.svc.CreateToken(h.accessToken.ExpiresIn, creds.Phone, h.accessToken.PrivateKey)
	if err != nil {
		JSON(w, http.StatusInternalServerError, nil, "", err)
		return
	}

	refreshToken, err := h.svc.CreateToken(h.refreshToken.ExpiresIn, creds.Phone, h.refreshToken.PrivateKey)
	if err != nil {
		JSON(w, http.StatusInternalServerError, nil, "", err)
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

	data := map[string]interface{}{"access_token": accessToken}
	JSON(w, http.StatusOK, data, "Authorized.", nil)
}

func (h *AuthHandler) HandleLogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     AccessTokenHeader,
		Domain:   "localhost",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     RefreshTokenHeader,
		Domain:   "localhost",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})
	JSON(w, http.StatusOK, nil, "Successfuly logged out.", nil)
}

func (h *AuthHandler) HandleRefreshToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(RefreshTokenHeader)
	if err != nil {
		JSON(w, http.StatusForbidden, nil, "Cookie do not set.", ErrForbidden)
		return
	}

	sub, err := h.svc.ValidateToken(cookie.Value, h.refreshToken.PublicKey)
	if err != nil {
		JSON(w, http.StatusForbidden, nil, "Invalid token.", ErrForbidden)
		return
	}

	user, err := h.svc.GetUserByPhone(r.Context(), fmt.Sprint(sub))
	if err != nil {
		JSON(w, http.StatusForbidden, nil, "Invalid token payload.", ErrForbidden)
		return
	}

	accessToken, err := h.svc.CreateToken(h.accessToken.ExpiresIn, user.ID, h.accessToken.PrivateKey)
	if err != nil {
		JSON(w, http.StatusForbidden, nil, "Can not to create access token.", err)
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

	data := map[string]interface{}{"access_token": accessToken}
	JSON(w, http.StatusOK, data, "Successfuly refreshed.", nil)
}

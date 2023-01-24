package handlers

import "net/http"

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
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

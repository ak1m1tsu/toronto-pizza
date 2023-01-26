package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type contextKey int

const (
	AuthorizationHeader string     = "Authorization"
	AccessTokenHeader   string     = "X-Access-Token"
	RefreshTokenHeader  string     = "X-Refresh-Token"
	ProductKey          contextKey = iota
)

type ApiResponse struct {
	Status int            `json:"status"`
	Body   map[string]any `json:"body"`
}

func NewApiResponse(status int, body map[string]any) *ApiResponse {
	return &ApiResponse{status, body}
}

func (r *ApiResponse) SetError(err error) {
	r.Body = map[string]any{"error": err.Error()}
}

func JSON(w http.ResponseWriter, status int, v interface{}) {
	buffer := &bytes.Buffer{}
	if err := json.NewEncoder(buffer).Encode(v); err != nil {
		http.Error(w, ErrInternalServer.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(buffer.Bytes())
}

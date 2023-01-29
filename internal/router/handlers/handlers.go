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
	keyProduct          contextKey = iota
	keyUserDTO
)

type ApiResponse struct {
	Status  int                    `json:"status"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
	Error   string                 `json:"error"`
}

func JSON(w http.ResponseWriter, status int, data map[string]interface{}, message string, err error) {
	resp := ApiResponse{
		Status:  status,
		Data:    data,
		Message: message,
	}

	if err != nil {
		resp.Error = err.Error()
	}

	buffer := &bytes.Buffer{}
	if err := json.NewEncoder(buffer).Encode(resp); err != nil {
		http.Error(w, ErrInternalServer.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(buffer.Bytes())
}

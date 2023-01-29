package models

import (
	"io"
	"encoding/json"
)

type Credentials struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}

func NewCredetials(body io.Reader) (*Credentials, error) {
	var creds *Credentials
	if err := json.NewDecoder(body).Decode(&creds); err != nil {
		return nil, err
	}
	return creds, nil
}

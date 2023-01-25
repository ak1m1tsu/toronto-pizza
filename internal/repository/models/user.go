package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                string `bson:"_id,omitempty"`
	Name              string `bson:"name"`
	Phone             string `bson:"phone"`
	EncryptedPassword string `bson:"encryptedPassword"`
}

func NewUser(name, phone, password string) (*User, error) {
	epwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:              name,
		Phone:             phone,
		EncryptedPassword: string(epwd),
	}, nil
}

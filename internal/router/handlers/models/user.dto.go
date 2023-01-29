package models

type UserDTO struct {
	ID string    `json:"id"`
	Name string  `json:"name"`
	Phone string `json:"phone"`
}

func NewUserDTO(id, name, phone string) *UserDTO {
	return &UserDTO{
		ID: id,
		Name: name,
		Phone: phone,
	}
}

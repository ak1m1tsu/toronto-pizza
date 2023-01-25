package models

type Category struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
}

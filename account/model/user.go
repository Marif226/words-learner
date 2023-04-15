package model

import "github.com/google/uuid"

// User defines domain model and its json and db dependecies
type User struct {
	UID 	uuid.UUID	`db:"uid" json:"uid"`
	Email		string	`db:"email" json:"email"`
	Password	string	`db:"password" json:"password"`
	Name		string	`db:"name" json:"name"`
	ImageURL	string	`db:"image_url" json:"imageUrl"`
	Website		string	`db:"website" json:"website"`
}
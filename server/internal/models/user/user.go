package models

import (
	"errors"
	"time"
)

type User struct {
	ID        string
	Username  string
	FirstName string
	LastName  string
	Email     string
	ImageURL  string
	Created   time.Time
	Updated   time.Time
}

func NewUser(id, username, firstName, lastName, email, imageURL string) *User {
	return &User{
		ID:        id,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		ImageURL:  imageURL,
		Created:   time.Now(),
	}
}

func (user *User) Validate() error {
	if user == nil {
		return errors.New("user validation failed")
	}
	if user.Username == "" {
		return errors.New("user validation failed")
	}
	return nil
}

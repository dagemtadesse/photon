package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var Validate = validator.New()

type User struct {
	Id       uuid.UUID `json:"id" `
	Email    string    `json:"email" validate:"required,email,max=100"`
	Password string    `json:"password" validate:"required,min=8,max=100"`
}

func (user *User) HashPassword() error {
	// hash the user password
	passwordSlice := []byte(user.Password)
	cost := bcrypt.DefaultCost

	// compute the hash
	hash, err := bcrypt.GenerateFromPassword(passwordSlice, cost)

	if err != nil {
		return err // halt if unable to hash the password
	}
	// change the plain password  on the user struct to
	// the hashed password
	user.Password = string(hash)

	return nil
}

func (user *User) CompareHashAndPassword(other *User) error {
	//hashed password and plain passwords
	hash := []byte(other.Password)
	plainPassword := []byte(user.Password)

	// compare the hash and the plain passwords
	return bcrypt.CompareHashAndPassword(hash, plainPassword)
}

func (user *User) Authenticate() error {
	return nil
}

package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Credential struct {
	Id       uuid.UUID `json:"id" `
	Email    string    `json:"email" validate:"required,email,max=100"`
	Password string    `json:"password" validate:"required,min=8,max=100"`
}

func (user *Credential) HashPassword() error {
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

func (user *Credential) CompareHashAndPassword(other *Credential) error {
	//hashed password and plain passwords
	hash := []byte(other.Password)
	plainPassword := []byte(user.Password)

	// compare the hash and the plain passwords
	return bcrypt.CompareHashAndPassword(hash, plainPassword)
}

func (user *Credential) CreateSessionId() string {
	// todo create session id
	return user.Id.String()
}

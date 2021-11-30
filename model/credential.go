package model

import (
	"github.com/google/uuid"
)

type Credential struct {
	Id       uuid.UUID `json:"id" `
	Email    string    `json:"email" validate:"required,email,max=100"`
	Password string    `json:"password" validate:"required,min=8,max=100"`
}

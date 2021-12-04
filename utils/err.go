package utils

import (
	"net/http"
)

type GenericErr struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func new(err, msg string) *GenericErr {
	return &GenericErr{
		Type:    err,
		Message: msg,
	}
}

var (
	ErrInvalidData            = new("invalida data", "some fields does not confirm with thier type")
	ErrInvalidDataFormat      = new("unsuported data format", "unable to parse equest body")
	ErrEmailAlreadyRegistered = new("database operation error", "provided email already exists")
	ErrInternalServerError    = new(http.StatusText(http.StatusInternalServerError), "")
)

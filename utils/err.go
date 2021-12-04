package utils

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidData            = errors.New("invalid Data provided")
	ErrInvalidDataFormat      = errors.New("unsuported data format")
	ErrEmailAlreadyRegistered = errors.New("email already taken")
	ErrInternalServerError    = errors.New(http.StatusText(http.StatusInternalServerError))
)

package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type GenericErrorMsg struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

var ErrEmailAlreadyExists = &GenericErrorMsg{
	Error:   "Database Error",
	Message: "Email already taken",
}

var ErrUnableToParseReqBody = &GenericErrorMsg{
	Error:   "Invalid Request Body",
	Message: "unable to parse json request",
}

func ErrInvalidDataForFields(err validator.ValidationErrors) *GenericErrorMsg {
	fields := make([]string, 0)
	for _, err := range err {
		fields = append(fields, err.Field())
	}

	return &GenericErrorMsg{
		Error:   "Validation Error",
		Message: fmt.Sprintf("Invalid data for fields [%v]", strings.Join(fields, ",")),
	}
}

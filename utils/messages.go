package utils

import "github.com/go-playground/validator/v10"

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "this field is required"
	case "email":
		return "enter valid email"
	case "alpha":
		return "only alpha characters allowed"
	}
	return "unknown error"
}

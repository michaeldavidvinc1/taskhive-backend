package utils

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate = validator.New()

func ValidationError(err error) string {
	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, err.Error())
	}
	return strings.Join(errors, ", ")
}

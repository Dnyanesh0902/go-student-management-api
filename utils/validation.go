package utils

import (
	"github.com/go-playground/validator/v10"
)

func GetValidationErrors(err error) []string {

	var errors []string

	for _, err := range err.(validator.ValidationErrors) {

		switch err.Field() {

		case "Name":
			errors = append(errors, "Name is required")

		case "Email":
			if err.Tag() == "required" {
				errors = append(errors, "Email is required")
			} else if err.Tag() == "email" {
				errors = append(errors, "Invalid email format")
			}

		case "Age":
			errors = append(errors, "Age must be greater than 0")

		case "Course":
			errors = append(errors, "Course is required")

		case "City":
			errors = append(errors, "City is required")
		}
	}

	return errors
}
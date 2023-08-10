package handlers

import "github.com/go-playground/validator/v10"

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "max":
		return "Should be less than " + fe.Param()
	case "min":
		return "Should be greater than " + fe.Param()
	}
	return "Unkonwn error"
}

package validators

import (
	"errors"
	"fmt"
	"simple-crud/models"

	"github.com/go-playground/validator/v10"
)

type UserValidator struct {
	validate *validator.Validate
}

func NewUserValidator() *UserValidator {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return &UserValidator{validate: validate}
}

func (userValidator *UserValidator) ValidateRequiredField(u models.User) ([]string, bool) {
	isValid := true
	if err := userValidator.validate.Struct(u); err != nil {
		isValid = false
		var validateErrs validator.ValidationErrors
		var messages []string
		if errors.As(err, &validateErrs) {
			messages = make([]string, len(validateErrs))

			for i, e := range validateErrs {
				fieldRequiredMessage := fmt.Sprintf("Field %s is required", e.Field())
				messages[i] = fieldRequiredMessage
			}
		}

		return messages, isValid
	}

	return []string{"Invalid body"}, isValid
}

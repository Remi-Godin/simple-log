package validators

import (
	"github.com/Remi-Godin/simple-log/internal/utils/validation"
)

type PasswordValidator struct {
	Validators    []validation.Validator
	FieldId       string
	FieldName     string
	FieldType     string
	FieldNameText string
	FieldValue    string
	Required      bool
	Valid         bool
	Invalid       bool
	Err           string
	Links         map[string]string
}

func NewPasswordValidator() PasswordValidator {
	var valList []validation.Validator
	valList = append(valList, validation.NewPasswordStrengthValidator(60))
	return PasswordValidator{
		Validators:    valList,
		FieldType:     "password",
		FieldName:     "password",
		FieldNameText: "Password",
		FieldId:       "password-input-field",
		Required:      true,
		Links:         make(map[string]string),
	}
}

func (nv PasswordValidator) GetValidators() []validation.Validator {
	return nv.Validators
}

func (nv PasswordValidator) GetFieldValue() string {
	return nv.FieldValue
}

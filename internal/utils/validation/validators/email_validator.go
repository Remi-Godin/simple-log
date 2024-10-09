package validators

import (
	"github.com/Remi-Godin/simple-log/internal/utils/validation"
)

type EmailValidator struct {
	Validators    []validation.Validator
	FieldId       string
	FieldName     string
	FieldType     string
	FieldNameText string
	FieldValue    string
	Valid         bool
	Invalid       bool
	Err           string
	Links         map[string]string
}

func NewEmailValidator() EmailValidator {
	var valList []validation.Validator
	valList = append(valList, validation.NewEmailFormatValidator())
	valList = append(valList, validation.NewEmailExistValidator())
	return EmailValidator{
		Validators:    valList,
		FieldType:     "email",
		FieldName:     "email",
		FieldNameText: "Email",
		FieldId:       "email-input-field",
		Links:         make(map[string]string),
	}
}

func (nv EmailValidator) GetValidators() []validation.Validator {
	return nv.Validators
}

func (nv EmailValidator) GetFieldValue() string {
	return nv.FieldValue
}

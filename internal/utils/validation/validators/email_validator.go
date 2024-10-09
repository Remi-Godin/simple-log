package validators

import (
	"github.com/Remi-Godin/simple-log/internal/utils/validation"
)

type emailValidator struct {
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

func NewEmailValidator() emailValidator {
	var valList []validation.Validator
	valList = append(valList, validation.NewEmailFormatValidator())
	valList = append(valList, validation.NewEmailExistValidator())
	return emailValidator{
		Validators: valList,
		Links:      make(map[string]string),
	}
}

func (nv emailValidator) GetValidators() []validation.Validator {
	return nv.Validators
}

func (nv emailValidator) GetFieldValue() string {
	return nv.FieldValue
}

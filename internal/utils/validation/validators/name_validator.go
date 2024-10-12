package validators

import (
	"github.com/Remi-Godin/simple-log/internal/utils/validation"
)

type NameValidator struct {
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

func NewNameValidator() NameValidator {
	var valList []validation.Validator
	valList = append(valList, validation.NewTextLengthValidator(2, 32))
	return NameValidator{
		FieldType:  "text",
		Validators: valList,
		Required:   true,
		Links:      make(map[string]string),
	}
}

func (nv NameValidator) GetValidators() []validation.Validator {
	return nv.Validators
}

func (nv NameValidator) GetFieldValue() string {
	return nv.FieldValue
}

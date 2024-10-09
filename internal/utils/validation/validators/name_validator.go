package validators

import (
	"github.com/Remi-Godin/simple-log/internal/utils/validation"
)

type nameValidator struct {
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

func NewNameValidator() nameValidator {
	var valList []validation.Validator
	valList = append(valList, validation.NewTextLengthValidator(2, 32))
	return nameValidator{
		Validators: valList,
		Links:      make(map[string]string),
	}
}

func (nv nameValidator) GetValidators() []validation.Validator {
	return nv.Validators
}

func (nv nameValidator) GetFieldValue() string {
	return nv.FieldValue
}

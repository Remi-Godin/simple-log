package validation

import (
	"context"
	"fmt"
)

type textLengthValidator struct {
	minLength int
	maxLength int
}

func NewTextLengthValidator(minLength int, maxLength int) textLengthValidator {
	return textLengthValidator{
		minLength: minLength,
		maxLength: maxLength,
	}
}

func (lv textLengthValidator) Validate(ctx context.Context, fieldData string) error {
	if len(fieldData) < lv.minLength {
		errMsg := fmt.Sprintf("Field value too short. Please have a minimum of %d characters.", lv.minLength)
		return ValidationError{errMsg}
	}
	if len(fieldData) > lv.maxLength {
		errMsg := fmt.Sprintf("Field value too long. Please have a maximum of %d characters.", lv.maxLength)
		return ValidationError{errMsg}
	}
	return nil
}

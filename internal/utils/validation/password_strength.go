package validation

import (
	"context"

	passwordValidator "github.com/wagslane/go-password-validator"
)

type passwordStrengthValidator struct {
	minEntropy float64
}

func NewPasswordStrengthValidator(minEntropy float64) passwordStrengthValidator {
	return passwordStrengthValidator{
		minEntropy: minEntropy,
	}
}

func (lv passwordStrengthValidator) Validate(ctx context.Context, fieldData string) error {
	err := passwordValidator.Validate(fieldData, float64(lv.minEntropy))
	if err != nil {
		return err
	}
	return nil
}

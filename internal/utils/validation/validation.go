package validation

import (
	"context"

	"github.com/rs/zerolog/log"
)

type Validator interface {
	Validate(ctx context.Context, fieldData string) error
}

type ValidationError struct {
	Message string
}

func NewValidationError(message string) ValidationError {
	return ValidationError{
		Message: message,
	}
}

func (ve ValidationError) Error() string {
	return ve.Message
}

type ValidatedInputField interface {
	GetValidators() []Validator
	GetFieldValue() string
}

func Validate(ctx context.Context, data ValidatedInputField) error {
	log.Info().Msg("Validating")
	for _, val := range data.GetValidators() {
		log.Info().Msg("Validating")
		err := val.Validate(ctx, data.GetFieldValue())
		if err != nil {
			return err
		}
	}
	return nil
}

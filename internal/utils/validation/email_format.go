package validation

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/rs/zerolog/log"
)

type emailFormatValidator struct{}

func NewEmailFormatValidator() emailFormatValidator {
	return emailFormatValidator{}
}

func (lv emailFormatValidator) Validate(ctx context.Context, fieldData string) error {
	log.Info().Msg(fmt.Sprintf("Checking email format for", fieldData))
	_, err := mail.ParseAddress(fieldData)
	if err != nil {
		return err
	}
	return nil
}

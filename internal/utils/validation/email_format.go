package validation

import "net/mail"

type emailFormatValidator struct{}

func NewEmailFormatValidator() emailFormatValidator {
	return emailFormatValidator{}
}

func (lv emailFormatValidator) Validate(fieldData string) error {
	_, err := mail.ParseAddress(fieldData)
	if err != nil {
		return err
	}
	return nil
}

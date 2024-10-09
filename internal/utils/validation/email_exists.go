package validation

import (
	"context"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
)

type emailExistValidator struct {
}

func NewEmailExistValidator() emailExistValidator {
	return emailExistValidator{}
}

func (lv emailExistValidator) Validate(fieldData string) error {
	_, err := database.New(global.AppData.Conn).GetUserInfo(context.Background(), fieldData)
	if err == nil {
		return ValidationError{"This email address is already in use. Please use a different one."}
	}
	return nil
}

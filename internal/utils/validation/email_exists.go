package validation

import (
	"context"
	"fmt"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/rs/zerolog/log"
)

type emailExistValidator struct {
}

func NewEmailExistValidator() emailExistValidator {
	return emailExistValidator{}
}

func (lv emailExistValidator) Validate(ctx context.Context, fieldData string) error {
	log.Info().Msg(fmt.Sprintf("Checking email exist for %s", fieldData))
	_, err := database.New(global.AppData.Conn).GetUserInfo(ctx, fieldData)
	if err == nil {
		return NewValidationError("This email address is already in use. Please use a different one.")
	}
	return nil
}

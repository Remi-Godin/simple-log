package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

// A simple Jwt handler that uses a secret key
type SimpleJwtHandler struct {
	secret       []byte
	expiryPeriod time.Duration
}

func NewSimpleJwtHandler(secret string, expPeriodMinutes time.Duration) *SimpleJwtHandler {
	return &SimpleJwtHandler{
		secret:       []byte(secret),
		expiryPeriod: expPeriodMinutes,
	}
}

func (this *SimpleJwtHandler) GenerateToken(user string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = jwt.NewNumericDate(time.Now().Add(this.expiryPeriod))
	claims["authorized"] = true
	claims["sub"] = user
	tokenString, err := token.SignedString(this.secret)
	if err != nil {
		log.Error().Err(err).Msg("An error was encountered when trying to sign an auth token for user: " + user)
		return "", err
	}
	return tokenString, nil
}

func (this *SimpleJwtHandler) ValidateToken(token string) (string, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method.")
		}
		return this.secret, nil
	})

	user, userErr := jwtToken.Claims.GetSubject()
	if userErr != nil {
		log.Error().Err(err).Msg("Could not find token user")
	}

	switch {
	case jwtToken.Valid:
		log.Info().Msg("Token successfuly validated for user: " + user)
	case errors.Is(err, jwt.ErrTokenMalformed):
		log.Error().Err(err).Msg("Invalid token format")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		log.Error().Err(err).Msg("Invalid signature")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		log.Error().Err(err).Msg("Token is expired")
	default:
		log.Error().Err(err).Msg("Couldn't handle this token")
	}

	if err != nil {
		log.Error().Err(err).Msg("Could not validate token.")
		return "", err
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	claims["exp"] = jwt.NewNumericDate(time.Now().Add(this.expiryPeriod))
	signedToken, err := jwtToken.SignedString(this.secret)
	if err != nil {
		log.Error().Err(err).Msg("Could not resign token for reissuance.")
	}

	return signedToken, nil
}

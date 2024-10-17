package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func AuthMiddleware(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Need to find a way to reuse the validator. time.Minute is not needed here since we are not giving a new token
		if !ValidateRequest(r, NewSimpleJwtHandler(secret, time.Minute)) {
			// If auth token not found or invalid, redirect to login screen
			//w.Header().Add("HX-Redirect", "/login")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Wrapper function to add Auth middleware to routes
func WithAuth(handleFunc http.HandlerFunc) http.Handler {
	return AuthMiddleware(handleFunc, global.AppData.Env.AuthSecret)
}

type JwtAuthTokenHandler interface {
	GenerateToken(param string) (string, error)
	ValidateToken(token string) (string, error)
}

func ExtractJwtFromCookie(r *http.Request, tag string, validator JwtAuthTokenHandler) (string, error) {
	cookie, err := r.Cookie(tag)
	if err != nil {
		return "", err
	}
	jwtString := cookie.Value
	token, err := validator.ValidateToken(jwtString)
	if err != nil {
		return "", err
	}

	return token, nil
}

func BindJwtToCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
}

func ValidateRequest(r *http.Request, validator JwtAuthTokenHandler) bool {
	token, err := ExtractJwtFromCookie(r, "Authorization", validator)
	if err != nil {
		return false
	}
	_, err = validator.ValidateToken(token)
	if err != nil {
		return false
	}
	return true
}

func ExtractUserFromJwt(r *http.Request) (string, error) {
	tokenStr, err := ExtractJwtFromCookie(r, "Authorization", NewSimpleJwtHandler(global.AppData.Env.AuthSecret, time.Minute))
	if err != nil {
		log.Error().Err(err).Msg("Could not get token from request.")
		return "", err
	}
	jwtToken, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method.")
		}
		return []byte(global.AppData.Env.AuthSecret), nil
	})

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		// Access your claims here
		user := claims["sub"] // Extract the "sub" claim
		email := fmt.Sprintf("%s", user)
		log.Info().Msg(fmt.Sprintf("%s", user))
		return email, nil
	} else {
		log.Error().Msg("Invalid token")
		return "", err
	}
}

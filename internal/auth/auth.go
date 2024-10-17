package auth

import (
	"net/http"
	"time"

	"github.com/Remi-Godin/simple-log/internal/global"
)

func AuthMiddleware(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Need to find a way to reuse the validator. time.Minute is not needed here since we are not giving a new token
		if !ValidateRequest(r, NewSimpleJwtHandler(secret, time.Minute)) {
			// If auth token not found or invalid, redirect to login screen
			//w.Header().Add("HX-Redirect", "/login")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
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

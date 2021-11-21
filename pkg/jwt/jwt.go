package jwt

import (
	"abdukhashimov/mybron.uz/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Jwt struct {
	cfg *config.Config
}

type TokenPayload struct {
	UserID string `json:"user_id"`
}

func (j *Jwt) GenerateToken(payload TokenPayload) (string, error) {
	var (
		tokenString string
		err         error
	)

	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)

	// set claims
	claims["payload"] = payload
	claims["expires_at"] = time.Now().Add(
		time.Hour * time.Duration(j.cfg.TokenExpireHour),
	).Unix()

	tokenString, err = token.SignedString(j.cfg.JWTSecretKey)

	return tokenString, err
}

func (j *Jwt) ParseToken(tokenStr string) (TokenPayload, error) {

	return TokenPayload{}, nil
}

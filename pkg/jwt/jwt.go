package jwt

import (
	"abdukhashimov/mybron.uz/config"
	"abdukhashimov/mybron.uz/pkg/logger"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Jwt struct {
	cfg *config.Config
	log logger.Logger
}

type TokenPayload struct {
	UserID string `json:"user_id"`
}

func (j *Jwt) GenerateToken(payload TokenPayload) (string, error) {
	var (
		tokenString string
		err         error
	)

	j.log.Info("generating jwt token...", logger.Any("payload", payload))

	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)

	// set claims
	claims["payload"] = payload
	claims["expires_at"] = time.Now().Add(
		time.Hour * time.Duration(j.cfg.TokenExpireHour),
	).Unix()

	tokenString, err = token.SignedString(j.cfg.JWTSecretKey)
	if err != nil {
		j.log.Warn("failed to generate jwt token", logger.Error(err))
	}

	return tokenString, err
}

func (j *Jwt) ParseToken(tokenStr string) (TokenPayload, error) {

	return TokenPayload{}, nil
}

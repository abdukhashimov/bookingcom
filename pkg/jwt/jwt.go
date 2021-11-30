package jwt

import (
	"abdukhashimov/mybron.uz/config"
	"abdukhashimov/mybron.uz/pkg/logger"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type Jwt struct {
	cfg *config.Config
	log logger.Logger
}

func NewJwt(cfg *config.Config, log logger.Logger) Jwt {
	return Jwt{
		cfg: cfg,
		log: log,
	}
}

type TokenPayload struct {
	UserID   string `json:"user_id"`
	UserType int    `json:"user_type"`
}

func (j *Jwt) GenerateToken(payload TokenPayload) (string, error) {
	var (
		tokenString string
		err         error
	)

	j.log.Debug("generating jwt token...", logger.Any("payload", payload))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// set claims
	claims["user_id"] = payload.UserID
	claims["user_type"] = payload.UserType
	claims["expires_at"] = time.Now().Add(
		time.Hour * time.Duration(j.cfg.TokenExpireHour),
	).Unix()

	tokenString, err = token.SignedString([]byte(j.cfg.JWTSecretKey))
	if err != nil {
		j.log.Warn("failed to generate jwt token", logger.Error(err))
	}

	return tokenString, err
}

func (j *Jwt) ParseToken(tokenStr string) (TokenPayload, error) {
	var (
		payload TokenPayload
		err     error
		token   *jwt.Token
	)

	j.log.Debug("parsing jwt token", logger.Any("token", tokenStr))

	token, err = jwt.Parse(
		tokenStr,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.cfg.JWTSecretKey), nil
		},
	)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expires := int64(claims["expires_at"].(float64))
		if expires < time.Now().Unix() {
			return TokenPayload{}, errors.New("token is expired")
		}

		payload.UserID = claims["user_id"].(string)
		j.log.Debug("parsed the token", logger.Any("payload", payload))
		return payload, err
	} else {
		j.log.Warn("failed to parse token", logger.Error(err))

		return TokenPayload{}, err
	}
}

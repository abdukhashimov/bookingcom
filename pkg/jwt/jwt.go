package jwt

import "abdukhashimov/mybron.uz/config"

type Jwt struct {
	cfg *config.Config
}

type TokenPayload struct {
	UserID string `json:"user_id"`
}

func (j *Jwt) GenerateToken(payload TokenPayload) (string, error) {

	return "", nil
}

func (j *Jwt) ParseToken(tokenStr string) (TokenPayload, error) {

	return TokenPayload{}, nil
}

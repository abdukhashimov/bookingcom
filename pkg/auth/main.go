package auth

import (
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/pkg/logger"
	"context"

	"github.com/gin-gonic/gin"
)

var (
	noAuthTokenMessage = "no authorization header is provided"
	authTokenIsInvalid = "authorization token is invalid"
)

type auth struct {
	jwt jwt.Jwt
	log logger.Logger
}

// ResponseModel ...
type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

func NewAuth(jwt jwt.Jwt, log logger.Logger) auth {
	return auth{
		jwt: jwt,
		log: log,
	}
}

func (a *auth) MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		// a.log.Debug("auth middleware is running", logger.String("Authorization", authorization))

		if authorization == "" {
			// a.log.Warn(noAuthTokenMessage)
			return
		}

		tokenPayload, err := a.jwt.ParseToken(authorization)
		if err != nil {
			a.log.Warn(authTokenIsInvalid)
			return
		}

		ctx := context.WithValue(c, "auth", tokenPayload)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

package auth

import (
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/pkg/logger"

	"github.com/gin-gonic/gin"
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

type ErrorModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func NewAuth(jwt jwt.Jwt, log logger.Logger) auth {
	return auth{
		jwt: jwt,
		log: log,
	}
}

func (a *auth) makeStatusHeader(c *gin.Context, statusCode int, message string, err error) {
	a.log.Error(message, logger.Int("code", statusCode), logger.Any("error", err))
	c.JSON(statusCode, ErrorModel{
		Code:    statusCode,
		Message: message,
		Error:   err,
	})
}

func (a *auth) MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
		return
	}
}

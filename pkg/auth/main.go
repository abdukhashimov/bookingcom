package auth

import (
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/pkg/logger"
	"context"
	"net/http"
)

type Auth struct {
	jwt jwt.Jwt
	log logger.Logger
}

func (a *Auth) MiddleWare() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get("Authorization")
			a.log.Debug("auth middleware is running", logger.String("Authorization", authorizationHeader))

			if authorizationHeader == "" {
				a.log.Warn("no authorization header is provided")
				next.ServeHTTP(rw, r)
				return
			}

			tokenPayload, err := a.jwt.ParseToken(authorizationHeader)
			if err != nil {
				a.log.Warn("authorization token is invalid")
				http.Error(rw, "Authorization token is invalid", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), tokenPayload, tokenPayload)
			r = r.WithContext(ctx)
			next.ServeHTTP(rw, r)
		})
	}
}

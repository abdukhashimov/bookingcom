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

			if authorizationHeader == "" {
				next.ServeHTTP(rw, r)
				return
			}

			tokenPayload, err := a.jwt.ParseToken(authorizationHeader)
			if err != nil {
				http.Error(rw, "Authorization token is invalid", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), tokenPayload, tokenPayload)
			r = r.WithContext(ctx)
			next.ServeHTTP(rw, r)
		})
	}
}

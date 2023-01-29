package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/romankravchuk/toronto-pizza/internal/config"
	"github.com/romankravchuk/toronto-pizza/internal/service"
)

type JWTAuthMiddleware struct {
	svc         service.IAuthService
	accessToken config.Token
}

func NewJWTAuthMiddleware(svc service.IAuthService, access config.Token) *JWTAuthMiddleware {
	return &JWTAuthMiddleware{svc: svc, accessToken: access}
}

func (m *JWTAuthMiddleware) JWTRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var accessToken string

		cookie, err := r.Cookie(AccessTokenHeader)
		authHeader := r.Header.Get(AuthorizationHeader)
		fields := strings.Fields(authHeader)

		switch {
		case len(fields) != 0 && fields[0] == "Bearer":
			accessToken = fields[1]
		case err == nil:
			accessToken = cookie.Value
		default:
			JSON(w, http.StatusUnauthorized, nil, "", ErrUnauthrized)
			return
		}

		if accessToken == "" {
			JSON(w, http.StatusUnauthorized, nil, "", ErrUnauthrized)
			return
		}

		sub, err := m.svc.ValidateToken(accessToken, m.accessToken.PublicKey)
		if err != nil {
			JSON(w, http.StatusUnauthorized, nil, "", err)
			return
		}

		userDto, err := m.svc.GetUserByPhone(r.Context(), fmt.Sprint(sub))
		if err != nil {
			JSON(w, http.StatusUnauthorized, nil, "", err)
			return
		}

		ctx := context.WithValue(r.Context(), keyUserDTO, userDto)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

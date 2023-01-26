package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/romankravchuk/toronto-pizza/internal/config"
	"github.com/romankravchuk/toronto-pizza/internal/service"
)

type key int

const (
	keyUserDTO key = iota
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
		var (
			resp        = NewApiResponse(http.StatusUnauthorized, map[string]any{"error": "unauthorized"})
			accessToken string
		)

		cookie, _ := r.Cookie(AccessTokenHeader)
		authHeader := r.Header.Get(AuthorizationHeader)
		fields := strings.Fields(authHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else {
			accessToken = cookie.Value
		}

		if accessToken == "" {
			JSON(w, resp.Status, resp)
			return
		}

		sub, err := m.svc.ValidateToken(accessToken, m.accessToken.PublicKey)
		if err != nil {
			resp.SetError(err)
			JSON(w, resp.Status, resp)
			return
		}

		userDto, err := m.svc.GetUserByPhone(r.Context(), fmt.Sprint(sub))
		if err != nil {
			resp.SetError(err)
			JSON(w, resp.Status, resp)
			return
		}

		ctx := context.WithValue(r.Context(), keyUserDTO, userDto)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

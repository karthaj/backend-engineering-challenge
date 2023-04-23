package middleware

import (
	"context"
	httpTransporter "github.com/go-kit/kit/transport/http"
	"net/http"
	"strings"
)

const bearer = "bearer"

func DecodeHttp() httpTransporter.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		token, ok := extractTokenFromAuthHeader(req.Header.Get("Authorization"))
		if !ok {
			return ctx
		}
		return context.WithValue(ctx, "JWT-TOKEN", token)
	}
}

func extractTokenFromAuthHeader(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")

	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != bearer {
		return "", false
	}

	return authHeaderParts[1], true
}

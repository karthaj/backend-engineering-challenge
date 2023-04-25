package codec

import (
	"backend-engineering-challenge/internals/app/middleware"
	"backend-engineering-challenge/internals/domain"
	"context"
	httpTransporter "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

const bearer = "bearer"

// DecodeHttp returns a middleware function that extracts the correlation ID and JWT token from HTTP request headers and adds them to the context.
func DecodeHttp() httpTransporter.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {

		// Set correlation id
		corId := req.Header.Get("correlation-id")
		if corId == "" {
			corId = uuid.New().String()
		}
		ctx = context.WithValue(ctx, domain.CorrelationIdContextKey, corId)

		// Extract JWT token from Authorization header
		token, ok := extractTokenFromAuthHeader(req.Header.Get("Authorization"))
		if !ok {
			return ctx
		}

		// Add JWT token to context
		return context.WithValue(ctx, middleware.JWTTokenContextKey, token)
	}
}

func extractTokenFromAuthHeader(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")

	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != bearer {
		return "", false
	}

	return authHeaderParts[1], true
}

package middleware

import (
	errors "backend-engineering-challenge/internals/app/error"
	"context"
	"github.com/go-kit/kit/endpoint"
)

const JWTTokenContextKey = "JWTToken"

func NewParser() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			_, ok := ctx.Value(JWTTokenContextKey).(string)
			if !ok {
				return nil, errors.NewAuthenticationError(`Need a token`, errors.ErrAuthNoToken)
			}

			// Validate the token with the required payload
			return next(ctx, request)
		}
	}
}

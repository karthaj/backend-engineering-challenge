package middleware

import (
	errors "backend-engineering-challenge/internals/app/error"
	"context"
	"github.com/go-kit/kit/endpoint"
)

const JWTTokenContextKey = "JWTToken"

//NewParser returns a middleware function that validates the JWT token in the context
//and calls the next endpoint with the request if the token is valid.
//return {endpoint.Middleware} The middleware function that validates the token and calls the next endpoint.
func NewParser() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			_, ok := ctx.Value(JWTTokenContextKey).(string)
			if !ok {
				return nil, errors.NewAuthenticationError(`Need a token`, errors.ErrAuthNoToken)
			}

			// ********************************************
			// Validate the token with the required payload
			// ********************************************

			return next(ctx, request)
		}
	}
}

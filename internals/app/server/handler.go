package server

import (
	errors "backend-engineering-challenge/internals/app/error"
	"backend-engineering-challenge/internals/app/middleware"
	"backend-engineering-challenge/internals/app/service"
	"backend-engineering-challenge/internals/transport/http/codec"
	"context"
	httpTransporter "github.com/go-kit/kit/transport/http"
	"net/http"
)

var opts = []httpTransporter.ServerOption{
	httpTransporter.ServerErrorEncoder(errors.ErrorEncoder),
	httpTransporter.ServerBefore(middleware.DecodeHttp()),
}

func Ping() http.Handler {

	var pingingService service.PingService
	pingingService = service.PingServiceStr{}

	return httpTransporter.NewServer(
		func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return pingingService.Pinging(ctx)
		},
		codec.DecodePing,
		codec.EncodePing,
	)
}

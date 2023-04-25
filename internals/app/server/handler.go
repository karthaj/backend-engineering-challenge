package server

import (
	errors "backend-engineering-challenge/internals/app/error"
	"backend-engineering-challenge/internals/app/middleware"
	"backend-engineering-challenge/internals/app/service"
	req_res "backend-engineering-challenge/internals/domain/req-res"
	"backend-engineering-challenge/internals/transport/http/codec"
	"context"
	httpTransporter "github.com/go-kit/kit/transport/http"
	"net/http"
)

var opts = []httpTransporter.ServerOption{
	httpTransporter.ServerErrorEncoder(errors.ErrorEncoder),
	httpTransporter.ServerBefore(codec.DecodeHttp()),
}

func Ping() http.Handler {

	svc := service.PingService{}

	return httpTransporter.NewServer(
		func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return svc.Pinging(ctx)
		},
		codec.DecodePing,
		codec.EncodePing,
	)
}

func DoTransaction() (handler http.Handler) {

	svc := service.AccountService{}

	return httpTransporter.NewServer(
		middleware.NewParser()(
			func(ctx context.Context, request interface{}) (response interface{}, err error) {
				req := request.(req_res.DoTransactionRequest)
				return svc.DoTransaction(ctx, req)
			}),
		codec.DecodeDoTransaction,
		codec.EncodeDoTransaction,
		opts...,
	)

}

func GetAccountDetailsByID() (handler http.Handler) {
	svc := service.AccountService{}

	return httpTransporter.NewServer(
		middleware.NewParser()(
			func(ctx context.Context, request interface{}) (response interface{}, err error) {
				req := request.(req_res.GetAccountDetailsByIDRequest)
				return svc.GetAccountDetailsByID(ctx, req)
			}),
		codec.DecodeGetAccountDetailsByID,
		codec.EncodeGetAccountDetailsByID,
		opts...,
	)

}

func GetAllAccountDetails() (handler http.Handler) {
	svc := service.AccountService{}

	return httpTransporter.NewServer(
		middleware.NewParser()(
			func(ctx context.Context, request interface{}) (response interface{}, err error) {

				return svc.GetAllAccountDetails(ctx)
			}),
		codec.DecodeGetAllAccountDetails,
		codec.EncodeGetAllAccountDetails,
		opts...,
	)

}

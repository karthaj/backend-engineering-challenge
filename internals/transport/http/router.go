package http

import (
	"backend-engineering-challenge/internals/app/server"
	"backend-engineering-challenge/internals/config"
	"backend-engineering-challenge/internals/domain/log"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

const logPrefixRouter = "backend-engineering-challenge.internals.transport.http.router"

var httpServer *http.Server

func Init(ctx context.Context) {
	r := mux.NewRouter()
	port := config.AppConf.Port

	r.PathPrefix("/debug").Handler(http.DefaultServeMux)

	r.Handle(`/v1.0/ping`, methodControl(http.MethodGet, server.Ping()))

	r.Handle(`/v1.0/ping`, methodControl(http.MethodGet, server.Ping()))

	StartServer(ctx, port, r)

}

// Validate the HTTP method
func methodControl(method string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			h.ServeHTTP(w, r)
		} else {
			http.Error(w, `Method not allowed`, http.StatusMethodNotAllowed)
		}
	})
}

// StartServer ...
func StartServer(ctx context.Context, port int, r http.Handler) {
	running := make(chan interface{}, 1)

	httpServer = &http.Server{
		Addr:         fmt.Sprintf(`:%d`, port),
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 5,
		Handler:      r,
	}

	go func(ctx context.Context) {
		err := httpServer.ListenAndServe()
		if err != nil {
			log.FatalContext(ctx, `Cannot start web server`, err)
		}
		running <- `done`
	}(ctx)

	go func() {
		log.FatalContext(ctx, logPrefixRouter, http.ListenAndServe(":6060", nil))
	}()

	log.InfoContext(ctx, logPrefixRouter, fmt.Sprintf("HTTP router started on port \033[0;32m[%d]\033[0m", port))

	<-running
}

// StopServer ...
func StopServer(ctx context.Context) {
	if err := httpServer.Shutdown(ctx); err != nil {
		log.FatalContext(ctx, logPrefixRouter, `Failed to gracefully shutdown server`)
	}

	log.FatalContext(ctx, logPrefixRouter, `Success gracefully shutting down server`)
}

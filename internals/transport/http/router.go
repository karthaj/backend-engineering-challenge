package http

import (
	"backend-engineering-challenge/internals/app/server"
	"backend-engineering-challenge/internals/config"
	"backend-engineering-challenge/internals/domain/log"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"time"
)

const logPrefixRouter = "backend-engineering-challenge.internals.transport.http.router"

var httpServer *http.Server

func Init(ctx context.Context) {
	r := mux.NewRouter()        // create new router
	port := config.AppConf.Port // get port number from configuration

	// register debug route handler
	r.PathPrefix("/debug").Handler(http.DefaultServeMux)

	// register routes with their corresponding handlers
	r.Handle(`/v1.0/ping`, methodControl(http.MethodGet, server.Ping()))

	r.Handle(`/v1.0/account/transaction`, methodControl(http.MethodPost, server.DoTransaction()))
	r.Handle(`/v1.0/account/get/id/{id}`, methodControl(http.MethodGet, server.GetAccountDetailsByID()))
	r.Handle(`/v1.0/account/get/all`, methodControl(http.MethodGet, server.GetAllAccountDetails()))

	// start server with given context, port number, and router
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

// StartServer starts an HTTP server with the provided context, port, and router. It spawns a goroutine to listen
// for incoming requests and returns once the server is running. If the server fails to start, it will log an error
// and terminate the program.
func StartServer(ctx context.Context, port int, r http.Handler) {
	// Create a channel to signal when the server is running
	running := make(chan interface{}, 1)
	address := net.JoinHostPort("0.0.0.0", fmt.Sprintf("%d", port))

	// Create an HTTP server with the provided settings
	httpServer = &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 5,
		Handler:      r,
	}

	// Spawn a goroutine to listen for incoming requests
	go func(ctx context.Context) {
		err := httpServer.ListenAndServe()
		if err != nil {
			log.FatalContext(ctx, `Cannot start web server`, err)
		}
		running <- `done`
	}(ctx)

	// Spawn a goroutine to listen for HTTP profiling requests on port 6060
	go func() {
		log.FatalContext(ctx, logPrefixRouter, http.ListenAndServe(":6060", nil))
	}()

	// Log that the server has started and on which port
	log.InfoContext(ctx, logPrefixRouter, fmt.Sprintf("HTTP router started on port \033[0;32m[%d]\033[0m", port))

	// Wait until the server is running before returning
	<-running
}

// StopServer ...
func StopServer(ctx context.Context) {
	if err := httpServer.Shutdown(ctx); err != nil {
		log.FatalContext(ctx, logPrefixRouter, `Failed to gracefully shutdown server`)
	}

	log.FatalContext(ctx, logPrefixRouter, `Success gracefully shutting down server`)
}

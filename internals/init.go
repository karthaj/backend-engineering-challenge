package internals

import (
	"backend-engineering-challenge/internals/config"
	"backend-engineering-challenge/internals/database"
	"backend-engineering-challenge/internals/domain"
	"backend-engineering-challenge/internals/domain/log"
	"backend-engineering-challenge/internals/transport/http"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"os"
)

const logPrefixInit = "backend-engineering-challenge.internals.init"

func Init() {

	// Create a channel to listen for OS signals, such as interrupts.
	sigs := make(chan os.Signal, 1)
	// Initialize the logging package.
	log.Init()

	// Load configuration settings for the system using the config package.
	config.InitConfig()

	// Create a new context with a unique correlation ID using uuid.New() function and a cancel function to ensure proper cleanup.
	parentCtx := context.WithValue(context.Background(), domain.CorrelationIdContextKey, uuid.New())
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	// Initialize the database package with the newly created context.
	database.Init(ctx)

	// Load the database using the context.
	err := database.LoadDB(ctx)
	if err != nil {
		log.ErrorContext(ctx, fmt.Sprintf("%s.%s", logPrefixInit, "database.DbCon.LoadDB()"), "error", err)
	}

	// Initialize the http package with the context.
	http.Init(ctx)

	// Wait for an OS signal to interrupt the program, and then stop the server using the http package.
	select {
	case <-sigs:
		logrus.Info(ctx, "Shutting down server", "OS interrupt")
		http.StopServer(ctx)
	}
}

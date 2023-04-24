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
	sigs := make(chan os.Signal, 1)

	// init log
	log.Init()

	// Load Config
	config.InitConfig()
	parentCtx := context.WithValue(context.Background(), domain.CorrelationIdContextKey, uuid.New())
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	database.Init(ctx)

	err := database.LoadDB(ctx)
	if err != nil {
		log.ErrorContext(ctx, fmt.Sprintf("%s.%s", logPrefixInit, "database.DbCon.LoadDB()"), "error", err)
	}

	http.Init(ctx)

	select {
	case <-sigs:
		logrus.Info(ctx, "Shutting down server", "OS interrupt")
		http.StopServer(ctx)
	}
}

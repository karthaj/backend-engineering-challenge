package internals

import (
	"backend-engineering-challenge/internals/config"
	"backend-engineering-challenge/internals/domain/logger"
	"backend-engineering-challenge/internals/transport/http"
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"os"
)

func Init() {
	sigs := make(chan os.Signal, 1)

	// init logger
	logger.Init()

	// Load Config
	config.InitConfig()
	parentCtx := context.WithValue(context.Background(), "uuid", uuid.New())
	ctx, _ := context.WithCancel(parentCtx)

	http.Init(ctx)

	select {
	case <-sigs:
		logrus.Info(ctx, "Shutting down server", "OS interrupt")
		http.StopServer(ctx)
	}
}

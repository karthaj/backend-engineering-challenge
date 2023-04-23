package service

import (
	"context"
	"fmt"
	"time"
)

type PingService interface {
	Pinging(ctx context.Context) (string, error)
}

type PingServiceStr struct{}

func (r PingServiceStr) Pinging(ctx context.Context) (string, error) {
	now := time.Now()
	formatted := now.Format(time.RFC1123)
	return fmt.Sprintf("P I N G - %v", formatted), nil
}

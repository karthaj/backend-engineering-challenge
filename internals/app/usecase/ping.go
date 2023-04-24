package usecase

import (
	"fmt"
	"time"
)

type PingUsecase struct{}

func NewPingUsecase() PingUsecase {
	i := PingUsecase{}
	return i
}

func (r PingUsecase) Ping() (string, error) {

	now := time.Now()
	formatted := now.Format(time.RFC1123)

	res := fmt.Sprintf("P I N G - %v", formatted)

	return res, nil
}

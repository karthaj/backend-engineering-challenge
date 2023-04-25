package codec

import (
	"context"
	"encoding/json"
	"net/http"
)

type PingResponse struct {
	Data struct {
		Status string
	}
	Meta struct {
		Code    int
		Message string
	}
}

func DecodePing(context.Context, *http.Request) (request interface{}, err error) {
	return nil, nil
}

func EncodePing(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	result := PingResponse{}
	result.Meta.Code = 200
	result.Meta.Message = "Success"
	result.Data.Status = response.(string)
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(result)
}

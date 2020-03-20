package logic

import (
	"context"
	"time"
)

func Success(ack Acker) (Acker, error) {
	var header = ack.GetHeader()
	header.Code = 200
	header.Msg = "Success"
	return ack, nil
}

func FailCode(ack Acker, code int32, msg string) (Acker, error) {
	var header = ack.GetHeader()
	header.Code = code
	header.Msg = msg
	return ack, nil
}

func Fail(ack Acker, msg string) (Acker, error) {
	var header = ack.GetHeader()
	header.Code = 400
	header.Msg = msg
	if msg == "" {
		header.Msg = "Unexpected error"
	}
	return ack, nil
}

func Error(ack Acker, err error) (Acker, error) {
	var header = ack.GetHeader()
	header.Code = 500
	header.Msg = "Unexpected error"
	return ack, err
}

func NewContext(ttls ...time.Duration) (context.Context, context.CancelFunc) {
	var ttl = 5 * time.Second
	if len(ttls) > 0 {
		ttl = ttls[0]
	}
	return context.WithTimeout(context.Background(), ttl)
}

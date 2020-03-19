package logic

import (
	"context"
	"time"
)

func Success(ack Acker) error {
	var header = ack.GetHeader()
	header.Code = 200
	header.Msg = "Success"
	return nil
}

func FailCode(ack Acker, code int32, msg string) error {
	var header = ack.GetHeader()
	header.Code = code
	header.Msg = msg
	return nil
}

func Fail(ack Acker, msg string) error {
	var header = ack.GetHeader()
	header.Code = 400
	header.Msg = msg
	return nil
}

func Error(ack Acker, err error) error {
	var header = ack.GetHeader()
	header.Code = 500
	header.Msg = "Unexpected error"
	return err
}

func NewContext(ttls ...time.Duration) (context.Context, context.CancelFunc) {
	var ttl = 5 * time.Second
	if len(ttls) > 0 {
		ttl = ttls[0]
	}
	return context.WithTimeout(context.Background(), ttl)
}

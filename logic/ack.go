package logic

import (
	"chat/protocol"
	"context"
	"reflect"
	"time"
)

func Success(ack protocol.Acker) (protocol.Acker, error) {
	var header = ack.GetHeader()
	if header == nil {
		header = &protocol.AckHeader{}
		var v = reflect.Indirect(reflect.ValueOf(ack))
		v.FieldByName("Header").Set(reflect.ValueOf(header))
	}
	header.Code = 200
	header.Msg = "Success"
	return ack, nil
}

func FailCode(ack protocol.Acker, code int32, msg string) (protocol.Acker, error) {
	var header = ack.GetHeader()
	if header == nil {
		header = &protocol.AckHeader{}
		var v = reflect.Indirect(reflect.ValueOf(ack))
		v.FieldByName("Header").Set(reflect.ValueOf(header))
	}
	header.Code = code
	header.Msg = msg
	return ack, nil
}

func Fail(ack protocol.Acker, msg string) (protocol.Acker, error) {
	var header = ack.GetHeader()
	if header == nil {
		header = &protocol.AckHeader{}
		var v = reflect.Indirect(reflect.ValueOf(ack))
		v.FieldByName("Header").Set(reflect.ValueOf(header))
	}
	header.Code = 400
	header.Msg = msg
	if msg == "" {
		header = &protocol.AckHeader{}
		var v = reflect.Indirect(reflect.ValueOf(ack))
		v.FieldByName("Header").Set(reflect.ValueOf(header))
	}
	return ack, nil
}

func AccessDeined(ack protocol.Acker) (protocol.Acker, error) {
	return Fail(ack, "access deined")
}

func Error(ack protocol.Acker, err error) (protocol.Acker, error) {
	var header = ack.GetHeader()
	if header == nil {
		header = &protocol.AckHeader{}
		var v = reflect.Indirect(reflect.ValueOf(ack))
		v.FieldByName("Header").Set(reflect.ValueOf(header))
	}
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

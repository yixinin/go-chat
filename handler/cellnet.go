package handler

import (
	"chat/protocol"
	"fmt"
	"reflect"
	"strings"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/gogopb"
	"github.com/davyxu/cellnet/util"
)

func init() {
	for k, v := range protocol.ProtocolMap {
		RegisterProtobuf(k, v...)
	}
}

func RegisterProtobuf(pkgName string, args ...interface{}) {
	for _, v := range args {
		var t = reflect.TypeOf(v).Elem()
		var fullName = strings.ToLower(fmt.Sprintf("%s.%s", pkgName, t.Name()))
		var id = int(util.StringHash(fullName))
		fmt.Printf("name:%s,hashid:%d\n", fullName, id)
		cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
			Codec: codec.MustGetCodec("gogopb"),
			Type:  t,
			ID:    id,
		})
	}
}

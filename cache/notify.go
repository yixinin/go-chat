package cache

import (
	"chat/protocol"
	"encoding/json"
	"fmt"
	"go-lib/db"
)

func GetNotifyMsgKey(uid int64) string {
	return fmt.Sprintf("list:user:msg:%d", uid)
}

func CacheNotifyMessage(uid int64, msg []byte) error {
	var key = GetNotifyMsgKey(uid)
	// var v, err = json.Marshal(msg)
	// if err != nil {
	// 	return err
	// }
	return db.Redis.RPush(key, msg).Err()
}

func GetNotifyMessage(uid int64) (*protocol.CacheMessage, error) {
	var key = GetNotifyMsgKey(uid)
	var buf, err = db.Redis.LPop(key).Bytes()
	var msg *protocol.CacheMessage
	err = json.Unmarshal(buf, msg)
	return msg, err
}

func GetAllNotifyMessage(uid int64) ([]*protocol.CacheMessage, error) {
	var key = GetNotifyMsgKey(uid)
	var msgs = make([]*protocol.CacheMessage, 0, 1)
	var err error
	for true {
		var buf []byte
		buf, err = db.Redis.LPop(key).Bytes()
		if err != nil {
			break
		}
		if len(buf) == 0 {
			break
		}
		var msg *protocol.CacheMessage
		err = json.Unmarshal(buf, msg)
		if err != nil {
			break
		}
		msgs = append(msgs, msg)
	}

	return msgs, err
}

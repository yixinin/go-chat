package cache

import (
	"encoding/json"
	"go-chat/protocol"
	"go-lib/db"
)

func GetNotifyMsgKey(uid string) string {
	return "list:user:msg"
}

func CacheNotifyMessage(uid string, msg protocol.CacheMessage) error {
	var key = GetNotifyMsgKey(uid)
	var v, err = json.Marshal(msg)
	if err != nil {
		return err
	}
	return db.Redis.RPush(key, v).Err()
}

func GetNotifyMessage(uid string) (*protocol.CacheMessage, error) {
	var key = GetNotifyMsgKey(uid)
	var buf, err = db.Redis.LPop(key).Bytes()
	var msg *protocol.CacheMessage
	err = json.Unmarshal(buf, msg)
	return msg, err
}

func GetAllNotifyMessage(uid string) ([]*protocol.CacheMessage, error) {
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

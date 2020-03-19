package cache

import (
	"chat/protocol"
	"encoding/json"
	"fmt"
	"go-lib/db"
)

func GetRealTimeNotifyKey(uid int64) string {
	return fmt.Sprintf("list:user:realtime:%d", uid)
}

func CacheRealTimeNotify(uid int64, msg *protocol.RealTimeNotify) error {
	var key = GetRealTimeNotifyKey(uid)
	var body, err = json.Marshal(msg)
	if err != nil {
		return err
	}
	return db.Redis.RPush(key, body).Err()
}
func GetRealTime(uid int64, msg *protocol.RealTimeNotify) (err error) {
	var key = GetRealTimeNotifyKey(uid)
	buf, err := db.Redis.LPop(key).Bytes()
	if err != nil {
		return
	}
	err = json.Unmarshal(buf, msg)
	return err
}

func GetAllNotifyMessage(uid int64) ([]*protocol.PollAck_Message, error) {
	var key = GetRealTimeNotifyKey(uid)
	var msgs = make([]*protocol.PollAck_Message, 0, 1)
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
		var msg *protocol.PollAck_Message
		err = json.Unmarshal(buf, msg)
		if err != nil {
			break
		}
		msgs = append(msgs, msg)
	}

	return msgs, err
}

//好友未读
func IncUserMessage(uid int64, count int) error {

}

func IncGroup()

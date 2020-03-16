package cache

import (
	"fmt"
	"go-lib/db"
	"strconv"
	"time"
)

func GetRoomKey(roomId int32) string {
	return fmt.Sprintf("hash:room:%d", roomId)
}
func GetUserRoomKey(uid string) string {
	return "kv:user:room:" + uid
}

func JoinRoom(uid string, roomId int32) error {
	var roomKey = GetRoomKey(roomId)
	var userRoomKey = GetUserRoomKey(uid)

	//设置用户的房间id
	err := db.Redis.SetNX(userRoomKey, strconv.FormatInt(int64(roomId), 10), 24*time.Hour).Err()
	if err != nil {
		return err
	}

	//设置房间用户列表
	err = db.Redis.SAdd(roomKey, uid).Err()
	if err != nil {
		return err
	}
	return nil
}

func LeaveRoom(uid string) error {
	var userRoomKey = GetUserRoomKey(uid)
	roomId, err := db.Redis.Get(userRoomKey).Int()
	if err != nil {
		return err
	}
	var roomKey = GetRoomKey(int32(roomId))
	db.Redis.Del(userRoomKey)
	db.Redis.SRem(roomKey, uid)
	return nil
}

func DiscardRoom(roomId int32) error {
	var roomKey = GetRoomKey((roomId))
	uids, err := db.Redis.SMembers(roomKey).Result()
	if err != nil {
		return err
	}
	var keys = make([]string, len(uids))

	for i, uid := range uids {
		var key = GetUserRoomKey(uid)
		keys[i] = key
	}
	db.Redis.Del(keys...)
	return nil
}

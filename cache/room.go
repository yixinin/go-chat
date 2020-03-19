package cache

import (
	"fmt"
	"go-lib/db"
	"go-lib/log"
	"strconv"
	"strings"
	"time"
)

func GetRoomKey(roomId int32) string {
	return fmt.Sprintf("hash:room:%d", roomId)
}
func GetUserRoomKey(uid int64) string {
	return fmt.Sprintf("kv:user:room:%d", uid)
}

func GetUserRoomInfo(uid int64) (int32, string, error) {
	var userRoomKey = GetUserRoomKey(uid)
	info, err := db.Redis.Get(userRoomKey).Result()
	s := strings.Split(info, ",")
	if len(s) != 2 {
		return 0, "", nil
	}
	rid, err := strconv.ParseInt(s[0], 10, 32)
	if err != nil {
		return 0, "", err
	}

	return int32(rid), string(s[1]), err
}

func GetRoomMembers(rid int32) ([]int64, error) {
	var roomKey = GetRoomKey(rid)
	res, err := db.Redis.SMembers(roomKey).Result()
	var uids = make([]int64, 0, len(res))
	for _, v := range res {
		var uid, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Error(err)
			continue
		}
		uids = append(uids, uid)
	}
	return uids, err
}

func JoinRoom(uid int64, addr string, roomId int32) error {
	var roomKey = GetRoomKey(roomId)
	var userRoomKey = GetUserRoomKey(uid)

	//设置用户的房间id
	var v = fmt.Sprintf("%d,%s", roomId, addr)
	err := db.Redis.SetNX(userRoomKey, v, 24*time.Hour).Err()
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

func LeaveRoom(uid int64) (int32, error) {
	var userRoomKey = GetUserRoomKey(uid)
	roomId, err := db.Redis.Get(userRoomKey).Int()
	if err != nil {
		return 0, err
	}
	var roomKey = GetRoomKey(int32(roomId))
	db.Redis.Del(userRoomKey)
	db.Redis.SRem(roomKey, uid)
	return int32(roomId), nil
}

func DiscardRoom(roomId int32) error {
	var roomKey = GetRoomKey((roomId))
	uids, err := db.Redis.SMembers(roomKey).Result()
	if err != nil {
		return err
	}
	var keys = make([]string, len(uids))
	keys = append(keys, roomKey)
	for i, v := range uids {
		var uid, _ = strconv.ParseInt(v, 10, 64)
		var key = GetUserRoomKey(uid)
		keys[i] = key
	}
	db.Redis.Del(keys...)
	return nil
}

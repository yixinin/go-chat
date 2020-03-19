package cache

import (
	"fmt"
	"go-lib/db"
	"go-lib/utils"
	"strconv"
	"time"
)

func GetTokenKey(token string) string {
	return "user:token:" + token
}

func GetUserDeviceToken(uid int64) string {
	return fmt.Sprintf("user:device:%d", uid)
}

func SetToken(uid int64, deviceType int32) (token, oldToken string, err error) {
	token = utils.UUID()
	var tokenKey = GetTokenKey(token)
	var deviceKey = GetUserDeviceToken(uid)

	oldToken = db.Redis.HGet(deviceKey, strconv.FormatInt(int64(deviceType), 10)).String()

	err = db.Redis.Set(tokenKey, uid, 30*time.Minute).Err()
	if err != nil {
		return
	}
	//存储用户登录设备
	err = db.Redis.HSet(deviceKey, strconv.FormatInt(int64(deviceType), 10), token).Err()
	if err != nil {
		db.Redis.Del(tokenKey)
	}
	return
}

func GetToken(token string) (uid int64, err error) {
	var tokenKey = GetTokenKey(token)
	uid, err = db.Redis.Get(tokenKey).Int64()
	if err != nil {
		return
	}

	return
}
func CheckToken(token string) (uid int64, err error) {
	var tokenKey = GetTokenKey(token)
	uid, err = db.Redis.Get(tokenKey).Int64()
	if err != nil {
		return
	}
	db.Redis.Expire(tokenKey, 30*time.Minute)
	return
}

func GetDeviceToken(uid int64, token string) (deviceType int32, err error) {

	var deviceKey = GetUserDeviceToken(uid)
	var m map[string]string
	m, _ = db.Redis.HGetAll(deviceKey).Result()
	if m != nil {
		for k, v := range m {
			if v == token {
				n, _ := strconv.ParseInt(k, 10, 32)
				deviceType = int32(n)
				return
			}
		}
	}
	return
}

func DelDevice(uid int64, deviceType int32) error {
	var deviceKey = GetUserDeviceToken(uid)
	return db.Redis.HDel(deviceKey, strconv.FormatInt(int64(deviceType), 10)).Err()
}

func DelToken(token string) error {
	var tokenKey = GetTokenKey(token)
	return db.Redis.Del(tokenKey).Err()
}

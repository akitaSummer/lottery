package utils

import (
	"fmt"
	"lottery/datasource"
)

func getLuckyLockKey(uid int) string {
	return fmt.Sprintf("lucky_lock_%d", uid)
}

func LockLucky(uid int) bool {
	key := getLuckyLockKey(uid)
	cacheObj := datasource.InstanceCache()
	// 设置3s作为释放时间，防止死锁
	rs, _ := cacheObj.Do("SET", key, 1, "EX", 3, "NX")
	if rs == "OK" {
		return true
	} else {
		return false
	}
}

func UnlockLucky(uid int) bool {
	key := getLuckyLockKey(uid)
	cacheObj := datasource.InstanceCache()
	rs, _ := cacheObj.Do("DEL", key)
	if rs == "OK" {
		return true
	} else {
		return false
	}
}

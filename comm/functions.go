package comm

import (
	"lottery/conf"
	"math/rand"
	"time"
)

func NowUnix() int {
	return int(time.Now().In(conf.SysTimeLocation).Unix())
}

// 将unix时间戳格式化为yyyymmdd H:i:s格式字符串
func FormatFromUnixTime(t int64) string {
	if t > 0 {
		return time.Unix(t, 0).Format(conf.SysTimeform)
	} else {
		return time.Now().Format(conf.SysTimeform)
	}
}

// 将字符串转成时间
func ParseTime(str string) (time.Time, error) {
	return time.ParseInLocation(conf.SysTimeform, str, conf.SysTimeLocation)
}

// 得到一个随机数
func Random(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if max < 1 {
		return r.Int()
	} else {
		return r.Intn(max)
	}
}

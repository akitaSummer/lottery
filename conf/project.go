package conf

import "time"

// 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

// cookie中的加密验证密钥
var CookieSecret = "hellolottery"

const SysTimeform = "2006-01-02 15:04:05"
const SysTimeformShort = "2006-01-02"

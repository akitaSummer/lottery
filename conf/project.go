package conf

import "time"

// 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

// cookie中的加密验证密钥
var CookieSecret = "hellolottery"

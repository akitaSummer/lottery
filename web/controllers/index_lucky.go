package controllers

import (
	"lottery/comm"
	"lottery/conf"
	"lottery/web/utils"
)

// localhost:8080/lucky
func (c *IndexController) GetLucky() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	// 验证登录用户
	loginuser := comm.GetLoginUser(c.Ctx.Request())
	if loginuser == nil || loginuser.Uid < 1 {
		rs["code"] = 101
		rs["msg"] = "请先登录，在来抽奖"
		return rs
	}
	// 用户分布式锁，防止用户多次重入
	ok := utils.LockLucky(loginuser.Uid)
	if ok {
		defer utils.UnlockLucky(loginuser.Uid)
	} else {
		rs["code"] = 102
		rs["msg"] = "正在抽奖，请稍后重试"
		return rs
	}
	// 验证用户今日参与次数
	ok = c.checkUserday(loginuser.Uid)
	if !ok {
		rs["code"] = 103
		rs["msg"] = "今日抽奖次数已用完，请明日再来"
		return rs
	}
	// 验证IP今日参与次数
	ip := comm.ClientIp(c.Ctx.Request())
	ipDayNum := utils.IncrIpLuckyNum(ip)
	if ipDayNum > conf.IpLimitMax {
		rs["code"] = 104
		rs["msg"] = "相同IP参与次数太多，明天再来参与吧"
		return rs
	}
	// 验证IP黑名单
	// 验证用户黑名单
	// 获得抽奖编码
	// 匹配奖品是否中奖
	// 有限制奖品发放
	// 不同编码优惠券发放
	// 记录中奖记录
	// 返回中奖结果

	return rs
}

package controllers

import (
	"lottery/comm"
	"lottery/models"
)

func (c *IndexController) prizeLarge(ip string, loginuser *models.ObjLoginuser, userInfo *models.LtUser, blackipInfo *models.LtBlackip) {
	nowTime := comm.NowUnix()
	blackTime := 30 * 86400
	if userInfo == nil || userInfo.Id <= 0 {
		userInfo = &models.LtUser{
			Id:         loginuser.Uid,
			Username:   loginuser.Username,
			Blacktime:  nowTime + blackTime,
			SysCreated: nowTime,
			SysIp:      ip,
		}
		c.ServiceUser.Create(userInfo)
	} else {
		userInfo = &models.LtUser{
			Id:         loginuser.Uid,
			Blacktime:  nowTime + blackTime,
			SysUpdated: nowTime,
		}
		c.ServiceUser.Update(userInfo, nil)
	}
	if blackipInfo == nil || blackipInfo.Id <= 0 {
		blackipInfo = &models.LtBlackip{
			Ip:         ip,
			BlackTime:  nowTime + blackTime,
			SysCreated: nowTime,
		}
		c.ServiceBlackip.Create(blackipInfo)
	} else {
		blackipInfo.BlackTime = nowTime + blackTime
		blackipInfo.SysUpdated = nowTime
		c.ServiceBlackip.Update(blackipInfo, nil)
	}
}

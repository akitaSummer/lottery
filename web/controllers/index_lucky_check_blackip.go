package controllers

import (
	"lottery/models"
	"time"
)

func (c *IndexController) checkBlackip(ip string) (bool, *models.LtBlackip) {
	info := c.ServiceBlackip.GetByIp(ip)
	if info == nil || info.Ip == "" {
		return true, nil
	}
	if info.BlackTime > int(time.Now().Unix()) {
		// IP黑名单存在，并且仍在黑名单有效期内
		return false, info
	}
	return true, info
}

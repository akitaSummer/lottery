package controllers

import (
	"fmt"
	"log"
	"lottery/conf"
	"lottery/models"
	"strconv"
	"time"
)

func (c *IndexController) checkUserday(uid int) bool {
	userdayInfo := c.ServiceUserday.GetUserToday(uid)
	if userdayInfo != nil && userdayInfo.Uid == uid {
		// 今日存在抽奖记录
		if userdayInfo.Num >= conf.UserPrizeMax {
			return false
		} else {
			userdayInfo.Num++
			err := c.ServiceUserday.Update(userdayInfo, nil)
			if err != nil {
				log.Println("index_lucky_check_userday ServiceUser.Update err= ", err)
			}
		}
	} else {
		// 创建用户今日的参与记录
		y, m, d := time.Now().Date()
		strDay := fmt.Sprintf("%d%02d%02d", y, m, d)
		day, _ := strconv.Atoi(strDay)
		userdayInfo = &models.LtUserday{
			Uid:        uid,
			Day:        day,
			Num:        1,
			SysCreated: int(time.Now().Unix()),
		}
		err := c.ServiceUserday.Create(userdayInfo)
		log.Println("index_lucky_check_userday ServiceUser.Create err= ", err)
	}
	return true
}

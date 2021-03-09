package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"lottery/comm"
	"lottery/services"
	"time"
)

type AdminUserController struct {
	Ctx            iris.Context
	ServiceUser    services.UserService
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserday services.UserDayService
	ServiceBlackip services.BlackipService
}

func (c *AdminUserController) Get() mvc.Result {
	page := c.Ctx.URLParamIntDefault("page", 1)
	size := 100
	pagePrev := ""
	pageNext := ""

	dataList := c.ServiceUser.GetAll(page, size)
	total := c.ServiceUser.CountAll()
	if len(dataList) < int(total) {
		pageNext = fmt.Sprintf("%d", page+1)
	}
	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}
	return mvc.View{
		Name: "admin/user.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Channel":  "user",
			"DataList": dataList,
			"Total":    total,
			"Now":      comm.NowUnix(),
			"PagePrev": pagePrev,
			"PageNext": pageNext,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminUserController) GetBlack() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	t := c.Ctx.URLParamIntDefault("time", 0)
	if id > 0 {
		user := c.ServiceUser.Get(id)
		user.Blacktime = int(time.Now().Unix()) + t*86400
		err := c.ServiceUser.Update(user, []string{"black_time", "sys_update"})
		if err != nil {
			return mvc.Response{
				Text: fmt.Sprintf("设置用户黑名单失败, err=%s", err),
			}
		}
	}
	return mvc.Response{
		Path: "/admin/user",
	}
}

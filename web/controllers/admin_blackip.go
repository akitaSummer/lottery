package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"lottery/comm"
	"lottery/models"
	"lottery/services"
)

type AdminBlackIpController struct {
	Ctx            iris.Context
	ServiceBlackIp services.BlackipService
}

func (c *AdminBlackIpController) Get() mvc.Result {
	page := c.Ctx.URLParamIntDefault("page", 1)
	size := 100
	pagePrev := ""
	pageNext := ""

	dataList := c.ServiceBlackIp.GetAll(page, size)
	total := len(dataList)
	if total >= size {
		total = int(c.ServiceBlackIp.CountAll())
		pageNext = fmt.Sprintf("%d", page+1)
	}
	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}

	return mvc.View{
		Name: "/admin/blackip.html",
		Data: iris.Map{
			"Title":    "黑名单管理",
			"Channel":  "blackip",
			"PagePrev": pagePrev,
			"PageNext": pageNext,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminBlackIpController) GetBlack() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	t := c.Ctx.URLParamIntDefault("time", 0)
	if err == nil {
		t = comm.NowUnix() + t*86400
		c.ServiceBlackIp.Update(&models.LtBlackip{Id: id, SysCreated: t, SysUpdated: comm.NowUnix()},
			[]string{"black_time", "sys_updated"})
	}
	return mvc.Response{
		Path: "/admin/blackip.html",
	}
}

package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"lottery/comm"
	"lottery/models"
	"lottery/services"
	"strings"
)

type AdminCodeController struct {
	Ctx         iris.Context
	ServiceCode services.CodeService
}

func (c *AdminCodeController) Get() mvc.Result {
	page := c.Ctx.URLParamIntDefault("page", 1)
	giftId := c.Ctx.URLParamIntDefault("gift_id", 0)
	size := 100
	pagePrev := ""
	pageNext := ""
	dataList := c.ServiceCode.GetAll(page, size)
	total := len(dataList)
	if total >= size {
		total = int(c.ServiceCode.CountAll())
		pageNext = fmt.Sprintf("%d", page+1)
	}
	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}
	return mvc.View{
		Name: "/admin/code.html",
		Data: iris.Map{
			"Title":    "优惠券管理",
			"Channel":  "code",
			"GiftId":   giftId,
			"DataList": dataList,
			"Total":    total,
			"PagePrev": pagePrev,
			"PageNext": pageNext,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminCodeController) PostImport() {
	giftId := c.Ctx.URLParamIntDefault("gift_id", 0)
	if giftId < 1 {
		c.Ctx.Text("没有指定奖品ID，无法导入，<a href='' onclick='history.go(-1)'>返回</a>")
		return
	}
	gift := c.ServiceCode.Get(giftId)
	if gift == nil || gift.Id < 1 {
		c.Ctx.Text("奖品信息不存在或者奖品类型不是优惠券，无法导入，<a href='' onclick='history.go(-1)'>返回</a>\"")
		return
	}
	codes := c.Ctx.PostValue("codes")
	now := comm.NowUnix()
	list := strings.Split(codes, "\n")
	sucNum := 0
	errNum := 0
	for _, code := range list {
		code := strings.TrimSpace(code)
		if code != "" {
			data := &models.LtCode{
				GiftId:     giftId,
				Code:       code,
				SysCreated: now,
			}
			err := c.ServiceCode.Create(data)
			if err != nil {
				errNum++
			} else {
				sucNum++
			}
		}
	}
	c.Ctx.HTML(fmt.Sprintf("成功导入%d条，调入失败%d条，<a href='/admin/code?gift_code=%d'>返回</a>", sucNum, errNum, giftId))
}

func (c *AdminCodeController) GetDelete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.ServiceCode.Delete(id)
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/admin/code"
	}
	return mvc.Response{
		Path: refer,
	}
}

func (c *AdminCodeController) GetReset() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.ServiceCode.Update(&models.LtCode{Id: id, SysStatus: 0}, []string{"sys_status"})
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/admin/code"
	}
	return mvc.Response{
		Path: refer,
	}
}

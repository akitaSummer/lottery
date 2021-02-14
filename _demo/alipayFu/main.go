/**
 * 五福
 * /lucky只有一个抽奖接口
 * 压力测试
 * wrk -t10 -c10 -d5 http://localhost:8080/lucky
 */
package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type gift struct {
	id      int    // 奖品ID
	name    string // 奖品名称
	pic     string // 奖品的图片
	link    string // 奖品类型
	inuse   bool   // 是否命中
	rate    int    // 中奖概率 万分之N
	rateMin int    // 大于等于中奖编码
	rateMax int    // 小于等于中奖编码
}

// 最大中奖号码
const rateMax = 10000

var logger *log.Logger

// 奖品列表
var giftList []*gift

var mu sync.Mutex

type lotteryController struct {
	Ctx iris.Context
}

// 初始化日志
func initLog() {
	f, _ := os.Create("./_demo/wechatShake/lottery_demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	initLog()

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func newGift() *[5]gift {
	giftList := new([5]gift)
	g1 := gift{
		id:      1,
		name:    "富强福",
		pic:     "",
		link:    "",
		inuse:   false,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[0] = g1
	g2 := gift{
		id:      2,
		name:    "和谐福",
		pic:     "",
		link:    "",
		inuse:   false,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[1] = g2
	g3 := gift{
		id:      3,
		name:    "友善福",
		pic:     "",
		link:    "",
		inuse:   false,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[2] = g3
	g4 := gift{
		id:      4,
		name:    "爱国福",
		pic:     "",
		link:    "",
		inuse:   false,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[3] = g4
	g5 := gift{
		id:      5,
		name:    "敬业福",
		pic:     "",
		link:    "",
		inuse:   false,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[4] = g5
	return giftList
}

func giftRage(rage string) *[5]gift {
	giftList := newGift()
	rates := strings.Split(rage, ",")
	ratesLen := len(rates)
	rateStart := 0
	for i, data := range giftList {
		if !data.inuse {
			continue
		}
		grate := 0
		if i < ratesLen {
			grate, _ = strconv.Atoi(rates[i])
		}
		giftList[i].rate = grate
		giftList[i].rateMin = rateStart
		giftList[i].rateMax = rateStart + grate
		if giftList[i].rateMax >= rateMax {
			giftList[i].rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += grate
		}
	}
	fmt.Printf("giftList=%v\n", giftList)
	return giftList
}

// GET http://localhost:8080/?rate=4,3,2,1,0
func (c *lotteryController) Get() string {
	rate := c.Ctx.URLParamDefault("rate", "4, 3, 2, 1, 0")
	giftList := giftRage(rate)
	return fmt.Sprintf("%v\n", giftList)
}

// http://localhost:8080/lucky
func (c *lotteryController) GetLucky() map[string]interface{} {
	uid, _ := c.Ctx.URLParamInt("uid")
	rate := c.Ctx.URLParamDefault("rate", "4, 3, 2, 1, 0")
	code := lucyCode()
	ok := false
	result := make(map[string]interface{})
	result["success"] = ok
	giftList := giftRage(rate)

	for _, data := range giftList {
		if !data.inuse {
			continue
		}
		if data.rateMin <= int(code) && data.rateMax > int(code) {
			sendData := data.name
			if ok {
				// 中奖后，成功得到奖品
				// 生成中奖记录
				saveLuckyData(code, data.id, data.name, data.link, sendData)
				result["success"] = ok
				result["id"] = data.id
				result["name"] = data.name
				result["link"] = data.link
				result["data"] = sendData
				result["uid"] = uid
				break
			}
		}
	}

	return result
}

func lucyCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(int32(rateMax))
	return code
}

func saveLuckyData(code int32, id int, name string, link string, sendData string) {
	logger.Printf("lucky, code=%d, gift=%d, name=%d, link=%s, data=%s,", code, id, name, link, sendData)
}

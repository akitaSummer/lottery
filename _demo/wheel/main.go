// wrk -t10 -c100 -d5 "http://localhost:8080/prize"
package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// 奖品中奖概率
type Prate struct {
	Rate  int    // 万分之N的中奖率
	Total int    // 总数量为限制，0则表示无限数量
	CodeA int    // 中奖概率起始编码
	CodeB int    // 中奖概率终止编码
	Left  *int32 // 剩余数
}

// 奖品列表
var prizeList []string = []string{
	"一等奖",
	"二等奖",
	"三等奖",
	"",
}

var mu sync.Mutex = sync.Mutex{}
var left = int32(1000)

// 奖品的中奖概率，与prizeList对应的设置
var rateList []Prate = []Prate{
	Prate{1, 1, 0, 0, &left},
	//Prate{2, 2, 1, 2, 2},
	//Prate{5, 10, 3, 5, 10},
	//Prate{100, 0, 0, 9999, 0},
}

// 抽奖控制器
type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

// http://localhost:8080/
func (c *lotteryController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return fmt.Sprintf("大转盘奖品列表: <br/> %s", strings.Join(prizeList, "<br/>\n"))
}

func (c *lotteryController) GetDebug() string {
	return fmt.Sprintf("获奖概率: %v\n", rateList)
}

func (c *lotteryController) GetPrize() string {
	// 第一步 抽奖 根据随机数匹配奖品
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	code := r.Intn(10000)
	var myprize string
	var prizeRate *Prate
	// 从奖品列表匹配是否中奖
	for i, prize := range prizeList {
		rate := &rateList[i]
		if code >= rate.CodeA && code <= rate.CodeB {
			// 满足条件
			myprize = prize
			prizeRate = rate
			break
		}
	}
	if myprize == "" {
		myprize = "很遗憾，再来一次吧"
		return myprize
	}
	if prizeRate.Total == 0 {
		return myprize
	} else if *prizeRate.Left > 0 {
		//mu.Lock()
		//defer mu.Unlock()
		left := atomic.AddInt32(prizeRate.Left, -1)
		if left >= 0 {
			return myprize
		}
		myprize = "很遗憾，再来一次吧"
		return myprize
	} else {
		myprize = "很遗憾，再来一次吧"
		return myprize
	}
}

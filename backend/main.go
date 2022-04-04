package main

import (
	"backend/config"
	"backend/dao"
	"backend/datasource"
	"backend/route"
	"backend/ticker"
	"backend/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	//读取配置文件
	config.Setup()
	//读取数据库连接池
	dao.Setup()
	//ticker配置
	go func() {
		t := time.NewTicker(time.Second * time.Duration(config.SystemConfig.TickerTimeInterval))
		for {
			select {
			case <-t.C:
				ticker.CheckContestAndSendEmailTicker()
			}
		}
	}()

	//gin启动
	router := gin.Default() // 获取 engine
	router.Use(route.Cors())
	router.GET("/ping", func(c *gin.Context) {
		var header map[string]string = make(map[string]string)
		header["AU"] = "au"
		var res = utils.Get("https://apps.game.qq.com/lol/match/apis/searchBMatchInfo_bak.php?p8=5&p1=167&p4=1&p2=%C8%AB%B2%BF&p9=&p10=&p6=3&p11=&p12=&page=1&pagesize=10&r1=retObj&_=1648892356929", header)
		res = res[11 : len(res)-1]
		//res = res[0: len(res) - 1]
		str := []byte(res)
		fmt.Println(res)
		resObj := datasource.ContestAPIResponse{}
		err := json.Unmarshal(str, &resObj)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resObj)
		c.String(http.StatusOK, res)
	})

	router.GET("/testSendEmail", func(c *gin.Context) {
		utils.SendEmail("289348588@qq.com;3483131756@qq.com", "测试发送邮件", "GO 发送邮件，官方连包都帮我们写好了，真是贴心啊！！！")
		c.String(http.StatusOK, "")
	})

	router.POST("/sendTest", func(c *gin.Context) {
		var req dao.Booking
		if err := c.BindJSON(&req); err != nil {
			fmt.Print(err)
			c.JSON(http.StatusBadRequest, err)
			return
		}
		fmt.Println(req)
		db := dao.GetDB()
		db.Create(&req)
		c.JSON(http.StatusOK, "sendTest")
	})

	//获取比赛
	router.GET("/getContestList", route.GetContestList)

	//获取比赛详情（可能不需要）
	router.GET("/getContestDetail", route.GetContestDetail)

	//获取预约信息
	router.GET("/getBookingList", route.GetBookingList)

	//发起预约
	router.POST("/startBooking", route.StartBooking)

	//取消预约
	router.DELETE("/cancelBooking", route.CancelBooking)

	err := router.Run(":8081") // 指定端口，运行 Gin
	if err != nil {
		log.Panicln("服务器启动失败：", err.Error())
	}
}

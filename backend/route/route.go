package route

import (
	"backend/dao"
	"backend/datasource"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

//跨域配置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session, x-custom-header")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}

//获取当前可预约比赛
func GetContestList(c *gin.Context) {
	var contestList = datasource.GetLatestContestInLPL()
	jsonStu, err := json.Marshal(contestList)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	fmt.Println(string(jsonStu))

	c.String(http.StatusOK, string(jsonStu))
}

//获取当前预约的比赛的详情
func GetContestDetail(c *gin.Context) {
	//contestId := c.Query("contestId")
	//res := datasource.GetContestDetailStr(contestId)

	filePtr, err := ioutil.ReadFile("test/contestDetail.txt") //config的文件目录
	if err != nil {

	}
	//fmt.Println(string(filePtr))
	c.String(http.StatusOK, string(filePtr))
}

//获取预约列表
func GetBookingList(c *gin.Context) {
	var bookingList []dao.Booking
	db := dao.GetDB()
	db.Where(" booking_status in  (?)", []int{1, 2, 3}).Find(&bookingList)
	jsonStu, err := json.Marshal(bookingList)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	fmt.Println(string(jsonStu))
	//c.String(http.StatusOK, string(jsonStu))
	c.String(http.StatusOK, string(jsonStu))
}

//开始预约请求
func StartBooking(c *gin.Context) {
	var req dao.Booking
	if err := c.BindJSON(&req); err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(req)
	req.BookingStatus = 1
	req.ResultScoreA = 0
	req.ResultScoreB = 0
	db := dao.GetDB()
	db.Create(&req)

	c.String(http.StatusOK, "ok")
}

//删除预约请求
func CancelBooking(c *gin.Context) {
	contestId := c.Query("bookingId")
	db := dao.GetDB()
	db.Where("booking_id = ?", contestId).Delete(&dao.Booking{})
	c.String(http.StatusOK, "ok")
}

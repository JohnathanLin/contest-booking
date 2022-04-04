package email

import (
	"backend/config"
	"backend/dao"
	"backend/utils"
	"fmt"
)

// SendBookingSuccEmail 预约成功发送邮件
func SendBookingSuccEmail(booking dao.Booking) {
	var subject = fmt.Sprintf("您预约的比赛 %s 预约成功了", booking.ContestName)
	var content = fmt.Sprintf("您预约的比赛 %s 预约成功了, 当前比分：%s %d:%d %s ", booking.ContestName,
		booking.TeamA, booking.BookingScoreA, booking.BookingScoreB, booking.TeamB)
	if config.SystemConfig.TestMode == 0 {
		utils.SendEmail(booking.BookingEmail, subject, content)
	} else {
		fmt.Println(subject + content)
	}
}

// SendBookingFailEmail 预约失败发送邮件
func SendBookingFailEmail(booking dao.Booking) {
	var subject = fmt.Sprintf("您预约的比赛 %s 预约失败了", booking.ContestName)
	var content = fmt.Sprintf("您预约的比赛 %s 预约失败了, 目标比分： %s %d:%d %s, 当前比分：%s %d:%d %s ", booking.ContestName,
		booking.TeamA, booking.BookingScoreA, booking.BookingScoreB, booking.TeamB,
		booking.TeamA, booking.ResultScoreA, booking.ResultScoreA, booking.TeamB)
	//utils.SendEmail(booking.BookingEmail, subject, content)
	if config.SystemConfig.TestMode == 0 {
		utils.SendEmail(booking.BookingEmail, subject, content)
	} else {
		fmt.Println(subject + content)
	}
}

package ticker

import (
	"backend/dao"
	"backend/datasource"
	"backend/email"
	"fmt"
)

func CheckContestAndSendEmailTicker() {
	db := dao.GetDB()
	var bookings []dao.Booking
	db.Where("booking_status in (?) and contest_time < now()", []int{1, 2}).Find(&bookings)
	var total int = 0
	var succ int = 0
	var fail int = 0
	for _, booking := range bookings {
		var dataChanged = false
		if booking.BookingStatus == 1 { //如果当前时间已经在比赛开始时间之后，则更新比赛状态为进行中
			booking.BookingStatus = 2
			dataChanged = true
		}
		var contestDetailRes = datasource.GetContestDetail(booking.ContestId)
		if contestDetailRes.Success == false {
			continue
		}
		var contestDetail = contestDetailRes.Data
		if contestDetail.ScoreA != booking.ResultScoreA {
			booking.ResultScoreA = contestDetail.ScoreA
			dataChanged = true

		}
		if contestDetail.ScoreB != booking.ResultScoreB {
			booking.ResultScoreB = contestDetail.ScoreB
			dataChanged = true
		}

		if booking.ResultScoreA > booking.BookingScoreA || booking.ResultScoreB > booking.BookingScoreB {
			booking.BookingStatus = 4
			dataChanged = true
			email.SendBookingFailEmail(booking)
			fail++
		}

		if booking.ResultScoreA == booking.BookingScoreA && booking.ResultScoreB == booking.BookingScoreB {
			booking.BookingStatus = 3
			dataChanged = true
			email.SendBookingSuccEmail(booking)
			succ++
		}
		if dataChanged {
			db.Save(&booking)
		}
		total++
	}
	str := fmt.Sprintf("更新了%d个预约信息，其中发送了%d个成功邮件 %d个失败邮件", total, succ, fail)
	fmt.Println(str)

}

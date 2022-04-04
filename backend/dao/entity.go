package dao

import "backend/utils"

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u Booking) TableName() string {
	//绑定MYSQL表名为users
	return "bookings"
}

//预约信息
type Booking struct {
	BookingId     int             `gorm:"primarykey" json:"bookingId"`
	ContestId     int             `gorm:"column:contest_id" json:"contestId"`
	ContestName   string          `gorm:"column:contest_name" json:"contestName"`
	BookingStatus int             `gorm:"column:booking_status" json:"bookingStatus"` // 1未开始 2计算中 3预约成功 4预约失败
	ContestTime   utils.LocalTime `gorm:"column:contest_time" json:"contestTime" time_format:"2006-01-02 12:34:56" time_utc:"8" `
	TeamA         string          `gorm:"column:team_a" json:"teamA"`
	TeamB         string          `gorm:"column:team_b" json:"teamB"`
	BookingScoreA int             `gorm:"column:booking_score_a" json:"bookingScoreA"`
	BookingScoreB int             `gorm:"column:booking_score_b" json:"bookingScoreB"`
	ResultScoreA  int             `gorm:"column:result_score_a" json:"resultScoreA"`
	ResultScoreB  int             `gorm:"column:result_score_b" json:"resultScoreB"`
	BookingEmail  string          `gorm:"column:booking_email" json:"bookingEmail"`
}

package config

import (
	"encoding/json"
	"os"
)

// Config 配置对象
type Config struct {
	Database   *Database   `json:"database"`   //数据库配置
	EmailConf  *EmailConf  `json:"emailConf"`  //发送邮件配置
	SystemConf *SystemConf `json:"systemConf"` //系统配置
}

// GlobalConfigSetting 配置实例
var GlobalConfigSetting = &Config{}

// Database 数据库配置对象
type Database struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Name        string `json:"name"`
	TablePrefix string `json:"table_prefix"`
}

//发送邮件配置
type EmailConf struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

//系统配置
type SystemConf struct {
	TestMode           int   `json:"testMode"`           //1测试 0正式
	TickerTimeInterval int64 `json:"tickerTimeInterval"` //ticker时间间隔 秒
}

// Setup 配置
func Setup() {
	filePtr, err := os.Open("config/config.json") //config的文件目录
	if err != nil {
		return
	}
	defer filePtr.Close()
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(GlobalConfigSetting)
	DatabaseSetting = GlobalConfigSetting.Database
	EmailConfig = GlobalConfigSetting.EmailConf
	SystemConfig = GlobalConfigSetting.SystemConf
}

// DatabaseSetting 数据库配置对象 实例
var DatabaseSetting = &Database{}
var EmailConfig = &EmailConf{}
var SystemConfig = &SystemConf{}

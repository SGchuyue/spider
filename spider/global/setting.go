package global

import (
	"myproject.cn/spider/common/logger"
	"myproject.cn/spider/http_server/config"
)

//配置并声明全居变量，方便关联
var (
	ServerSetting   *config.ServerSettingS
	AppSetting      *config.AppSettingS
	DatabaseSetting *config.DatabaseSettingS

	Logger *logger.Logger
)

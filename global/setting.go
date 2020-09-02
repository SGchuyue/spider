package global

import (
	"writetest/pkg/logger"
	"writetest/pkg/setting"
)

//配置并声明全居变量，方便关联
var (
	ServerSetting *setting.ServerSettingS
	AppSetting     *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	Logger *logger.Logger
)

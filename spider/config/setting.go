package config

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

// 配置文件读取封装设定viper
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	// 允许设置多个配置路径，尝试解决路径查找问题
	vp.AddConfigPath("configs/")

	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

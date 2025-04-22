package config

import (
	"github.com/spf13/viper"
	"log"
)

// Config 获取配置文件中配置信息的类
type Config struct {
	App struct {
		Name string
		Port string
	}

	Database struct {
		Host     string
		Port     string
		Username string
		Password string
		scheme   string
	}
}

// AppConfig Config类型的指针变量
var AppConfig *Config

// InitConfig 初始化config.yml中的配置信息
func InitConfig() {
	viper.SetConfigName("config")   // 设置配置文件名
	viper.SetConfigType("yml")      // 配置文件类型
	viper.AddConfigPath("./config") //配置文件位置

	// 加载配置信息
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config：%v", err)
	}
	AppConfig = &Config{}
	// 将 Viper 当前加载和合并的所有配置数据解组（unmarshal）到一个指定的结构体变量
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("unable to decode into AppConfig struct：%v", err)
	}

	// 初始化数据库连接
	InitDb()
}

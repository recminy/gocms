package config

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var Viper *viper.Viper

type Map map[string]interface{}

func init()  {
	Viper = viper.New()
	//设置配置文件名
	Viper.SetConfigName(".env")
	//配置文件类型
	Viper.SetConfigType("env")
	//配置文件路径
	Viper.AddConfigPath(".")

	err := Viper.ReadInConfig()
	if err != nil {
		fmt.Println("配置读取失败", err)
	}
	viper.SetEnvPrefix("appEnv")
	Viper.AutomaticEnv()
}

func Add(name string, configuration map[string]interface{}) {
	Viper.Set(name, configuration)
}

func Get(path string, defaultValue ...interface{}) interface{} {
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0  {
			return defaultValue[0]
		}
		return nil
	}
	return Viper.Get(path)
}

func Env(path string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(path, defaultValue[0])
	}
	return Get(path)
}

func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue...))
}

func GetUint(path string, defaultValue ...interface{}) uint  {
	return cast.ToUint(Get(path, defaultValue...))
}

func GetBool(path string, defaultValue ...interface{}) bool  {
	return cast.ToBool(Get(path, defaultValue...))
}
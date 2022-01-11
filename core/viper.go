package core

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var v *viper.Viper

type StrMap map[string]interface{}

func init() {

	v = viper.New()

	//设置文件名
	v.SetConfigName(".env")

	//文件类型
	v.SetConfigType("env")

	//文件路径
	v.AddConfigPath(".")

	//读根目录下的 .env 文件，读不到会报错
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Viper.Get() 时, 优先读取环境变量
	v.AutomaticEnv()
}

func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}
	return Get(envName)
}

func Get(path string, defaultValue ...interface{}) interface{} {
	if !v.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return v.Get(path)
}

func Set(name string, configuration map[string]interface{}) {
	v.Set(name, configuration)
}

func GetToString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

func GetToInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

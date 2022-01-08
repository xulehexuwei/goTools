package config_log

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func GetConf() func(string) interface{} {
	viper.SetConfigName("conf")     //把json文件换成yaml文件，只需要配置文件名 (不带后缀)即可
	viper.AddConfigPath("./config") //添加配置文件所在的路径
	//viper.SetConfigType("json")       //设置配置文件类型
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}

	viper.WatchConfig() //监听配置变化
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置发生变更：", e.Name)
	})
	return viper.Get
}

//
//func main() {
//	conf := config_log.GetConf()
//	r := conf("mysql.host")
//	fmt.Println(r)
//}

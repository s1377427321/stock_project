package main

import (
	"github.com/json-iterator/go/extra"
	"web_frame/datacache"
	"web_frame/util/config"
	"github.com/spf13/pflag"
	"time"
	"web_frame/datacache/commoncache"
	"github.com/astaxie/beego/logs"
)

var (
	cfg = pflag.StringP("config", "c", "", "config file path.")
)

func main() {

	//这里是设置，支持模糊的json匹配，比如说把一个“1.29”string类型，可以转1.29 float类型
	//更多详情请移步：https://www.colabug.com/233300.html
	extra.RegisterFuzzyDecoders()
	// 导入配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	datacache.InitDataCache()

	go dotest()
	for range time.NewTicker(1 * time.Hour).C {

	}
}

func dotest() {
	//commoncache.SaveIDPublishTime(123, time.Now().Unix())

	//time.Sleep(5 * time.Second)

	before := commoncache.GetIDPublishTime(123)

	result := time.Now().Unix() - before

	hour := int(result) / 3600
	min := (int(result) % 3600) / 60
	second := (int(result) % 3600) % 60

	logs.Info("  ", hour, min, second)

}

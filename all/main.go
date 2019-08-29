package all

import (
	"github.com/json-iterator/go/extra"
	"web_frame/util/config"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "config file path.")
)

func main()  {
	extra.RegisterFuzzyDecoders()
	// 导入配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}


}
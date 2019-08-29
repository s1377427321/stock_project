package baseconfig

import (
	"github.com/spf13/viper"
	"github.com/spf13/pflag"
	"github.com/leisc/aliyun-acm"
	"github.com/lexkong/log"
	"bytes"
)

var (
	configStr     = ""
	runtime_viper *viper.Viper
	callBack      func()
	diamond       *aliacm.Diamond
)

var (
	cfg         = pflag.StringP("config", "c", "", "server config file path.")
	config_flag = false
)

func init() {
	runtime_viper = viper.New()
}

func GetViper() *viper.Viper {
	return runtime_viper
}

func StartConfigTask(callback func()) {
	callBack = callback
	if err := Init(*cfg); err != nil {
		panic(err)
	}
	GetACMConfig(viper.GetString("acm.namespace"), viper.GetString("acm.accessKey"), viper.GetString("acm.secretKey"), viper.GetString("acm.group"), viper.GetString("acm.dataId"))
}

func GetACMConfig(namespace, accessKey, secretKey, group, dataId string) {
	d, err := aliacm.New(
		aliacm.SZAddr,
		namespace,
		accessKey,
		secretKey)
	if err != nil {
		log.Errorf(err, "new acm failed")
		return
	}
	var f = func(h aliacm.Unit, err error) {
		log.Errorf(err, "hook faile")
	}
	d.SetHook(f)
	unit := aliacm.Unit{
		Group:     group,
		DataID:    dataId,
		FetchOnce: false,
		OnChange:  Handle,
	}
	d.Add(unit)
	diamond = d
	log.Infof("start acm task success=======")
	select {}
}

func Handle(config aliacm.Config) {
	log.Infof("===============================Handle new config\n%s", string(config.Content))

	runtime_viper.SetConfigType("yaml")
	err := runtime_viper.ReadConfig(bytes.NewReader(config.Content))
	if err != nil {
		log.Errorf(err, "viper read config string failed")
	}
	if !config_flag {
		callBack()
		config_flag = true
	}
}

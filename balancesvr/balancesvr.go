package main

import (
	"github.com/lexkong/log"
	"gobaselib/baseconfig"
)

func main() {
	log.InitWithFile("./conf/logcfg.yaml")
	baseconfig.StartConfigTask(ConfigCallBack)
}

func ConfigCallBack() {
	log.Infof("------ ",baseconfig.GetViper().GetString("acm.dataId"))
}
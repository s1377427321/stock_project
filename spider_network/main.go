package main

import (
	"spider_network/spider_struct"
	"time"
)

const(
	//webUrl ="http://quote.eastmoney.com/sh511990.html"
	//webUrl="https://www.msn.com/en-us/money/etfdetails/fi-136.8.511990?ocid=INSFIST10"
	webUrl = "https://www.jisilu.cn/data/etf/#index"
)


func main()  {
	fund511:=&spider_struct.Spider{}
	fund511.Url = webUrl

	fund511.Start(5*time.Second)
	if fund511.GetIsSendEmail() {

	}

	for range time.NewTicker(1000*time.Minute).C{

	}
}


package main

import (
	"time"
	. "email_notice/common"
	"sync"
	"common/email"
)

var mainUrl = "http://hq.sinajs.cn/list=%s"
var httpPort = ":4661"
var emailUrl = "1377427321@qq.com"

var NoticeStockS map[string]*Stock
var mx sync.Mutex

var BuyStocks map[string]*BuyStock
var bmx sync.Mutex

var emailServer = &email.EmailServers{
	ServerEmail:    "1377427321@qq.com",
	ServerPassword: "atpncirernxrhchj",
	ServerPort:     465,
	ServerIP:       "smtp.qq.com",
}

func NoticeEmail(s *Stock, content string) {
	ec := &email.EmailContent{
		NickName:     "文健",
		Subject:      content,
		BodyContent:  content,
		NoticeEmails: []string{emailUrl},
	}
	s.SendEmail(emailServer, ec)
}

func init() {
	NoticeStockS = make(map[string]*Stock, 0)
	BuyStocks = make(map[string]*BuyStock, 0)

	DoInitStock()
}

//这个程序负债监控股票价格，设置一个最高价格，最低价格，到了这个价格，就会通知用户，去操作
func main() {
	go RunHttpServer()
	for range time.NewTicker(100 * time.Hour).C {
	}
}


var InitNoticeStocks []InitNoticeStocksS = []InitNoticeStocksS{
	InitNoticeStocksS{
		Money:130000,
		Name:"华宝添益",
		Code :"sh511990",
		Height:100.05,
		Low:99.905,
	},
}

var InitBuyStocks []InitBuyStocksS = []InitBuyStocksS{
	InitBuyStocksS{
		Name:"白云机场",
		Code:"sh600004",
		Money:130000,
		Copies:20,
		Price:12.21,
	},
	InitBuyStocksS{
		Name:"伊利股份",
		Code:"sh600887",
		Money:130000,
		Copies:10,
		Price:25.18,
	},
	InitBuyStocksS{
		Name:"招商银行",
		Code:"sh600036",
		Money:130000,
		Copies:10,
		Price:29.32,
	},
	InitBuyStocksS{
		Name:"复星医药",
		Code:"sh600004",
		Money:130000,
		Copies:10,
		Price:25.76,
	},
	InitBuyStocksS{
		Name:"贵州茅台",
		Code:"sh600519",
		Money:130000,
		Copies:10,
		Price:717.92,
	},
}

func DoInitStock()  {
	for _,v:=range InitNoticeStocks{
		AddNoticeStock(v.Code,v.Height,v.Low,v.Money)
	}

	for _,v:=range InitBuyStocks{
		AddBuyStock(v.Code,v.Price,v.Money,v.Copies)
	}
}




/*
华宝添益
120.79.154.53:4661/add?code=sh511990&height=100.05&low=99.905&money=130000
白云机场
120.79.154.53:4661/addstock?code=sh600004&money=130000&copies=20&price=12.21
伊利股份
120.79.154.53:4661/addstock?code=sh600887&money=130000&copies=10&price=25.18
招商银行
120.79.154.53:4661/addstock?code=sh600036&money=130000&copies=10&price=29.32
复星医药
120.79.154.53:4661/addstock?code=sh600196&money=130000&copies=10&price=25.76
贵州茅台
120.79.154.53:4661/addstock?code=sh600519&money=130000&copies=10&price=717.92
*/


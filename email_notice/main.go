package main

import (
	. "email_notice/common"
	"sync"
	"common/email"
	"common"
	"time"
	"github.com/spf13/viper"
	"strings"
	"strconv"
	"fmt"
)

var mainUrl = "http://hq.sinajs.cn/list=%s"
var httpPort = ":5555"
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

var NoticeLimit = 1

func init() {
	common.CreateBeegoLog("log_email_notice")

	NoticeStockS = make(map[string]*Stock, 0)
	BuyStocks = make(map[string]*BuyStock, 0)

}

//这个程序负债监控股票价格，设置一个最高价格，最低价格，到了这个价格，就会通知用户，去操作
func main() {

	common.StartConfigTask("conf/email_notice.yaml", CallBack)

	go RunHttpServer()

	DoInitStock()

	//StopWinLose
	//swl:=&StopWinLose{}
	//swl.Init("000895","双汇发展",20.05,1)

	for range time.NewTicker(31 * time.Minute).C {
		h := time.Now().Hour()
		if h >= 9 && h < 10 {
			noticeIsAlive()
		}

		if h >= 15 && h < 16 {
			noticeIsAlive()
		}
	}
}

func CallBack() {
	noticeStocks := viper.GetStringSlice("notice_stocks")
	for _, v := range noticeStocks {
		data := strings.Split(v, "|")
		money, _ := strconv.ParseFloat(data[0], 64)
		name := data[1]
		code := data[2]
		heigh, _ := strconv.ParseFloat(data[3], 64)
		low, _ := strconv.ParseFloat(data[4], 64)

		DeleteNoticeStock(code)

		fmt.Println(fmt.Sprintf("%v:%v:%v:%v:%v", money, name, code, heigh, low))
		AddNoticeStock(code, heigh, low, money)
	}

	stopwinloses := viper.GetStringSlice("stop_win_lose")
	for _, v := range stopwinloses {
		data := strings.Split(v, "|")

		code := data[0]
		name := data[1]
		price, _ := strconv.ParseFloat(data[2], 64)
		magnification, _ := strconv.Atoi(data[3])

		InstanceStopWinLoseManage().DeleteStock(code)

		InstanceStopWinLoseManage().Add(code, name, price, magnification)
	}
}

func noticeIsAlive() {
	ec := &email.EmailContent{
		NickName:     "文健",
		Subject:      "alive",
		BodyContent:  "alive",
		NoticeEmails: []string{emailUrl},
	}
	email.SendEmailTo(emailServer, ec)
	common.DingDingNotify1("alive")
}

var InitNoticeStocks []InitNoticeStocksS = []InitNoticeStocksS{
	InitNoticeStocksS{
		Money:  130000,
		Name:   "华宝添益",
		Code:   "sh511990",
		Height: 100.05,
		Low:    99.95,
	},
	//InitNoticeStocksS{
	//	Money:  130000,
	//	Name:   "R001",
	//	Code:   "sz131810",
	//	Height: 6,
	//	Low:    0,
	//},
	//InitNoticeStocksS{
	//	Money:  130000,
	//	Name:   "三全食品",
	//	Code:   "sz002216",
	//	Height: 8,
	//	Low:    5.3,
	//},
}

var InitBuyStocks []InitBuyStocksS = []InitBuyStocksS{
	//InitBuyStocksS{
	//	Name:   "白云机场",
	//	Code:   "sh600004",
	//	Money:  130000,
	//	Copies: 20,
	//	Price:  13.15,
	//},
	//InitBuyStocksS{
	//	Name:   "伊利股份",
	//	Code:   "sh600887",
	//	Money:  130000,
	//	Copies: 10,
	//	Price:  27.81,
	//},
	//InitBuyStocksS{
	//	Name:   "招商银行",
	//	Code:   "sh600036",
	//	Money:  130000,
	//	Copies: 10,
	//	Price:  32.9,
	//},
	//InitBuyStocksS{
	//	Name:   "复星医药",
	//	Code:   "sh600004",
	//	Money:  130000,
	//	Copies: 10,
	//	Price:  25.76,
	//},
	//InitBuyStocksS{
	//	Name:   "贵州茅台",
	//	Code:   "sh600519",
	//	Money:  130000,
	//	Copies: 10,
	//	Price:  717.92,
	//},
	//InitBuyStocksS{
	//	Name:   "比亚迪",
	//	Code:   "sz002594",
	//	Money:  130000,
	//	Copies: 20,
	//	Price:  53.80,
	//},
	//
	//
	////ETF
	//InitBuyStocksS{
	//	Name:   "50ETF",
	//	Code:   "sh510050",
	//	Money:  130000,
	//	Copies: 20,
	//	Price:  2.802,
	//},
	//
	//InitBuyStocksS{
	//	Name:   "300ETF",
	//	Code:   "sh510300",
	//	Money:  130000,
	//	Copies: 20,
	//	Price:  3.735,
	//},
	//
	//InitBuyStocksS{
	//	Name:   "证券ETF",
	//	Code:   "sh512880",
	//	Money:  130000,
	//	Copies: 20,
	//	Price:  1.038,
	//},
	//
	//InitBuyStocksS{
	//	Name:   "环保ETF",
	//	Code:   "sh512580",
	//	Money:  130000,
	//	Copies: 20,
	//	Price:  0.763,
	//},
	//
	//InitBuyStocksS{
	//	Name:   "银行ETF",
	//	Code:   "sh512800",
	//	Money:  130000,
	//	Copies: 20,
	//	Price:  1.070,
	//},
	//
	//InitBuyStocksS{
	//	Name:   "医药ETF",
	//	Code:   "sh512010",
	//	Money:  130000,
	//	Copies: 20,
	//	Price:  1.577,
	//},
}

func DoInitStock() {
	for _, v := range InitNoticeStocks {
		AddNoticeStock(v.Code, v.Height, v.Low, v.Money)
	}

	for _, v := range InitBuyStocks {
		AddBuyStock(v.Code, v.Price, v.Money, v.Copies)
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
/*
120.79.154.53:4661/add?code=sh000001&height=3056&low=3005&money=130000

http://hq.sinajs.cn/list=sh000001

 */

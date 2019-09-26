package common

import (
	"sync"
	"time"
	"fmt"
	"common/email"
	"common"
)

var UpdataTicket = 5 * time.Second

type NoticeEmailCallBack func(s *Stock, content string)

type StockInterface interface {
	SendEmail(es *email.EmailServers, ec *email.EmailContent)
	UpdateCurrentPrice() (isSendEmail bool, emialContent string)
	Close()
	Start()
}

type Average struct {
	Day5 float64
}

type StockCommon struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	AllMoney  float64 `json:"all_money"`
	BuyPrice  float64 `json:"buy_price"`
	BuyMoney  float64 `json:"buy_money"`
	SellPrice float64 `json:"sell_price"`
	BuyNums   int `json:"buy_nums"`
}

type Stock struct {
	BuyMoney       float64
	Code           string
	Url            string
	Name           string
	HightPrice     float64
	LowPrice       float64
	mx             sync.Mutex
	NoticeCallBack NoticeEmailCallBack
	closeCh        chan interface{}
	NoticeLimit    int
	Count          int
	CountOverTime  time.Time
}

func (s *Stock) Start() {
	s.closeCh = make(chan interface{}, 0)
	s.UpdateCurrentPrice()
	go func() {
		for range time.NewTicker(UpdataTicket).C {

			select {
			case <-s.closeCh:
				return
			default:
			}

			s.UpdateCurrentPrice()

		}
	}()
}

func (s *Stock) Close() {
	close(s.closeCh)
}

func (s *Stock) SendEmail(es *email.EmailServers, ec *email.EmailContent) {
	s.Count ++
	if s.Count > s.NoticeLimit {
		//一小时以后恢复发送邮件
		addTime := s.CountOverTime.Add(2 * time.Hour)
		//addTime:=s.CountOverTime.Add(2*time.Second)
		if time.Now().After(addTime) {
			s.Count = 0
		}
		return
	} else {
		//记录发送了的时间
		s.CountOverTime = time.Now()
	}
	email.SendEmailTo(es, ec)
	common.DingDingNotify1(ec.BodyContent)
}

func (s *Stock) SendToDingDing(content string) {
	common.DingDingNotify1(content)
}

func (s *Stock) UpdatePrice(heigh, low float64) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.HightPrice = heigh
	s.LowPrice = low
}

func (s *Stock) UpdateCurrentPrice() (isSendEmail bool, emialContent string) {

	h := time.Now().Hour()
	if h < 9 || h >= 15 {
		return
	}

	s.mx.Lock()
	defer s.mx.Unlock()
	currentPrice, err := common.GetPriceFromUrl(s.Url)
	if err != nil || currentPrice < 0.001 {
		return false, ""
	}

	strFormat := "%s current price : %v  and  height price：%v, low price：%v ,buy money:%v , buy copies:%v"
	var content = fmt.Sprintf(strFormat, s.Code, currentPrice, s.HightPrice, s.LowPrice, s.BuyMoney, 0)
	fmt.Println(content)

	if currentPrice >= s.HightPrice {
		bc := (int(s.BuyMoney/s.HightPrice) / 100) * 100
		var content = fmt.Sprintf(strFormat, s.Code, currentPrice, s.HightPrice, s.LowPrice, s.BuyMoney, bc)
		if s.NoticeCallBack != nil {
			s.NoticeCallBack(s, content)
		}
		s.SendToDingDing(content)
		return true, content

	} else if currentPrice <= s.LowPrice {
		bc := (int(s.BuyMoney/s.LowPrice) / 100) * 100
		var content = fmt.Sprintf(strFormat, s.Code, currentPrice, s.HightPrice, s.LowPrice, s.BuyMoney, bc)
		if s.NoticeCallBack != nil {
			s.NoticeCallBack(s, content)
		}
		s.SendToDingDing(content)
		return true, content
	}

	return false, content
}

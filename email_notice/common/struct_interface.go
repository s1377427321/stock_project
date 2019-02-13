package common

import (
	"sync"
	"time"
	"fmt"
	"common/email"
)

var UpdataTicket = 5 * time.Second

type NoticeEmailCallBack func(s *Stock, content string)

type StockInterface interface {
	SendEmail(es *email.EmailServers, ec *email.EmailContent)
	UpdateCurrentPrice() (isSendEmail bool, emialContent string)
	Close()
	Start()
}

type Stock struct {
	BuyMoney       float64
	Code           string
	Url            string
	HeightPrice    float64
	LowPrice       float64
	mx             sync.Mutex
	NoticeCallBack NoticeEmailCallBack
	closeCh        chan interface{}
	NoticeLimit    int
	Count          int
	CountOverTime time.Time
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
		addTime:=s.CountOverTime.Add(1*time.Hour)
		//addTime:=s.CountOverTime.Add(2*time.Second)
		if time.Now().After(addTime) {
			s.Count = 0
		}
		return
	}else {
		//记录发送了的时间
		s.CountOverTime = time.Now()
	}
	email.SendEmailTo(es, ec)
}

func (s *Stock) UpdateCurrentPrice() (isSendEmail bool, emialContent string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	currentPrice, err := GetPriceFromUrl(s.Url)
	if err != nil {
		return false, ""
	}

	strFormat:="%s current price : %v  and  height price：%v, low price：%v ,buy money:%v , buy copies:%v"
	var content = fmt.Sprintf(strFormat, s.Code, currentPrice, s.HeightPrice, s.LowPrice,s.BuyMoney,0)
	fmt.Println(content)

	if currentPrice >= s.HeightPrice {
		bc:= (int(s.BuyMoney/s.HeightPrice) / 100) * 100
		var content = fmt.Sprintf(strFormat, s.Code, currentPrice, s.HeightPrice, s.LowPrice,s.BuyMoney,bc)
		s.NoticeCallBack(s, content)
		return true, content

	} else if currentPrice <= s.LowPrice {
		bc:= (int(s.BuyMoney/s.LowPrice) / 100) * 100
		var content = fmt.Sprintf(strFormat, s.Code, currentPrice, s.HeightPrice, s.LowPrice,s.BuyMoney,bc)
		s.NoticeCallBack(s, content)
		return true, content
	}

	return false, content
}

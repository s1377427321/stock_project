package common

import (
	"sync"
	"time"
	"fmt"
	"common/email"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
	"errors"
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
	Code           string
	Url            string
	HeightPrice    float64
	LowPrice       float64
	mx             sync.Mutex
	NoticeCallBack NoticeEmailCallBack
	closeCh        chan interface{}
	NoticeLimit    int
	Count          int
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
		return
	}
	email.SendEmailTo(es, ec)
}

func (s *Stock) UpdateCurrentPrice() (isSendEmail bool, emialContent string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	currentPrice ,err:= GetPriceFromUrl(s.Url)
	if err!= nil {
		return false,""
	}

	var content = fmt.Sprintf("%s current price : %v  and  height price：%v, low price：%v ", s.Code, currentPrice, s.HeightPrice, s.LowPrice)
	fmt.Println(content)
	if currentPrice >= s.HeightPrice {
		s.NoticeCallBack(s, content)
		return true, content
	} else if currentPrice <= s.LowPrice {
		s.NoticeCallBack(s, content)
		return true, content
	}

	return false, content
}



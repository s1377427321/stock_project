package main

import (
	"email_notice/common"
	stockdata "fund/stock_data_save"
	datastruct "fund/stock_data_save/structs"
	"time"
	"sort"
	"fmt"
	"strconv"
	"encoding/json"
	"io/ioutil"
)

const (
	savePath = "./save.json"

	NOTHING = 0
	BUY     = 1
	SELL    = 2
	KEEP    = 3
)

type BuyOpt struct {
	BuyCode  string
	BuyMoney float64
	BuyDay   string
}

type SaveLocal struct {
	OriginMoney float64 `json:"origin_money"`
	RemainMoney float64 `json:"remain_money"`
	WinMoney    float64 `json:"win_money"`
	WinPercent  float64 `json:"win_percent"`
	BuyDay      string  `json:"buy_day"`
	ChangDay    string  `json:"chang_day"`
	Status      int     `json:"status"`

	Code     string  `json:"code"`
	BuyNums  int     `json:"buy_nums"`
	BuyMoney float64 `json:"buy_money"`
}

type TrainStock struct {
	Stock       *common.StockCommon
	OriginMoney float64 `json:"origin_money"`
	AllMoney    float64 `json:"all_money"`
	RemainMoney float64 `json:"remain_money"`
	WinMoney    float64 `json:"win_money"`
	WinPercent  float64 `json:"win_percent"`
	IndexDay    int     `json:"index_day"`
	BuyDay      string  `json:"buy_day"`
	StockData   datastruct.DialyStockInfoSlice
	len         int    `json:"len"`
	ChangDay    string `json:"chang_day"`
	Status      int    `json:"status"`
}

//func (t *TrainStock) Init(money float64) {
//
//}

func (t *TrainStock) NewBuy(money float64, opt *BuyOpt) {
	t.OriginMoney = money
	newStock := &common.StockCommon{
		Code:     opt.BuyCode,
		BuyMoney: opt.BuyMoney,
	}
	t.BuyDay = opt.BuyDay
	info := &datastruct.StockBasicInfo{
		Symbol: opt.BuyCode,
	}
	nowDay := time.Now().Format("20060102")
	t.StockData = stockdata.GetDailTradeFromCSVSlice(info, opt.BuyDay, nowDay)
	sort.Sort(sort.Reverse(t.StockData))
	t.len = len(t.StockData)
	t.Stock = newStock
	t.Stock.BuyPrice = t.StockData[t.len-1].Close
	t.Stock.BuyNums = int(opt.BuyMoney/t.Stock.BuyPrice/100) * 100
	t.Stock.BuyMoney = float64(t.Stock.BuyNums) * t.Stock.BuyPrice

	t.RemainMoney = t.OriginMoney - t.Stock.BuyMoney
	t.WinMoney = t.RemainMoney + t.Stock.BuyMoney
	t.IndexDay = 0
	t.Status = BUY
}

//再次购买
func (t *TrainStock) OldBuy(buyDay string) (bool) {
	if t.Status != NOTHING {
		fmt.Println("not equit NOTHING")
		return false
	}
	if buyDay == "" {
		fmt.Println("not buyDay")
		return false
	}
	info := &datastruct.StockBasicInfo{
		Symbol: t.Stock.Code,
	}
	nowDay := time.Now().Format("20060102")
	t.StockData = stockdata.GetDailTradeFromCSVSlice(info, buyDay, nowDay)
	sort.Sort(sort.Reverse(t.StockData))
	t.len = len(t.StockData)
	t.IndexDay = 0

	t.Stock.BuyPrice = t.StockData[t.len-1].Close
	t.Stock.BuyNums = int(t.Stock.BuyMoney/t.Stock.BuyPrice/100) * 100
	t.Status = BUY
	return true
}

func (t *TrainStock) NextDay(status int) {
	switch status {
	case SELL:
		t.Sell()
		break
	case KEEP:
		t.Keep()
		break
	}
}

func (t *TrainStock) Sell() {
	if t.Status != BUY {
		fmt.Println("not buy")
		return
	}
	if t.IndexDay >= t.len {
		panic("Sell IndexDay > len")
	}
	info := t.StockData[t.len-1-t.IndexDay]

	t.Stock.BuyMoney = float64(t.Stock.BuyNums) * info.Close
	t.AllMoney = t.Stock.BuyMoney + t.RemainMoney
	t.WinMoney = t.AllMoney - t.OriginMoney
	t.WinPercent = (t.WinMoney) / t.OriginMoney
	t.StockData = nil
	t.IndexDay = 0
	t.ChangDay = info.Date
	fmt.Println(t.ToString())
	t.Status = NOTHING

	SaveLocalFile(t)
}

func (t *TrainStock) KeepRange(endDay string) (bool) {
	if t.Status != BUY {
		fmt.Println("not NOTHING")
		return false
	}
	for {
		info := t.StockData[t.len-1-t.IndexDay]
		tt, _ := time.Parse("2006-01-02", info.Date)
		startDay := tt.Format("20060102")
		start, _ := strconv.Atoi(startDay)
		end, _ := strconv.Atoi(endDay)
		if start < end {
			t.Keep()
		} else {
			break
		}

	}

	SaveLocalFile(t)
	return true
}

func (t *TrainStock) Keep() {
	if t.Status != BUY {
		fmt.Println("not buy")
		return
	}
	t.IndexDay += 1
	if t.IndexDay >= t.len {
		panic("Buy IndexDay > len")
	}
	info := t.StockData[t.len-1-t.IndexDay]
	t.Stock.BuyMoney = float64(t.Stock.BuyNums) * info.Close
	t.AllMoney = t.Stock.BuyMoney + t.RemainMoney
	t.WinMoney = t.AllMoney - t.OriginMoney
	t.WinPercent = (t.WinMoney) / t.OriginMoney
	t.ChangDay = info.Date
	fmt.Println(t.ToString())

	SaveLocalFile(t)
}

func (t *TrainStock) ToString() (string) {
	return fmt.Sprintf("WinMoney=%.2f ,WinPercent=%.2f,ChangeDay=%s", t.WinMoney, t.WinPercent, t.ChangDay)
}

//OriginMoney float64 `json:"origin_money"`
//RemainMoney float64 `json:"remain_money"`
//WinMoney    float64 `json:"win_money"`
//WinPercent  float64 `json:"win_percent"`
//BuyDay      string  `json:"buy_day"`
//ChangDay    string `json:"chang_day"`
//Status      int    `json:"status"`
func SaveLocalFile(t *TrainStock) error {

	tt, _ := time.Parse("2006-01-02", t.ChangDay)
	BuyDay := tt.Format("20060102")

	save := &SaveLocal{
		OriginMoney: t.OriginMoney,
		RemainMoney: t.RemainMoney,
		WinMoney:    t.WinMoney,
		WinPercent:  t.WinPercent,
		BuyDay:      BuyDay,
		ChangDay:    t.ChangDay,
		Status:      t.Status,
		Code:        t.Stock.Code,
		BuyNums:     t.Stock.BuyNums,
		BuyMoney:    t.Stock.BuyMoney,
	}

	data, err := json.Marshal(save)
	if err != nil {
		panic("json marshal error " + err.Error())
	}

	err = ioutil.WriteFile(savePath, data, 0660)
	if err != nil {
		panic("write file error " + err.Error())
	}

	return nil
}

func LoadSaveFileRecover() (t *TrainStock) {

	s := &SaveLocal{}
	data, err := ioutil.ReadFile(savePath)
	if err != nil {
		panic("read file error " + err.Error())
	}

	err = json.Unmarshal(data, s)
	if err != nil {
		panic("json unmarshal error " + err.Error())
	}
	st := &common.StockCommon{
		Code:     s.Code,
		BuyNums:  s.BuyNums,
		BuyMoney: s.BuyMoney,
	}

	t = &TrainStock{
		OriginMoney: s.OriginMoney,
		RemainMoney: s.RemainMoney,
		WinMoney:    s.WinMoney,
		WinPercent:  s.WinPercent,
		BuyDay:      s.BuyDay,
		ChangDay:    s.ChangDay,
		Status:      s.Status,
		Stock:       st,
	}

	info := &datastruct.StockBasicInfo{
		Symbol: t.Stock.Code,
	}
	nowDay := time.Now().Format("20060102")
	t.StockData = stockdata.GetDailTradeFromCSVSlice(info, t.BuyDay, nowDay)
	sort.Sort(sort.Reverse(t.StockData))
	t.len = len(t.StockData)
	t.Stock.BuyPrice = t.StockData[t.len-1].Close

	fmt.Println("recover success status =", t.Status)
	return t
}

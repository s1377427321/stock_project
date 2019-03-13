package common

import (
	"time"
	"common"
	"sort"
	"fmt"
	"stock_statistics/model"
	"strings"
	"net/http"
	"io/ioutil"
	"github.com/axgle/mahonia"
	"encoding/csv"
	"io"
	"strconv"
)

var CompareUpdateTicker = 15 * time.Second

type CompareCallBack func(code string, heightPrice, lowPrice float64, buyMoney float64)
type DeleteCompareCallBack func(code string) bool

type CompareStock struct {
	StockName    string
	StockUrl     string
	ForwardDay   int //前
	AfterwardDay int //后
	BeforeDay    string
	DayNum       int
	IsKeepStocks bool

	DayTrades model.DayTradeSlice

	closeCh chan interface{}

	AddNoticeFunc    CompareCallBack
	DeleteNoticeFunc DeleteCompareCallBack
}

func (this *CompareStock) Start() {
	this.BeforeDay = time.Unix(time.Now().Unix()-3600*int64(time.Now().Hour()+1), 0).Format("20060102")
	go func() {
		this.Update()
	}()
}

func (this *CompareStock) Update() {

	this.DoUpdate()
	for range time.NewTicker(CompareUpdateTicker).C {
		select {
		case <-this.closeCh:
			return
		default:
		}

		this.DoUpdate()
	}
}

func (this *CompareStock) DoUpdate() {

	h := time.Now().Hour()
	if h < 9 || h >= 15 {
		return
	}

	beforDay := time.Unix(time.Now().Unix()-3600*int64(time.Now().Hour()+1), 0).Format("20060102")

	if this.BeforeDay != beforDay {
		this.DayTrades = this.GetStockPriceData(this.StockName, int64(this.DayNum))
	} else {
		currentPrice, err := common.GetPriceFromUrl(this.StockUrl)
		if err != nil {
			return
		}

		this.DayTrades[0].Close = float32(currentPrice)
	}

	var heightPrice float64
	var lowPrice float64
	for i := 0; i < this.ForwardDay; i++ {
		heightPrice += common.Round64(float64(this.DayTrades[i].Close/float32(this.ForwardDay)), 3)
	}

	for i := 0; i < this.AfterwardDay; i++ {
		lowPrice += common.Round64(float64(this.DayTrades[i].Close/float32(this.AfterwardDay)), 3)
	}

	if this.IsKeepStocks{
		if heightPrice < lowPrice {

		}
	}else {
		if heightPrice > lowPrice {

		}
	}

	//this.AddNoticeFunc(this.StockName, heightPrice, lowPrice, 0)

	//currentPrice, err := common.GetPriceFromUrl(this.StockUrl)
	//if err != nil {
	//	return
	//}
	//
	//var lowPrice float64 = b.CurrentUpdateBuyLimit.Low
	//var heightPrice float64 = b.CurrentUpdateBuyLimit.Height
	//
	//if currentPrice < lowPrice {
	//	b.CurrentUpdateBuyLimit = b.CurrentUpdateBuyLimit.nextLow
	//} else if currentPrice > heightPrice {
	//	b.CurrentUpdateBuyLimit = b.CurrentUpdateBuyLimit.nextHeight
	//} else {
	//	return
	//}
	//
	//b.AddNoticeFunc(b.StockName, b.CurrentUpdateBuyLimit.Height, b.CurrentUpdateBuyLimit.Low, b.oneCopiesMoney)
}

func (this *CompareStock) GetStockPriceData(code string, dayNum int64) model.DayTradeSlice {
	dts := this.GetBeforData(code, dayNum)
	currentPrice, err := common.GetPriceFromUrl(this.StockUrl)
	if err != nil {
		return nil
	}
	tn := time.Now().Format("20060102")
	temp := &model.DayTrade{
		Date:  tn,
		Close: float32(currentPrice),
	}
	dts = append(dts, temp)
	sort.Sort(dts)
	return dts
}

//获取多天的数据
func (this *CompareStock) GetBeforData(code string, dayNum int64) model.DayTradeSlice {
	var url string
	//end := time.Now().Format("20060102")
	end := time.Unix(time.Now().Unix()-3600*int64(time.Now().Hour()+1), 0).Format("20060102")
	before := time.Unix(time.Now().Unix()-3600*24*dayNum, 0).Format("20060102")
	if strings.Contains(code, "sz") {
		s := strings.Split(code, "sz")
		url = fmt.Sprintf(this.StockUrl, "1"+s[1], before, end)
	} else if strings.Contains(code, "sh") {
		s := strings.Split(code, "sh")
		url = fmt.Sprintf(this.StockUrl, "0"+s[1], before, end)
	} else {
		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil
	}
	enc := mahonia.NewDecoder("gbk")
	_, utf8Body, _ := enc.Translate(body, true)

	days := make(model.DayTradeSlice, 0)

	r := csv.NewReader(strings.NewReader(string(utf8Body)))
	r.Read()
	for {
		cols, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil
		}

		Close, _ := strconv.ParseFloat(cols[3], 32)
		Open, _ := strconv.ParseFloat(cols[6], 32)
		Low, _ := strconv.ParseFloat(cols[5], 32)
		High, _ := strconv.ParseFloat(cols[4], 32)
		Volume, _ := strconv.Atoi(cols[11])
		Money, _ := strconv.ParseFloat(cols[12], 32)
		Code, _ := strconv.Atoi(strings.TrimLeft(cols[1], "'"))

		if Open == 0 {
			continue
		}

		//获取本地location
		timeLayout := "2006-01-02" //转化所需模板
		targetLayout := "20060102"
		loc, _ := time.LoadLocation("Local")                         //重要：获取时区
		theTime, _ := time.ParseInLocation(timeLayout, cols[0], loc) //使用模板在对应时区转化为time.time类型
		timeN := theTime.Format(targetLayout)                        //转化为时间戳 类型是int64

		days = append(days, &model.DayTrade{
			Date:   timeN,
			Code:   Code,
			Close:  float32(Close),
			High:   float32(High),
			Low:    float32(Low),
			Open:   float32(Open),
			Volume: Volume,
			Money:  int(Money),
		})

	}

	sort.Sort(days)

	var heightPrice float64
	var lowPrice float64
	for i := 0; i < this.ForwardDay; i++ {
		heightPrice += common.Round64(float64(this.DayTrades[i].Close/float32(this.ForwardDay)), 3)
	}

	for i := 0; i < this.AfterwardDay; i++ {
		lowPrice += common.Round64(float64(this.DayTrades[i].Close/float32(this.AfterwardDay)), 3)
	}

	if heightPrice > lowPrice {
		this.IsKeepStocks = true
	} else {
		this.IsKeepStocks = false
	}

	return days
}

func (this *CompareStock) Close() {
	this.DeleteNoticeFunc(this.StockName)
	close(this.closeCh)
}

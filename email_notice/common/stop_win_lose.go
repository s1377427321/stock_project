package common

import (
	"math"
	"github.com/astaxie/beego"
	"strings"
	"fmt"
	"encoding/json"
	"time"
)

var MyAllMoney = float64(160000)

var dataUrl = "http://hq.sinajs.cn/list=%s"

type StockBuyInfo struct {
	AllMoney   float64 `json:"all_money"`
	BuyDate    string  `json:"buy_date"`
	BuyPrice   float64 `json:"buy_price"`
	EndPrice   float64 `json:"end_price"`
	WinPercent float64 `json:"win_percent"`
	WinMoney   float64 `json:"win_money"`
}

type StopWinLose struct {
	Code            string          `json:"code"`
	Name            string          `json:"name"`
	BuyPrice        float64         `json:"buy_price"`
	OriginBuyPrice  float64         `json:"origin_buy_price"`
	BuyDate         string          `json:"buy_date"`
	LosePrice       float64         `json:"lose_price"`      //输到多少钱退出
	WinPrice        float64         `json:"win_price"`       //赢多少钱退出
	StopWinMoney    float64         `json:"stop_win_money"`  //止盈钱
	StopLoseMoney   float64         `json:"stop_lose_money"` //止损钱
	BuyMoney        float64         `json:"buy_money"`
	CurrentBuyNums  int             `json:"current_buy_nums"` //当前随持有的份数
	RealBuyNums     int             `json:"real_buy_nums"`    //第一次真实买入的份数
	Magnification   int             `json:"magnification"`    //倍率
	OperationRecord []*StockBuyInfo `json:"-"`
	Url             string          `json:"-"`
	Stock           *Stock          `json:"-"`
}

func (s *StopWinLose) Init(code, name string, price float64, magnification int) {

	/*
	根据策略重新计算出止盈止损的价格   止损 挂单卖 收盘前11分钟 直接平仓      止盈 挂单卖  收盘前 直接平仓
	基本策略
		1.总资金的百分之一的0.618止盈  总资金的百分之一的0.382止损
		2.根据当前投资环境，判断趋势，现在两种环境上升趋势 止损止盈计划x2   止损止盈计划x1
	 */
	s.OperationRecord = make([]*StockBuyInfo, 0)
	s.Code = code
	s.Name = name
	s.BuyPrice = price
	s.OriginBuyPrice = price
	s.Magnification = magnification
	s.BuyMoney = float64(MyAllMoney) / 10
	s.BuyDate = time.Now().Format("2006-01-02 15:04:05")

	s.RealBuyNums = int(math.Floor(s.BuyMoney/(s.BuyPrice*100)) * 100)
	s.CurrentBuyNums = s.RealBuyNums

	s.StopLoseMoney = MyAllMoney * 0.01 * 0.382 * float64(magnification)
	s.StopWinMoney = MyAllMoney * 0.01 * 0.618 * float64(magnification)

	s.LosePrice = (float64(s.RealBuyNums)*s.BuyPrice - s.StopLoseMoney) / float64(s.RealBuyNums)

	s.WinPrice = (float64(s.RealBuyNums)*s.BuyPrice + s.StopWinMoney) / float64(s.RealBuyNums)

	if strings.Index(code, "0") == 0 {
		s.Url = fmt.Sprintf(dataUrl, "sz"+code)
	} else if strings.Index(code, "6") == 0 || strings.Index(code, "5") == 0 {
		s.Url = fmt.Sprintf(dataUrl, "sh"+code)
	}

	s.Stock = &Stock{
		BuyMoney:       s.BuyMoney,
		Code:           s.Code,
		NoticeCallBack: nil,
		NoticeLimit:    0,
		Url:            s.Url,
		Name:           s.Name,
		LowPrice:       s.LosePrice,
		HightPrice:     s.WinPrice,
	}

	go s.Stock.Start()
	beego.Info(s)
}

func (s *StopWinLose) Close() {
	s.Stock.Close()
	s.Stock = nil
	s = nil
}

func (s *StopWinLose) UpdatePrice(high, low float64) {
	s.StopWinMoney = high
	s.StopLoseMoney = low
	s.Stock.UpdatePrice(s.StopWinMoney, s.StopLoseMoney)
}

func (s *StopWinLose) ShowInfo() string {
	b, err := json.Marshal(s)
	if err != nil {
		panic("*StopWinLose)ShowInfo  " + err.Error())
	}

	beego.Info(string(b))
	return string(b)
}

//高抛低吸 增加
func (s *StopWinLose) AddLowSuction(buyNums int, price float64) {
	allBuyNums := buyNums + s.CurrentBuyNums
	currentPrice := (float64(buyNums)*price + s.BuyPrice*float64(s.CurrentBuyNums)) / float64(allBuyNums)
	s.BuyPrice = currentPrice
	s.CurrentBuyNums = allBuyNums
	s.BuyMoney = s.BuyMoney + price*float64(buyNums)

	s.LosePrice = (float64(s.RealBuyNums)*s.BuyPrice - s.StopLoseMoney) / float64(s.RealBuyNums)

	s.WinPrice = (float64(s.RealBuyNums)*s.BuyPrice + s.StopWinMoney) / float64(s.RealBuyNums)

	s.Stock.UpdatePrice(s.WinPrice, s.LosePrice)
}

//高抛低吸 减少
func (s *StopWinLose) ReduceHighThrow(buyNums int, price float64) {
	allBuyNums := s.CurrentBuyNums - buyNums

	currentPrice := (s.BuyPrice*float64(s.CurrentBuyNums) - float64(buyNums)*price) / float64(allBuyNums)
	s.BuyPrice = currentPrice
	s.CurrentBuyNums = allBuyNums
	s.BuyMoney = s.BuyMoney - float64(buyNums)*price

	s.LosePrice = (float64(s.RealBuyNums)*s.BuyPrice - s.StopLoseMoney) / float64(s.RealBuyNums)

	s.WinPrice = (float64(s.RealBuyNums)*s.BuyPrice + s.StopWinMoney) / float64(s.RealBuyNums)

	s.Stock.UpdatePrice(s.WinPrice, s.LosePrice)
}

var _swl *StopWinLoseManage

func InstanceStopWinLoseManage() (*StopWinLoseManage) {
	if _swl == nil {
		_swl = &StopWinLoseManage{
			Items: make(map[string][]*StopWinLose, 0),
		}
	}
	return _swl
}

type StopWinLoseManage struct {
	Items map[string][]*StopWinLose
}

func (s *StopWinLoseManage) Add(code, name string, price float64, magnification int) {
	stockItem := &StopWinLose{}
	stockItem.Init(code, name, price, magnification)
	_, ok := s.Items[code]
	if !ok {
		s.Items[code] = make([]*StopWinLose, 0)
	}

	s.Items[code] = append(s.Items[code], stockItem)
}

func (s *StopWinLoseManage) DeleteStock(code string) {
	items, ok := s.Items[code]
	if !ok {
		beego.Error(fmt.Sprintf("HighThrow code not find %v", code))
		return
	}

	index := 0
	minPrice := float64(100000)
	for i := 0; i < len(items); i++ {
		item := items[i]
		if minPrice > item.BuyPrice {
			minPrice = item.BuyPrice
			index = i
		}
	}
	items[index].Close()
	items[index] = nil
	items = append(items[:index], items[index+1:]...)

	s.Items[code] = items
	if len(items) == 0 {
		delete(s.Items, code)
	}
}

func (s *StopWinLoseManage) ShowItems(code string) string {
	var data []byte
	var err error
	//var result string
	tempData:=""
	if code == "" {
		//data, err = json.Marshal(s.Items)
		//if err != nil {
		//	panic("ShowItems  Error 11 " + err.Error())
		//}
		for _,v:=range s.Items{

			d, err := json.Marshal(v)
			if err != nil {
				panic("ShowItems  Error 11 " + err.Error())
			}
			tempData = tempData + string(d)+"\n\n"
		}
	}else {
		items, _ := s.Items[code]
		data, err = json.Marshal(items)
		if err != nil {
			panic("ShowItems  Error 22 " + err.Error())
		}

		tempData = string(data)
	}


	return tempData

}

func (s *StopWinLoseManage) AddLowSuction(code string, buyNums int, price float64) {
	items, ok := s.Items[code]
	if !ok {
		beego.Error(fmt.Sprintf("HighThrow code not find %v", code))
		return
	}

	index := 0
	minPrice := float64(100000)
	for i := 0; i < len(items); i++ {
		item := items[i]
		if minPrice > item.OriginBuyPrice {
			minPrice = item.OriginBuyPrice
			index = i
		}
	}

	items[index].AddLowSuction(buyNums, price)
}

func (s *StopWinLoseManage) ReduceHighThrow(code string, buyNums int, price float64) {
	items, ok := s.Items[code]
	if !ok {
		beego.Error(fmt.Sprintf("HighThrow code not find %v", code))
		return
	}

	index := 0
	minPrice := float64(100000)
	for i := 0; i < len(items); i++ {
		item := items[i]
		if minPrice > item.OriginBuyPrice {
			minPrice = item.OriginBuyPrice
			index = i
		}
	}

	items[index].ReduceHighThrow(buyNums, price)
}

func (s *StopWinLoseManage) GetStockBuyNum(money, price float64) int {
	num := int(math.Floor(money/(price*100)) * 100)

	return num
}

/*
增加
http://127.0.0.1:5555/addswl?code=601628&name=中国人寿&price=26.030&magnification=1
http://swjswj.vip:5555/addswl?code=000001&name=测试&price=10.000&magnification=1
http://swjswj.vip:5555/addswl?code=000895&name=双汇发展&price=23.968&magnification=1
删除
http://127.0.0.1:5555/deleteswl?code=000001
http://swjswj.vip:5555/deleteswl?code=601628
显示
http://127.0.0.1:5555/showstocksswl?code=000001
http://swjswj.vip:5555/showstocksswl?code=000001
低吸
http://127.0.0.1:5555/addlowsuctionswl?code=000001&buy_num=1400&price=10.500
http://swjswj.vip:5555/addlowsuctionswl?code=000001&buy_num=1400&price=10.500
高抛
http://127.0.0.1:5555/reducehighthrowswl?code=000001&buy_num=1400&price=10.500
http://swjswj.vip:5555/reducehighthrowswl?code=000001&buy_num=1400&price=10.500
 */

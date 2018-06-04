package model

import "encoding/json"


//百分比策略操作数据结构
type StockTacticsOperate struct {
	Sp *StockStructStactics
	IsBeBuy bool
	IsBeSell bool
	Next *StockTacticsOperate
}

//百分比策略
type StockStructStactics struct {
	Id int
	Code int
	Percent float64
	Price float64
	BuyShare float64
	SellShare float64
	NeedMoney float64
}

func (self *StockStructStactics)ToString() string  {
	byteResult,err:= json.Marshal(self)
	if err !=nil{
		panic("convert json error")
	}

	return string(byteResult)
}

type DataChange struct {
	Data string       //时间
	TotalMoney float64  //所以的钱
	WinMoney float64 //赢了多少钱
	RemainMoney float64  //没有买股票的钱
	StockMoney float64  //持股的钱
	ShareHolding float64 //持股的份额
	StockPrice float64  //持股的价格
}

func (self *DataChange)ToString() string  {
	byteResult,err:= json.Marshal(self)
	if err !=nil{
		panic("convert json error")
	}

	return string(byteResult)
}

type DataChanges []*DataChange

func (p DataChanges) Len() int { return len(p) }

func (p DataChanges) Less(i, j int) bool {
	return p[i].StockPrice > p[j].StockPrice
}

func (p DataChanges) Swap(i, j int) { p[i], p[j] = p[j], p[i] }








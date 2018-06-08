package model

import (
	"encoding/json"
	"time"
	"sort"

)


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

	timeLayout := "2006-01-02"
	loc, _ := time.LoadLocation("Local")
	pi, _ := time.ParseInLocation(timeLayout, p[i].Data, loc)
	pj,_:=time.ParseInLocation(timeLayout,p[j].Data,loc)
	return pi.Unix() < pj.Unix()
}

func (p DataChanges) Swap(i, j int) { p[i], p[j] = p[j], p[i] }


type BeforeDayStruct struct {
	Code int         //代码
	StartDay string   //开始时间
	BeforeDay int     //向前多少天开始测试
	OriginMoney float64  //原始的钱
	TotalMoney float64  //所有的钱
	WinMoney float64   //赢的钱
	OneMoney float64    //十分之一的钱
	RemainMoney float64  //没有买股票的钱
	StockMoney float64  //持股的钱
	StockPrice float64  //持股的价格
	ShareHolding float64 //持股的份额
	//StockTactics map[int]*StockProcess  //股票里面的操作网格数据
	StockTacticsOperate map[int]*StockTacticsOperate  //股票里面的操作网格数据
	Tactics string //保存到数据库的策略 是 StockTacticsOperate的json
	TacticsWin float64    //执行策略挣的钱
	UpdateDay  string    //最新追踪时间
	DoTacticsNums int		//执行策略次数
	DataChanges DataChanges       //记录每次买入卖出
	TacticsChange string    //是 DataChanges的json
	Idx int              //有多少个多余十分之一的个数
}

func (self *BeforeDayStruct)ToStringTactics() string{
	var outString string=""
	for i:=self.Idx;i<len(self.StockTacticsOperate)+self.Idx;i++{
		if v,ok := self.StockTacticsOperate[i];ok && v!=nil{
			outString =outString + ","+ self.StockTacticsOperate[i].Sp.ToString()
		}else {
			panic("key value error")
		}
	}
	return outString
}


func (self *BeforeDayStruct)ToStringDataChanges() string  {
	sort.Sort(self.DataChanges)
	var outString string=""
	for _,v:=range self.DataChanges{
		outString = outString+","+v.ToString()
	}
	return  outString
}


func (self *BeforeDayStruct)AddADataChanges(data string,totalM,winM,remainM,staockM,shareH,stockP float64)  {
	temp:=&DataChange{
		Data:data,
		TotalMoney:totalM,
		WinMoney:winM,
		RemainMoney:remainM,
		StockMoney:staockM,
		ShareHolding:shareH,
		StockPrice:stockP,
	}
	self.DataChanges = append(self.DataChanges,temp)
}
/*
stocks := storage.GetAllStocks()
	storage.InsertResult(self.Code,self.StartDay,self.BeforeDay,self.TotalMoney,self.WinMoney,self.OneMoney,self.RemainMoney,self.StockMoney,self.StockPrice,
		self.ShareHolding,self.Tactics,self.DoTacticsNums,self.TacticsWin,self.TacticsChange,self.UpdateDay)
 */
//func GetAllStocks() []*Stock {
//	return storage.GetAllStocks()
//}
//
//func InsertResult(self *BeforeDayStruct)  {
//	storage.InsertResult(self.Code,self.StartDay,self.BeforeDay,self.TotalMoney,self.WinMoney,self.OneMoney,self.RemainMoney,self.StockMoney,self.StockPrice,
//		self.ShareHolding,self.Tactics,self.DoTacticsNums,self.TacticsWin,self.TacticsChange,self.UpdateDay)
//}










package percentTen

import (
	"time"
	"storage"
	."model"
	"math"
	"fmt"
	."commo"
	"sort"
)

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
	TacticsWin float64    //执行策略挣的钱
	UpdateDay  string    //最新追踪时间
	ChangeNums int		//执行策略次数
	DataChanges DataChanges       //记录每次买入卖出
	Idx int              //有多少个多余十分之一的个数
}

func (self *BeforeDayStruct)Do() {

	//before := time.Unix(time.Now().Unix()-int64(3600*24*self.BeforeDay), 0).Format("2006-01-02")

	data := storage.GetLatestSomeDataFromDay(self.Code,self.BeforeDay)
	record:=make(map[int]*DayTrade)
	var j int = 0
	for i:=0;i<self.BeforeDay ; {
		dayKey:=time.Unix(time.Now().Unix() - int64(3600*24*j),0).Format("2006-01-02")
		if v,ok:= data[dayKey]; v !=nil && ok {
			record[self.BeforeDay-i] = v

			i++
		}else {
			//panic("没有发现key"+dayKey)
		}
		j++
	}

	self.OneMoney =self.TotalMoney*0.1

	if v,ok:=record[1];ok && v!=nil{
		if self.StockPrice == 0 {
			self.StockPrice =Round64(float64(v.Close),3)
		}
		self.StartDay = v.Date
		self.ShareHolding=math.Floor(math.Floor(self.OneMoney/self.StockPrice)/100)*100

		self.StockMoney = math.Floor(self.ShareHolding*self.StockPrice)

		self.RemainMoney = math.Floor(self.TotalMoney - self.StockMoney)

		for i := self.Idx; i <= 9; i++ {
			temp01:=&StockTacticsOperate{
				Sp:&StockStructStactics{
					Code:self.Code,
					Id:i,
				},
				IsBeBuy:false,
				IsBeSell:false,
			}

			if i == 0 {
				//temp01.Percent =Round64(1+0.1,1)
				temp01.Sp.Percent =Round64(1,1)
				temp01.IsBeBuy = true
			//} else if i ==0  {
			//	temp01.Percent =Round64(1,1)
			}else {
				temp01.Sp.Percent =Round64(1-0.1*(math.Floor(float64(i))),1)
			}

			temp01.Sp.Price= Round64(self.StockPrice*temp01.Sp.Percent,3)
			//fmt.Println(temp01.Price)
			temp01.Sp.BuyShare=math.Floor(math.Floor(self.OneMoney/temp01.Sp.Price)/100)*100
			temp01.Sp.NeedMoney =math.Ceil(temp01.Sp.BuyShare*temp01.Sp.Price)


			self.StockTacticsOperate[i] = temp01

			if i != self.Idx {
				self.StockTacticsOperate[i-1].Sp.SellShare = temp01.Sp.BuyShare
				self.StockTacticsOperate[i-1].Next = temp01
			}else{
				self.StockTacticsOperate[i].Next = nil
			}
			//else {
			//	temp01.isBeBuy = true
			//}

		}

		self.AddADataChanges(self.StartDay,self.TotalMoney,self.WinMoney,self.RemainMoney,self.StockMoney,self.ShareHolding,self.StockPrice)

	}else {
		panic("没有发现key")
	}

	self.DealTransaction(record)



	//func InsertResult(code int,start_day ,befor_day string,total_money ,one_money,remain_money,stock_money,stock_price,share_holding float64,strProcess string)  {
	storage.InsertResult(self.Code,self.StartDay,self.BeforeDay,self.TotalMoney,self.WinMoney,self.OneMoney,self.RemainMoney,self.StockMoney,self.StockPrice,self.ShareHolding,self.ToStringTactics(),self.ChangeNums,self.TacticsWin,self.UpdateDay,self.ToStringDataChanges())

	//保存到数据库
	//fmt.Println(self.ToStringDataChanges())
	//fmt.Println(self.ToStringTactics())
	fmt.Println("")
}

func (self *BeforeDayStruct)DealTransaction(record map[int]*DayTrade)  {
	for i:=1;i<=self.BeforeDay ; i++ {

		//isEnd :=false
		isChange:=false
		changDay := ""
		if v,ok := record[i];ok && v!=nil{

			fmt.Println("------------------------------------")
			fmt.Println(v.Low,"  ",v.High,"  ",v.Close)
			fmt.Println("------------end------------------------")
			fmt.Println("")

			minPrice ,maxPrice :=self.GetMinMax(float64(v.Close),float64(v.High),float64(v.Low))
			changDay = v.Date
			for ii:=self.Idx;ii<len(self.StockTacticsOperate)+self.Idx;ii++ {
				if vv,ok := self.StockTacticsOperate[ii];ok && vv!=nil{
					nextVV:=vv.Next
					if minPrice < vv.Sp.Price && maxPrice > vv.Sp.Price {
						if nextVV !=nil && nextVV.IsBeBuy==true{           //卖出
							//if maxPrice> vv.Price {  //卖出  34  48
							self.StockPrice = vv.Sp.Price

							self.ShareHolding = self.ShareHolding - nextVV.Sp.BuyShare
							self.TacticsWin = self.TacticsWin+Round64(nextVV.Sp.BuyShare * (vv.Sp.Price-nextVV.Sp.Price),1)
							self.RemainMoney = self.RemainMoney + Round64(vv.Sp.Price* nextVV.Sp.BuyShare,1)
							nextVV.IsBeBuy = false
							//vv.isBeBuy = false
							//isEnd = true
							isChange = true
							break
							//}
						}

						if vv.IsBeBuy== false {
							//if minPrice < vv.Price && maxPrice > vv.Price{  //买入
							self.ShareHolding =self.ShareHolding + vv.Sp.BuyShare
							self.RemainMoney = self.RemainMoney-vv.Sp.NeedMoney
							vv.IsBeBuy = true
							//vv.isBeSell = false
							//isEnd = true
							isChange = true
							break
							//}
						}

					}
				}
			}
			self.StockPrice = float64(v.Close)
			self.StockMoney = Round64(self.ShareHolding*self.StockPrice,1)
			self.TotalMoney = self.RemainMoney + self.StockMoney
			self.WinMoney = self.TotalMoney - self.OriginMoney

			if isChange == true{
				self.ChangeNums = self.ChangeNums +1
				self.AddADataChanges(changDay,self.TotalMoney,self.WinMoney,self.RemainMoney,self.StockMoney,self.ShareHolding,self.StockPrice)
			}
			self.UpdateDay=v.Date
		}else {
			panic("key value error")
			self.UpdateDay = ""
		}

	}
}


func (self *BeforeDayStruct)GetMinMax(a,b,c float64) (aret float64 , bret float64) {
	temp:=[]float64{a,b,c}

	sort.Float64s(temp)

	return temp[0],temp[len(temp)-1]
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

func (self *BeforeDayStruct)ToStringDataChanges() string  {
	sort.Sort(self.DataChanges)
	var outString string=""
	for _,v:=range self.DataChanges{
		outString = outString+","+v.ToString()
	}
	return  outString
}


func  Start(code ,beforeDay int,menoy float64)  {
	beforeStruct:= NewBeforeDayStruct(code,beforeDay,menoy,-5)
	//beforeStruct.StockPrice = 4.3
	beforeStruct.Do()
}


func NewBeforeDayStruct(code ,beforeDayNum int,momey float64,idx int)  *BeforeDayStruct {
	return &BeforeDayStruct{
		Code:code,
		BeforeDay:beforeDayNum,
		OriginMoney:momey,
		TotalMoney:momey,
		StockTacticsOperate:make(map[int]*StockTacticsOperate),
		Idx:idx,
		ChangeNums:0,
		StockPrice:0,
		TacticsWin:0,
	}
}





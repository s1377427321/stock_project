package percentTen

import (
	"time"
	."model"
	"math"
	"fmt"
	."commo"
	"storage"
	"log"
)

func Do(self *BeforeDayStruct) {

	//before := time.Unix(time.Now().Unix()-int64(3600*24*self.BeforeDay), 0).Format("2006-01-02")

	data :=storage.GetLatestSomeDataFromDay(self.Code,self.BeforeDay) //GetLatestSomeDataFromDay(self.Code,self.BeforeDay) // storage.GetLatestSomeDataFromDay(self.Code,self.BeforeDay)
	record:=make(map[int]*DayTrade)
	var j int = 0
	log.Println("---------begin do for-----------")
	log.Println(len(data))

	if len(data) < self.BeforeDay {
		self.BeforeDay = len(data)
	}


	if  self.BeforeDay == 0{
		return
	}


	for i:=0;i<self.BeforeDay  ; {
		dayKey:=time.Unix(time.Now().Unix() - int64(3600*24*j),0).Format("2006-01-02")
		if v,ok:= data[dayKey]; v !=nil && ok {
			record[self.BeforeDay-i] = v

			i++
		}else {
			//panic("没有发现key"+dayKey)
		}
		j++
	}

	log.Println("---------end do for-----------")
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

	DealTransaction(self,record)



	//func InsertResult(code int,start_day ,befor_day string,total_money ,one_money,remain_money,stock_money,stock_price,share_holding float64,strProcess string)  {
	self.Tactics =self.ToStringTactics()
	self.TacticsChange = self.ToStringDataChanges()
	log.Println("-------- storage.InsertResult --------------------")
	log.Println(self.TacticsWin)
	storage.InsertResult(self.Code,self.StartDay,self.BeforeDay,self.TotalMoney,self.WinMoney,self.OneMoney,self.RemainMoney,self.StockMoney,self.StockPrice,
		self.ShareHolding,self.Tactics,self.DoTacticsNums,self.TacticsWin,self.TacticsChange,self.UpdateDay)

	log.Println("-------- storage.InsertResult  end--------------------")
	//storage.InsertResult(self)
	//保存到数据库
	//fmt.Println(self.ToStringDataChanges())
	//fmt.Println(self.ToStringTactics())
	fmt.Println("")
}

func DealTransaction(self *BeforeDayStruct,record map[int]*DayTrade)  {
	log.Println("-------- DealTransaction --------------------")
	for i:=1;i<=self.BeforeDay ; i++ {

		//isEnd :=false
		isChange:=false
		changDay := ""
		if v,ok := record[i];ok && v!=nil{

			//fmt.Println("------------------------------------")
			//fmt.Println(v.Low,"  ",v.High,"  ",v.Close)
			//fmt.Println("------------end------------------------")
			//fmt.Println("")

			minPrice ,maxPrice :=GetMinMax(float64(v.Close),float64(v.High),float64(v.Low))
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
				self.DoTacticsNums = self.DoTacticsNums +1
				self.AddADataChanges(changDay,self.TotalMoney,self.WinMoney,self.RemainMoney,self.StockMoney,self.ShareHolding,self.StockPrice)
			}
			self.UpdateDay=v.Date
		}else {
			panic("key value error")
			self.UpdateDay = ""
		}

	}
}

func  Start(code ,beforeDay int,menoy float64)  {
	beforeStruct:= NewBeforeDayStruct(code,beforeDay,menoy,-5)
	//beforeStruct.StockPrice = 4.3
	Do(beforeStruct)
}



func StartForAll(code,beforDay int ,menoy float64)  {
	//首先获取现有数据库里面有没有这条记录
	//如果有，则填充BeforeDayStruct数据，然后再更新到最新一天股价操作
	//如果没有，直接新建一个BeforeDayStruct数据 ，重新执行


}

func BeforeAllDay(beforeDay int,menoy float64)  {
	sem := make(chan int)
	stocks := storage.GetAllStocks()
	for _,v :=range stocks{
		log.Println(v)
		go doBeforeAllDay(v.Id,beforeDay,menoy,sem)
		//<-sem
		select {
		case <-time.After(5*time.Second):
			fmt.Println("chao shi")
		case <-sem:
		}
	}
}

func doBeforeAllDay(id, beforeDay int, menoy float64, sem chan int) {
	Start(id, beforeDay, menoy)
	fmt.Println("-----doBeforeAllDay--------")
	sem <- 1
}

func NewBeforeDayStruct(code ,beforeDayNum int,momey float64,idx int)  *BeforeDayStruct {
	return &BeforeDayStruct{
		Code:code,
		BeforeDay:beforeDayNum,
		OriginMoney:momey,
		TotalMoney:momey,
		StockTacticsOperate:make(map[int]*StockTacticsOperate),
		Idx:idx,
		DoTacticsNums:0,
		StockPrice:0,
		TacticsWin:0,
	}
}





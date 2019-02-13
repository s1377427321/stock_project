package common

import (
	"common"
	"time"
	"sort"
	"fmt"
)

var BuyUpdateTicker = 30 * time.Second

type AddNoticeCallBack func(code string, heightPrice, lowPrice float64,buyMoney float64)
type DeleteNoticeCallBack func(code string) bool

type BuyStock struct {
	StockName                  string
	StockUrl                   string
	AllMoney                   float64
	BuyPrice                   float64
	NumberOfCopies             int
	oneCopiesMoney       float64
	NumberOfCopiesPice         map[float64]*BuyLimit //价格、
	OrderNumberOfCopiesPiceKey []float64  //价格(排序好的)
	closeCh                    chan interface{}

	CurrentUpdateBuyLimit *BuyLimit

	AddNoticeFunc    AddNoticeCallBack
	DeleteNoticeFunc DeleteNoticeCallBack
}

type BuyLimit struct {
	Height float64
	Middle float64
	Low float64
	nextHeight *BuyLimit
	nextLow *BuyLimit
}

func (b *BuyStock) Start() {
	b.closeCh = make(chan interface{}, 0)
	b.oneCopiesMoney= common.Round64(b.AllMoney/IntToFloat64(b.NumberOfCopies), 3)
	oneCopiesStockPrice := common.Round64(b.BuyPrice/IntToFloat64(b.NumberOfCopies), 3)
	b.BuyPrice = oneCopiesStockPrice*IntToFloat64(b.NumberOfCopies)

	var i int = 1
	var max int = b.NumberOfCopies * 2
	for ; i <= b.NumberOfCopies; i++ {
		index := IntToFloat64(i)*oneCopiesStockPrice
		//copys := (int(b.oneCopiesMoney/index) / 100) * 100
		lt:=&BuyLimit{}
		lt.Middle = index
		b.NumberOfCopiesPice[index] = lt
		b.OrderNumberOfCopiesPiceKey = append(b.OrderNumberOfCopiesPiceKey, index)
	}

	for ; i <= max; i++ {
		index := IntToFloat64(i)*oneCopiesStockPrice
		//copys := (int(b.oneCopiesMoney/index) / 100) * 100
		lt:=&BuyLimit{}
		lt.Middle = index
		b.NumberOfCopiesPice[index] = lt
		b.OrderNumberOfCopiesPiceKey = append(b.OrderNumberOfCopiesPiceKey, index)
	}

	sort.Float64s(b.OrderNumberOfCopiesPiceKey)

	for j:=0;j<len(b.OrderNumberOfCopiesPiceKey);j++   {
		var beforB *BuyLimit
		if j == 0 {
			beforB = nil
		}else {
			beforB=b.NumberOfCopiesPice[b.OrderNumberOfCopiesPiceKey[j-1]]
		}

		var afterB *BuyLimit
		if j ==  len(b.OrderNumberOfCopiesPiceKey)-1{
			afterB = nil
		}else {
			afterB=b.NumberOfCopiesPice[b.OrderNumberOfCopiesPiceKey[j+1]]
		}

		current:=b.NumberOfCopiesPice[b.OrderNumberOfCopiesPiceKey[j]]
		if afterB != nil {
			current.Height = afterB.Middle
		}else {
			current.Height = current.Middle*2
		}

		if beforB != nil {
			current.Low = beforB.Middle
		}else {
			current.Low = -current.Middle
		}

		current.nextHeight = afterB
		current.nextLow = beforB
	}

	b.CurrentUpdateBuyLimit = b.NumberOfCopiesPice[b.BuyPrice]

	fmt.Println("AAA")

	go func() {
		b.Update()
	}()
}

func (b *BuyStock) Update() {

	//b.DoUpdate()
	b.AddNoticeFunc(b.StockName, b.CurrentUpdateBuyLimit.Height, b.CurrentUpdateBuyLimit.Low,b.oneCopiesMoney)
	for range time.NewTicker(BuyUpdateTicker).C {
		select {
		case <-b.closeCh:
			return
		default:
		}


		b.DoUpdate()
	}
}

func (b *BuyStock) DoUpdate() {

	currentPrice, err := GetPriceFromUrl(b.StockUrl)
	if err != nil {
		return
	}

	var lowPrice float64  = b.CurrentUpdateBuyLimit.Low
	var heightPrice float64 = b.CurrentUpdateBuyLimit.Height

	if currentPrice <lowPrice{
		b.CurrentUpdateBuyLimit = b.CurrentUpdateBuyLimit.nextLow
	}else if  currentPrice > heightPrice  {
		b.CurrentUpdateBuyLimit = b.CurrentUpdateBuyLimit.nextHeight
	}else {
		return
	}

	b.AddNoticeFunc(b.StockName, b.CurrentUpdateBuyLimit.Height, b.CurrentUpdateBuyLimit.Low,b.oneCopiesMoney)
}

func (b *BuyStock) Close() {
	b.DeleteNoticeFunc(b.StockName)
	close(b.closeCh)
}

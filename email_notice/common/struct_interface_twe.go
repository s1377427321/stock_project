package common

import (
	"common"
	"time"
	"sort"
	)

var BuyUpdateTicker = 60*time.Second

type AddNoticeCallBack func(code string, heightPrice, lowPrice float64)
type DeleteNoticeCallBack func(code string) bool

type BuyStock struct {
	StockName                  string
	StockUrl                   string
	AllMoney                   float64
	BuyPrice                   float64
	NumberOfCopies             int
	NumberOfCopiesPice         map[float64]int //价格、份数
	OrderNumberOfCopiesPiceKey []float64
	closeCh                    chan interface{}

	CurrentUpdateLowPrice float64

	AddNoticeFunc    AddNoticeCallBack
	DeleteNoticeFunc DeleteNoticeCallBack
}

func (b *BuyStock) Start() {
	b.closeCh = make(chan interface{},0)
	oneCopyMoneyPrice := common.Round64(b.AllMoney/float64(b.NumberOfCopies), 3)
	oneCopyStockPrice := common.Round64(b.BuyPrice/float64(b.NumberOfCopies), 3)

	var i int = 1
	var max int = b.NumberOfCopies*2
	for ; i <= b.NumberOfCopies; i++ {
		index := common.Round64( float64(i) * oneCopyStockPrice,3)
		copys := (int(oneCopyMoneyPrice/index) / 100) * 100
		b.NumberOfCopiesPice[index] = copys
		b.OrderNumberOfCopiesPiceKey = append(b.OrderNumberOfCopiesPiceKey, index)
	}

	for ; i <= max; i++ {
		index := common.Round64( float64(i) * oneCopyStockPrice,3)
		copys := (int(oneCopyMoneyPrice/index) / 100) * 100
		b.NumberOfCopiesPice[index] = copys
		b.OrderNumberOfCopiesPiceKey = append(b.OrderNumberOfCopiesPiceKey, index)
	}


	sort.Float64s(b.OrderNumberOfCopiesPiceKey)

	//fmt.Println(b.OrderNumberOfCopiesPiceKey)

	go func() {
		b.Update()
	}()
}

func (b *BuyStock) Update() {
	b.DoUpdate()
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
	currentPrice,err := GetPriceFromUrl(b.StockUrl)
	if err!= nil {
		return
	}

	var lowPrice float64 = 0
	var heightPrice float64

	for _, v := range b.OrderNumberOfCopiesPiceKey {
		if v <= currentPrice {
			lowPrice = v
		} else {
			heightPrice = v
			break
		}
	}

	if lowPrice != b.CurrentUpdateLowPrice {
		b.CurrentUpdateLowPrice = lowPrice
		b.AddNoticeFunc(b.StockName, heightPrice, lowPrice)
	}
}

func (b *BuyStock) Close() {
	b.DeleteNoticeFunc(b.StockName)
	close(b.closeCh)
}

package common

import (
	"strconv"
	"encoding/json"
	"common"
	"sort"
)

type Grid struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	StockNum  string `json:"stockNum"`
	NeedMoney string `json:"needMoney"`
	Percent   string `json:"percent"`
}

func NewStockGrid(code, name string, buyPrice, sellPrice float64, bl float64, hm float64, div float64) *StockGrid {
	buyPrice = common.Round64(buyPrice, 3)
	bl = common.Round64(bl, 3)
	hm = common.Round64(hm, 3)
	div = common.Round64(div, 3)
	tc := &StockGrid{
		BearLose: bl,
		Divide:   div,
	}
	tc.Code = code
	tc.AllMoney = hm
	tc.Name = name
	tc.BuyPrice = buyPrice
	tc.SellPrice = sellPrice

	gri := make(map[float64]string)
	tc.Grids = gri
	return tc
}

type StockGrid struct {
	StockCommon
	AllLoseMoney float64
	BearLose     float64
	Divide       float64
	Grids        map[float64]string
}

func (this *StockGrid) Do() {
	oneMoney := this.AllMoney / this.Divide
	onePercent := this.BearLose / this.Divide

	for i := 0; i <= int(this.Divide); i++ {
		//	fmt.Println(float64(i))
		needPercent := onePercent * (float64(i))
		divPrice := this.BuyPrice * (1 - needPercent)

		divPrice = common.Round64(divPrice, 3)

		divStockNum := (int(oneMoney/divPrice) / 100) * 100
		divNeedMoney := float64(divStockNum) * divPrice
		tempGrids := Grid{}
		tempGrids.Code = this.Code
		tempGrids.Price = strconv.FormatFloat(float64(divPrice), 'f', 3, 64)
		tempGrids.StockNum = strconv.Itoa(divStockNum)
		tempGrids.NeedMoney = strconv.FormatFloat(float64(divNeedMoney), 'f', 3, 64)
		tempGrids.Percent = strconv.FormatFloat(needPercent, 'f', 3, 64)
		tempGrids.Name = this.Name

		t, _ := json.Marshal(tempGrids)
		//index := strconv.FormatFloat(divPrice, 'f', 3, 64)
		this.Grids[divPrice] = string(t)

		this.AllLoseMoney += (divPrice - this.SellPrice) * float64(divStockNum)
	}

	for i := 1; i <= int(this.Divide); i++ {
		//	fmt.Println(float64(i))
		needPercent := onePercent * (float64(i))
		divPrice := this.BuyPrice * (1 + needPercent)

		divPrice = common.Round64(divPrice, 3)

		divStockNum := (int(oneMoney/divPrice) / 100) * 100
		divNeedMoney := float64(divStockNum) * divPrice
		tempGrids := Grid{}
		tempGrids.Code = this.Code
		tempGrids.Price = strconv.FormatFloat(float64(divPrice), 'f', 3, 64)
		tempGrids.StockNum = strconv.Itoa(divStockNum)
		tempGrids.NeedMoney = strconv.FormatFloat(float64(divNeedMoney), 'f', 3, 64)
		tempGrids.Percent = strconv.FormatFloat(needPercent, 'f', 3, 64)
		tempGrids.Name = this.Name
		t, _ := json.Marshal(tempGrids)
		//	this.Grids[divPrice] = string(t)
		//index := strconv.FormatFloat(divPrice, 'f', 3, 64)
		this.Grids[divPrice] = string(t)
	}

}

func (this *StockGrid) ToString() string {
	sortTmep := make([]float64, 0)

	for key, _ := range this.Grids {
		sortTmep = append(sortTmep, key)
	}

	sort.Float64s(sortTmep)

	var str string

	for _, val := range sortTmep {
		temp := this.Grids[val]
		//	tempStr := fmt.Sprintf("price =%s ,stockNum =%s, needMoney=%s \n ", temp.Price, temp.StockNum, temp.NeedMoney)
		str += temp + "\n"
	}

	str += strconv.FormatFloat(this.AllLoseMoney, 'f', 3, 64) + "\n"

	str += strconv.FormatFloat(this.SellPrice, 'f', 3, 64) + "\n"

	return str
}

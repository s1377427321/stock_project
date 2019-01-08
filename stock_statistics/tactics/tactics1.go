package tactics

import (
	"encoding/json"
	//	"sort"
	"strconv"
	."stock_statistics/commo"
)

/*

code  起始价格 能承受最大损失  总资金  分成多少档
		10          50%			6           5

买入价格  买入股数


*/

/*

http://120.79.154.53:4343/tactics1?code=600034&
origPrice=10.23&bearLose=0.5&haveMoney=50000&divide=5
*/



type Grid struct {
	Code      int `json:"code"`
	Price     string `json:"price"`
	StockNum  string `json:"stockNum"`
	NeedMoney string `json:"needMoney"`
	Percent   string `json:"percent"`
}

func NewTactics1(code int, op float64, bl float64, hm float64, div float64) *Tactics1 {
	op = Round64(op, 3)
	bl = Round64(bl, 3)
	hm = Round64(hm, 3)
	div = Round64(div, 3)
	tc := Tactics1{
		code:      code,
		origPrice: op,
		bearLose:  bl,
		haveMoney: hm,
		divide:    div,
	}
	gri := make(map[string]string)
	tc.Grids = gri
	return &tc
}

type Tactics1 struct {
	code      int
	origPrice float64
	bearLose  float64
	haveMoney float64
	divide    float64
	Grids     map[string]string
}

func (this *Tactics1) Do() {
	oneMoney := this.haveMoney / this.divide
	onePercent := this.bearLose / this.divide

	for i := 0; i <= int(this.divide); i++ {
		//	fmt.Println(float64(i))
		needPercent := onePercent * (float64(i))
		divPrice := this.origPrice * (1 - needPercent)

		divPrice = Round64(divPrice, 3)

		divStockNum := (int(oneMoney/divPrice) / 100) * 100
		divNeedMoney := float64(divStockNum) * divPrice
		tempGrids := Grid{}
		tempGrids.Code = this.code
		tempGrids.Price = strconv.FormatFloat(float64(divPrice), 'f', 3, 64)
		tempGrids.StockNum = strconv.Itoa(divStockNum)
		tempGrids.NeedMoney = strconv.FormatFloat(float64(divNeedMoney), 'f', 3, 64)
		tempGrids.Percent = strconv.FormatFloat(needPercent, 'f', 3, 64)

		t, _ := json.Marshal(tempGrids)
		index := strconv.FormatFloat(divPrice, 'f', 3, 64)
		this.Grids[index] = string(t)
	}

	for i := 1; i <= int(this.divide); i++ {
		//	fmt.Println(float64(i))
		needPercent := onePercent * (float64(i))
		divPrice := this.origPrice * (1 + needPercent)

		divPrice = Round64(divPrice, 3)

		divStockNum := (int(oneMoney/divPrice) / 100) * 100
		divNeedMoney := float64(divStockNum) * divPrice
		tempGrids := Grid{}
		tempGrids.Code = this.code
		tempGrids.Price = strconv.FormatFloat(float64(divPrice), 'f', 3, 64)
		tempGrids.StockNum = strconv.Itoa(divStockNum)
		tempGrids.NeedMoney = strconv.FormatFloat(float64(divNeedMoney), 'f', 3, 64)
		tempGrids.Percent = strconv.FormatFloat(needPercent, 'f', 3, 64)
		t, _ := json.Marshal(tempGrids)
		//	this.Grids[divPrice] = string(t)
		index := strconv.FormatFloat(divPrice, 'f', 3, 64)
		this.Grids[index] = string(t)
	}

}

func (this *Tactics1) ToString() string {
	/*	sortTmep := make([]string, 0)

		for key, _ := range this.Grids {
			sortTmep = append(sortTmep, key)
		}

		sort.Float64s(sortTmep)

		var str string

		for _, val := range sortTmep {
			temp := this.Grids[val]
			//	tempStr := fmt.Sprintf("price =%s ,stockNum =%s, needMoney=%s \n ", temp.Price, temp.StockNum, temp.NeedMoney)
			str += temp
		}
	*/
	return "str"
}

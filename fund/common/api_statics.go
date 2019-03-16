package common

import (
	. "fund/common/structs"
	"time"
	"github.com/astaxie/beego/logs"
	"strconv"
	"sort"
	"fmt"
	"fund/stocks"
)



//income  计算四个季度的动态市盈率
func GetShiYingLv() {
	var allIncome float64
	var allMarketValue float64

	now := time.Now().Format("20060102")
	nowInt, _ := strconv.Atoi(now)
	nowInt -= 1
	nowS := strconv.Itoa(nowInt)
	logs.Info(now)

	for _, v := range stocks.Hongli_etf {
		//if k>5 {
		//	break
		//}
		codeByte := []byte(v.Code)
		first := codeByte[0]
		if first == []byte("6")[0] {
			v.Code += ".SH"
		} else if first == []byte("0")[0] {
			v.Code += ".SZ"
		}

		perIncome := GetLatestYearPerIncome(v.Code)
		allIncome += perIncome * v.BuyNums

		currentV := GetDateClosePrice(v.Code, nowS)
		allMarketValue += v.BuyNums * currentV
	}

	logs.Info(fmt.Sprintf("%0.2f", allIncome))
	logs.Info(fmt.Sprintf("%0.2f", allMarketValue))
	logs.Info(fmt.Sprintf("%0.2f", allIncome/allMarketValue))
}

//获取最新四个季度的每股收入
func GetLatestYearPerIncome(code string) float64 {
	now := time.Now().Format("20060102")
	logs.Info(now)
	fields, items, _, _ := GetIncome(code, "20170101", now)
	ann_date_index := 0
	basic_eps := 0
	for k, v := range fields {
		if v == "ann_date" {
			ann_date_index = k

		}
		if v == "basic_eps" {
			basic_eps = k
		}
	}

	saveSlice := make(IncomeSS, 0)
	tempSave := make(map[int]int, 0)
	for _, v := range items {
		t, err := strconv.Atoi(v[ann_date_index])
		if err != nil {
			panic("time.Parse " + err.Error())
		}

		if _, ok := tempSave[t]; ok {
			continue
		}

		tempSave[t] = t
		newS := &IncomeS{
			Data: v,
			Time: t,
		}

		saveSlice = append(saveSlice, newS)
	}

	sort.Sort(saveSlice)

	for i := 0; i < len(saveSlice)-1; i++ {
		beforMoth := (saveSlice[i].Time / 100) % 100
		afterMoth := (saveSlice[i+1].Time / 100) % 100

		pc, err := strconv.ParseFloat(saveSlice[i].Data[basic_eps], 64)
		if err != nil {
			panic("ParseFloat  " + err.Error())
		}

		pc2, err := strconv.ParseFloat(saveSlice[i+1].Data[basic_eps], 64)
		if err != nil {
			panic("ParseFloat  " + err.Error())
		}

		if beforMoth >= 5 && afterMoth >= 4 {
			saveSlice[i].PerInCome = pc - pc2
		} else if beforMoth == 4 && afterMoth <= 4 {
			saveSlice[i].PerInCome = pc
		} else if beforMoth <= 4 && afterMoth >= 9 {
			saveSlice[i].PerInCome = pc - pc2
		} else {
			panic(fmt.Sprintf("beformaoth = %v ,afterMoth = %v", beforMoth, afterMoth))
		}

		logs.Info("+++ ", saveSlice[i].Data[0], saveSlice[i].Time, saveSlice[i].PerInCome)
	}

	var totolYearIncome float64 = 0
	for i := 0; i < 4; i++ {
		totolYearIncome += saveSlice[i].PerInCome
	}

	logs.Info("totolYearIncome =", totolYearIncome)
	logs.Info(fields)
	logs.Info(items)

	return totolYearIncome
}

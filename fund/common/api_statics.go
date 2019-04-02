package common

import (
	. "fund/common/structs"
	"time"
	"github.com/astaxie/beego/logs"
	"strconv"
	"sort"
	"fmt"
	"fund/stocks"
	commons "common"
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

//查找指定天数内，股价上涨不超过多少的股票 1
func FindHowMuchDayPriceUpOrDown() {

	howMuchRice := 0.20                 //n天内最低价格到当前价格的涨幅
	riceHowMuchDay := 60                //多少天之内的股价
	var yearIncome float64 = 1000000000 //年净收入以上的股票

	stocks, err := GetAllStocksInfo()
	if err != nil {
		panic(err)
	}

	FindResults := ""
	result := make(IncomeGoodSS, 0)
	for _, v := range stocks {
		res, ok := FindIncomeGood(v.TsCode, v.Name, howMuchRice, riceHowMuchDay, yearIncome)
		if ok {
			logs.Info(res)

			result = append(result, res)
			//FindResults = append(FindResults,res)
			//FindResults +="\n"+"\n"+ res + "\n"+"\n"
		}
	}

	sort.Sort(result)

	for i := 0; i < len(result); i++ {
		temp := result[i]
		FindResults += "\n" + "\n" + temp.Des + "\n" + "\n"
	}

	timeN := time.Now().Format("200601021543")
	commons.WriteWithIoutil(fmt.Sprintf("AAAA_%s.txt", timeN), FindResults)
}

//查找指定天数内，股价上涨不超过多少的股票 2
//查找指定代码的的Income数据
func QueryMySQLStockIncome(stock string) StockInComes {
	sql := fmt.Sprintf("SELECT ts_code,end_date,n_income_attr_p FROM stock_income where ts_code=\"%s\"", stock)
	res, err := Engine.QueryString(sql)
	if err != nil {
		panic("AAAAAAAAAAAAAAA" + err.Error())
	}

	sis := make([]*StockIncome, 0)
	for _, v := range res {
		si := &StockIncome{}
		commons.DataToStruct(v, si)
		sis = append(sis, si)
	}

	return sis
}

//查找指定天数内，股价上涨不超过多少的股票 3
//通过给定条件，查找对应符合条件的代码
func FindIncomeGood(tsCode, name string, howMuchRice float64, riceHowMuchDay int, yearIncome float64) (*IncomeGoodS, bool) {

	incomes := QueryMySQLStockIncome(tsCode)
	incomesLen := len(incomes)
	if  incomesLen/4 < 4 {
		return nil, false
	}

	sort.Sort(incomes)
	//logs.Info(incomes)

	var yearWin float64 = 0

	yearWins := make([]float64, 0)
	yearIndex := 0
	addCount := 0
	for i := 0; i < len(incomes)-1; i++ {

		//if addCount > 3 {
		//	break
		//}

		beforeTime, _ := strconv.Atoi(incomes[i].End_date)
		afterTime, _ := strconv.Atoi(incomes[i+1].End_date)

		beforMoth := (beforeTime / 100) % 100
		afterMoth := (afterTime / 100) % 100

		pc := incomes[i].N_income_attr_p

		pc2 := incomes[i+1].N_income_attr_p

		if beforMoth == 3 {
			addCount++
			yearWin += pc
		} else if beforMoth == 6 || beforMoth == 9 || beforMoth == 12 {
			//yearWin += pc - pc2
			addCount++
			if pc2 > 0 {
				yearWin += pc - pc2
			} else {
				yearWin += pc + pc2
			}

		} else {
			logs.Info(fmt.Sprintf("beformaoth = %v ,afterMoth = %v", beforMoth, afterMoth))
			continue
		}

		okNum:=(i+1)
		if okNum/4 > yearIndex {
			temp := yearWin
			yearWins = append(yearWins, temp)
			yearWin = 0
			yearIndex +=1
		}
	}

	isSelect := true
	for i:=0;i<4;i++ {
		if yearWins[i] < yearIncome {
			isSelect = false
			break
		}
	}


	if incomes[0].Ts_code == "000895.SZ" {
		fmt.Println()
	}

	if isSelect == false {
		logs.Error(fmt.Sprintf("%s  Not win money ok %f", incomes[0].Ts_code, yearWin))
		return nil, false
	}

	logs.Info(fmt.Sprintf("yearWin ok %f", yearWin))

	//if yearWin < yearIncome {
	//	logs.Error(fmt.Sprintf("%s  Not win money ok %f", incomes[0].Ts_code, yearWin))
	//	return nil, false
	//} else {
	//	logs.Info(fmt.Sprintf("yearWin ok %f", yearWin))
	//}

	dsi := GetLatestDayStock(commons.TsCodeToCode(incomes[0].Ts_code), riceHowMuchDay)

	if len(dsi) == 0 {
		logs.Error("Not Find GetLatestDayStock", incomes[0].Ts_code)
		return nil, false
	}

	//找到三十天里面的最大值，最小值，还有从三十天开始的到现在，价格变化最小的值
	var max float64
	var min float64 = dsi[len(dsi)-1].Close
	var follow float64 //记录价格变化值

	var before float64 = 0

	for i := len(dsi) - 1; i >= 0; i-- {
		v := dsi[i]

		if v.Close > max {
			max = v.Close
		}

		if v.Close < min {
			min = v.Close
		}

		if before == 0 {
			before = v.Close
		} else {
			follow = v.Close - before
		}

		if follow > 100 {

		}
	}

	if riceHowMuchDay > len(dsi) {
		return nil, false
	}

	//zuichujiage := dsi[riceHowMuchDay-1].Close
	zuihoujiage := dsi[0].Close

	//shangzhangxiajialv := (zuihoujiage - zuichujiage) / zuichujiage

	zuidazhangfulv := (max - min) / min

	zuididaoxianzaishangzhanglv := (zuihoujiage - min) / min

	huitiaolv := (max - zuihoujiage) / max

	des := fmt.Sprintf("股票%s  %s 在最近的%d天内，最低价格%f，最高价格%f，当前价格%f,其中最大上涨幅度为：%f,最低价格到当前价格上涨率为：%f,最高价格到当前价格回调率：%f",
		incomes[0].Ts_code, name, riceHowMuchDay, min, max, zuihoujiage, zuidazhangfulv, zuididaoxianzaishangzhanglv, huitiaolv)

	tempRes := &IncomeGoodS{
		Des:   des,
		Order: zuididaoxianzaishangzhanglv,
	}

	return tempRes, true
	//if howMuchRice > zuididaoxianzaishangzhanglv {
	//	return des, true
	//}
	//return "", false
}

func FindConceptIncomeGood() {
	concepts := []string{"水泥", "医药", "酒业", "环保", "钢铁", "银行", "建筑", "装饰"}
	sts := FindConceptCode(concepts)

	FindResults := ""
	howMuchRice := 0.20                 //n天内最低价格到当前价格的涨幅
	riceHowMuchDay := 80                //多少天之内的股价
	var yearIncome float64 = 1000000000 //年净收入以上的股票

	for _, v := range sts {
		scd := GetConceptDetailInfosByIdFromMySQL(v.Code)

		for _, v := range scd {
			res, ok := FindIncomeGood(v.TsCode, v.Name, howMuchRice, riceHowMuchDay, yearIncome)
			if ok {
				logs.Info(res)

				//FindResults = append(FindResults,res)
				FindResults += res.Des + "\n"
			}
		}

	}

	commons.WriteWithIoutil("BBB.txt", FindResults)

}

package main

import (
	_ "github.com/go-sql-driver/mysql"
		"fund/common/good_company"
	"time"
)

func main() {
	gm := &good_company.GoodCompanyManager{
		Companys: make(map[string]*good_company.GoodCompany, 0),
	}

	gm.Start()

	for range time.NewTicker(time.Hour*10).C{
	}
}

//
//func main() {
//
//	//common.SaveConceptDetailInfosToMySQL()
//	//reres:=common.FindConceptCode([]string{"水泥","医药","酒业","环保","钢铁","银行","建筑","装饰"})
//	//logs.Info(reres)
//	//common.SaveConceptInfosToMySQL()
//	//common.GetAllStocksInfo()
//	//common.UpdateDayTradeData()
//	//common.GetDealBuyStaticsInfos()
//	//common.GetDealBuyStaticsInfos()
//	//common.SaveIncomeDatasToMySQL()
//
//
//
//	//howMuchRice := 0.20
//	//riceHowMuchDay := 80
//	//var yearIncome float64 = 1000000000
//	//
//	//common.FindIncomeGood("601988.SH", howMuchRice, riceHowMuchDay, yearIncome)
//
//	//common.FindConceptIncomeGood()
//
//
//
//	common.SaveFinaIndicatorFromTuShare()
//
//
//
//	//common.FindHowMuchDayPriceUpOrDown()
//
//	//更新每日
//	//common.UpdateDayTradeData()
//
//	//HuiCheDaYu20YiShang()
//}

//func BuyTest() {
//	Monney := 100000
//	dayNum := 60
//	stocks, err := common.GetAllStocksInfo()
//	if err != nil {
//		panic(err)
//	}
//
//
//}

//
//func HuiCheDaYu20YiShang() {
//	stocks, err := common.GetAllStocksInfo()
//	if err != nil {
//		panic(err)
//	}
//	dayNum := 60
//
//	resultS := time.Now().Format("2006-01-02 15:04:05") + "\n"
//
//	for _, v := range stocks {
//
//		incomes := common.QueryMySQLStockIncome(v.TsCode)
//		if len(incomes) == 0 {
//			return
//		}
//
//		sort.Sort(incomes)
//		//logs.Info(incomes)
//
//		var yearWin float64 = 0
//
//		addCount := 0
//		for i := 0; i < len(incomes)-1; i++ {
//			if addCount > 3 {
//				break
//			}
//
//			beforeTime, _ := strconv.Atoi(incomes[i].End_date)
//			afterTime, _ := strconv.Atoi(incomes[i+1].End_date)
//
//			beforMoth := (beforeTime / 100) % 100
//			afterMoth := (afterTime / 100) % 100
//
//			pc := incomes[i].N_income_attr_p
//
//			pc2 := incomes[i+1].N_income_attr_p
//
//			if beforMoth == 3 {
//				addCount++
//				yearWin += pc
//			} else if beforMoth == 6 || beforMoth == 9 || beforMoth == 12 {
//				addCount++
//				if pc2 > 0 {
//					yearWin += pc - pc2
//				} else {
//					yearWin += pc + pc2
//				}
//
//			} else {
//				logs.Info(fmt.Sprintf("beformaoth = %v ,afterMoth = %v", beforMoth, afterMoth))
//				continue
//			}
//
//		}
//
//		fmt.Println(fmt.Sprintf("%s  %s year win :%f", v.TsCode, v.Name, yearWin))
//		if yearWin < 1000000000 {
//			continue
//		}
//
//		//收盘价的
//		dsi := common.GetLatestDayStock(v.Symbol, dayNum)
//
//		var max float64 = 0
//		//find max price
//		for i := len(dsi) - 1; i > 0; i-- {
//			vv := dsi[i]
//			price := vv.Close
//			if price > max {
//				max = price
//			}
//		}
//
//		currPrice := dsi[0].Close
//
//		lv := (max - currPrice) / max
//
//		logs.Info(fmt.Sprintf("%s huiCheLV:%f \n", v.TsCode, lv))
//
//		if lv >= 0.15 {
//			resultS += fmt.Sprintf("%s %s maxPrice:%f , currentPrice:%f, huiCheLV:%f \n", v.TsCode, v.Name, max, currPrice, lv)
//		}
//	}
//
//	resultS += "\n" + time.Now().Format("2006-01-02 15:04:05") + "\n"
//	commons.WriteWithIoutil("HuiCheDaYu20YiShang3.txt", resultS)
//
//}

//
func YeJiYouLiangYeJiZengZhang() {
	//incomes := common.QueryMySQLStockIncome(tsCode)
	//if len(incomes) == 0 {
	//	return "", false
	//}
	//
	//sort.Sort(incomes)
	////logs.Info(incomes)
	//
	//var yearWin float64 = 0
	//
	//addCount := 0
	//for i := 0; i < len(incomes)-1; i++ {
	//	if addCount > 3 {
	//		break
	//	}
	//
	//	beforeTime, _ := strconv.Atoi(incomes[i].End_date)
	//	afterTime, _ := strconv.Atoi(incomes[i+1].End_date)
	//
	//	beforMoth := (beforeTime / 100) % 100
	//	afterMoth := (afterTime / 100) % 100
	//
	//	pc := incomes[i].N_income_attr_p
	//
	//	pc2 := incomes[i+1].N_income_attr_p
	//
	//	if beforMoth == 3 {
	//		addCount++
	//		yearWin += pc
	//	}else if beforMoth == 6 || beforMoth == 9 || beforMoth == 12 {
	//		yearWin += pc - pc2
	//	}else {
	//		logs.Info(fmt.Sprintf("beformaoth = %v ,afterMoth = %v", beforMoth, afterMoth))
	//		continue
	//	}
	//
	//}
	//
	//if yearWin < yearIncome {
	//	logs.Error(fmt.Sprintf("%s  Not win money ok %f", incomes[0].Ts_code, yearWin))
	//	return "", false
	//} else {
	//	logs.Info(fmt.Sprintf("yearWin ok %f", yearWin))
	//}
	//
	//dsi :=common.GetLatestDayStock(commons.TsCodeToCode(incomes[0].Ts_code), riceHowMuchDay)
	//
	//if len(dsi) == 0 {
	//	logs.Error("Not Find GetLatestDayStock", incomes[0].Ts_code)
	//	return "", false
	//}
	//
	////找到三十天里面的最大值，最小值，还有从三十天开始的到现在，价格变化最小的值
	//var max float64
	//var min float64 = dsi[len(dsi)-1].Close
	//var follow float64 //记录价格变化值
	//
	//var before float64 = 0
	//
	//for i := len(dsi) - 1; i >= 0; i-- {
	//	v := dsi[i]
	//
	//	if v.Close > max {
	//		max = v.Close
	//	}
	//
	//	if v.Close < min {
	//		min = v.Close
	//	}
	//
	//	if before == 0 {
	//		before = v.Close
	//	} else {
	//		follow = v.Close - before
	//	}
	//
	//	if follow > 100 {
	//
	//	}
	//}
	//
	//if riceHowMuchDay > len(dsi) {
	//	return "", false
	//}
	//
	//zuichujiage := dsi[riceHowMuchDay-1].Close
	//zuihoujiage := dsi[0].Close
	//
	//shangzhangxiajialv := (zuihoujiage - zuichujiage) / zuichujiage
	//
	//zuidazhangfulv := (max - min) / zuichujiage
	//
	//zuididaoxianzaishangzhanglv := (zuihoujiage - min) / min
	//
	//des := fmt.Sprintf("股票%s  在最近的%d天内，价格从最初的%f，上涨或下降到%f,上涨率为：%f,其中最大上涨幅度为：%f,最低价格到当前价格上涨率为：%f", incomes[0].Ts_code, riceHowMuchDay, zuichujiage, zuihoujiage, shangzhangxiajialv, zuidazhangfulv, zuididaoxianzaishangzhanglv)
	//
	//if howMuchRice > zuididaoxianzaishangzhanglv {
	//	return des, true
	//}
	//return "", false
}

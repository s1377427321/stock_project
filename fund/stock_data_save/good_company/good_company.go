package good_company

import (
	"fund/common/structs"
	"strconv"
	"fmt"
	"github.com/astaxie/beego/logs"
	commons "common"
	. "fund/common"
	"time"
)

type GoodCompany struct {
	TsCode   string
	NameCode string
	AddDay   string

	Roe_dts             []float64 //5年平均 净资产收益率(扣除非经常损益)
	Netprofit_margins   []float64 //5 年平均销售净利率
	Grossprofit_margins []float64 //5年平均 销售毛利率
	N_income_attr_ps    []float64 //净利润(不含少数股东损益)
	TotalShare          float64   //总股本S

	YearNums int //多少年
	DayNums  int //多少天内

	CurrentPE                   float64 //当前pe
	ZuiDaZhangFuLv              float64 //最大上升比 (max - min) / min
	ZuiDiDaoXianZaiShangZhangLv float64 //最低到现在上涨率
	ZuiGaoDaoXianZaiHuiTiaoLv   float64 //最高到现在价格的回调率
	AVG5                        float64 //5日平均价格
	AVG10                       float64 //10日平均价格
	AVG20                       float64 //20日平均价格

	CurrentPrice float64 //当前价格
	HeighPrice   float64
	LowPrice     float64
}

func (g *GoodCompany) Init(tsCode, name string) {
	g.TsCode = tsCode
	g.NameCode = name
	g.AddDay = time.Now().Format("20060101")
	g.Roe_dts = make([]float64, g.YearNums)
	g.Netprofit_margins = make([]float64, g.YearNums)
	g.Grossprofit_margins = make([]float64, g.YearNums)
	g.N_income_attr_ps = make([]float64, g.YearNums)
}

func (g *GoodCompany) InitFinaIndicator(finaData []*structs.StockFinaIndicator) (bool) {
	if len(finaData) <= 4 {
		return false
	}

	first := finaData[0]

	firstYears, _ := strconv.Atoi(first.End_date)
	firstYears = (firstYears / 100) % 100

	//因为这两个不是累加的
	g.Netprofit_margins[0] = first.Netprofit_margin
	g.Grossprofit_margins[0] = first.Grossprofit_margin

	g.TotalShare = first.Total_share

	if firstYears == 12 {
		g.Roe_dts[0] = first.Roe_dt
	} else {

		var yearWin float64

		//最近的四次
		for i := 0; i < 4; i++ {
			//temp:= finaData[i]

			beforeTime, _ := strconv.Atoi(finaData[i].End_date)
			afterTime, _ := strconv.Atoi(finaData[i+1].End_date)

			beforMoth := (beforeTime / 100) % 100
			afterMoth := (afterTime / 100) % 100

			pc := finaData[i].Roe_dt
			pc2 := finaData[i+1].Roe_dt

			if beforMoth == 3 {
				yearWin += pc
			} else if beforMoth == 6 || beforMoth == 9 || beforMoth == 12 {
				//yearWin += pc - pc2
				if pc2 > 0 {
					yearWin += pc - pc2
				} else {
					yearWin += pc + pc2
				}

			} else {
				//logs.Info(fmt.Sprintf("beformaoth = %v ,afterMoth = %v", beforMoth, afterMoth))
				panic(fmt.Sprintf("beformaoth = %v ,afterMoth = %v", beforMoth, afterMoth))
				continue
			}
		}

		g.Roe_dts[0] = yearWin
	}

	//计算剩下的年份年报数据
	yearCount := 1
	for i := 1; i < len(finaData); i++ {
		temp := finaData[i]
		moth, _ := strconv.Atoi(temp.End_date)

		moth = (moth / 100) % 100

		if moth == 12 {
			g.Roe_dts[yearCount] = temp.Roe_dt
			g.Netprofit_margins[yearCount] = temp.Netprofit_margin
			g.Grossprofit_margins[yearCount] = temp.Grossprofit_margin
			yearCount++
		}

		if yearCount > 4 {
			break
		}
	}

	return true
}

func (g *GoodCompany) InitIncome(incomeData []*structs.StockIncome) (bool) {
	if len(incomeData) <= 4 {
		return false
	}

	first := incomeData[0]

	firstYears, _ := strconv.Atoi(first.End_date)
	firstYears = (firstYears / 100) % 100

	//if g.TsCode == "002736.SZ" {
	//	fmt.Println("")
	//}

	if firstYears == 12 {
		g.N_income_attr_ps[0] = first.N_income_attr_p
	} else {

		var yearWin float64

		//最近的四次
		for i := 0; i < 4; i++ {
			//temp:= finaData[i]

			beforeTime, _ := strconv.Atoi(incomeData[i].End_date)
			afterTime, _ := strconv.Atoi(incomeData[i+1].End_date)

			beforMoth := (beforeTime / 100) % 100
			afterMoth := (afterTime / 100) % 100

			pc := incomeData[i].N_income_attr_p
			pc2 := incomeData[i+1].N_income_attr_p

			if beforMoth == 3 {
				yearWin += pc
			} else if beforMoth == 6 || beforMoth == 9 || beforMoth == 12 {
				//yearWin += pc - pc2
				if pc2 > 0 {
					yearWin += pc - pc2
				} else {
					yearWin += pc + pc2
				}

			} else {
				//logs.Info(fmt.Sprintf("beformaoth = %v ,afterMoth = %v", beforMoth, afterMoth))
				panic(fmt.Sprintf("beformaoth = %v ,afterMoth = %v", beforMoth, afterMoth))
				continue
			}
		}

		g.N_income_attr_ps[0] = yearWin
	}

	//计算剩下的年份年报数据
	yearCount := 1
	for i := 1; i < len(incomeData); i++ {
		temp := incomeData[i]
		moth, _ := strconv.Atoi(temp.End_date)

		moth = (moth / 100) % 100

		if moth == 12 {
			g.N_income_attr_ps[yearCount] = temp.N_income_attr_p
			yearCount++
		}

		if yearCount > 4 {
			break
		}
	}

	return true
}

func (g *GoodCompany) InitOtherData() {

	dsi := GetLatestDayStock(commons.TsCodeToCode(g.TsCode), g.DayNums)

	if len(dsi) == 0 {
		logs.Error("Not Find GetLatestDayStock", g.TsCode)
		return
	}

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

	g.HeighPrice = max
	g.LowPrice = min

	g.CurrentPrice = dsi[0].Close

	g.ZuiDaZhangFuLv = (max - min) / min

	g.ZuiDiDaoXianZaiShangZhangLv = (g.CurrentPrice - min) / min

	g.ZuiGaoDaoXianZaiHuiTiaoLv = (max - g.CurrentPrice) / max

	g.CurrentPE = (g.TotalShare * g.CurrentPrice) / g.N_income_attr_ps[0]

	g.UpdateAVG()
}

func (g *GoodCompany) UpdateAVG() (bool) {

	dsi := GetLatestDayStock(commons.TsCodeToCode(g.TsCode), g.DayNums)

	if len(dsi) == 0 {
		logs.Error("Not Find GetLatestDayStock", g.TsCode)
		return false
	}

	index5 := 0
	var value5 float64 = 0
	index10 := 0
	var value10 float64 = 0
	index20 := 0
	var value20 float64 = 0
	for i := 0; i < len(dsi); i++ {
		v := dsi[i]
		if index5 < 5 {
			index5++
			value5 += v.Close
		}

		if index10 < 10 {
			index10++
			value10 += v.Close
		}

		if index20 < 20 {
			index20++
			value20 += v.Close
		} else {
			break
		}
	}

	g.AVG5 = value5 / 5
	g.AVG10 = value10 / 10
	g.AVG20 = value20 / 20
	return true
}

func (g *GoodCompany) IsGoodCompany() (bool) {

	//净利润
	var avgYearWin float64 = 0
	for _, v := range g.N_income_attr_ps {
		avgYearWin += v
	}
	avgYearWin = avgYearWin / float64(len(g.N_income_attr_ps))

	if avgYearWin < configGoodComany.YearIncome {
		return false
	}

	//roe
	var avgRoe float64 = 0
	for _, v := range g.Roe_dts {
		avgRoe += v
	}
	avgRoe = avgRoe / float64(len(g.Roe_dts))

	if avgRoe < configGoodComany.Roe {
		fmt.Println(fmt.Sprintf("avgRoe %f  configGoodComany %f",avgRoe,configGoodComany.Roe))
		return false
	}

	//净利率
	var avgNM float64 = 0
	for _, v := range g.Netprofit_margins {
		avgNM += v
	}
	avgNM = avgNM / float64(len(g.Netprofit_margins))

	if avgNM < configGoodComany.Netprofit_margins {
		fmt.Println(fmt.Sprintf("avgNM %f  configGoodComany.Netprofit_margins %f",avgNM,configGoodComany.Netprofit_margins))
		return false
	}

	////毛利率
	//var avgGM float64 = 0
	//for _, v := range g.Grossprofit_margins {
	//	avgGM += v
	//}
	//avgGM = avgGM / float64(len(g.Grossprofit_margins))
	//if avgGM != 0 {
	//	if avgGM < configGoodComany.Grossprofit_margins {
	//		fmt.Println(fmt.Sprintf("avgGM %f  configGoodComany.Grossprofit_margins %f",avgGM,configGoodComany.Grossprofit_margins))
	//		return false
	//	}
	//}

	//Roe_dts             []float64 //5年平均 净资产收益率(扣除非经常损益)
	//Netprofit_margins   []float64 //5 年平均销售净利率
	//Grossprofit_margins []float64 //5年平均 销售毛利率

	//	ZuiDaZhangFuLv              float64 //最大上升比 (max - min) / min
	//	ZuiDiDaoXianZaiShangZhangLv float64 //最低到现在上涨率
	//	ZuiGaoDaoXianZaiHuiTiaoLv   float64 //最高到现在价格的回调率
	//
	//ZuiDiDaoXianZaiShangZhangLvXiaoYu: 60,
	//	ZuiGaoDaoXianZaiHuiTiaoLvDaYu:     10,
	//		ZuiDaZhangFuLvXiaoYu:              90,



	//if g.ZuiGaoDaoXianZaiHuiTiaoLv < configGoodComany.ZuiGaoDaoXianZaiHuiTiaoLvDaYu {
	//	fmt.Println(fmt.Sprintf("g.ZuiGaoDaoXianZaiHuiTiaoLv %f  configGoodComany.ZuiGaoDaoXianZaiHuiTiaoLvDaYu %f",g.ZuiGaoDaoXianZaiHuiTiaoLv,configGoodComany.ZuiGaoDaoXianZaiHuiTiaoLvDaYu))
	//	return false
	//}
	//
	//if g.ZuiDaZhangFuLv > configGoodComany.ZuiDaZhangFuLvXiaoYu {
	//	fmt.Println(fmt.Sprintf("g.ZuiDaZhangFuLv %f  configGoodComany.ZuiDaZhangFuLvXiaoYu %f",g.ZuiDaZhangFuLv,configGoodComany.ZuiDaZhangFuLvXiaoYu))
	//	return false
	//}
	//
	//if g.ZuiDiDaoXianZaiShangZhangLv > configGoodComany.ZuiDiDaoXianZaiShangZhangLvXiaoYu {
	//	fmt.Println(fmt.Sprintf("g.ZuiDiDaoXianZaiShangZhangLv %f  configGoodComany.ZuiDiDaoXianZaiShangZhangLvXiaoYu %f",g.ZuiDiDaoXianZaiShangZhangLv,configGoodComany.ZuiDiDaoXianZaiShangZhangLvXiaoYu))
	//	return false
	//}

	return true
}

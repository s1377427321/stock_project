package good_company

import (
	"fund/common/structs"
	"strings"
	"time"
	"fmt"
	. "fund/common"
	commons "common"
	"reflect"
	"common/email"
)


var emailUrl = "1377427321@qq.com"
var emailServer = &email.EmailServers{
	ServerEmail:    "1377427321@qq.com",
	ServerPassword: "atpncirernxrhchj",
	ServerPort:     465,
	ServerIP:       "smtp.qq.com",
}

type DesOut struct {
	TsCode                      string
	NameCode                    string
	DayNums                     int
	CurrentPE                   float64 `json:"current_pe"`
	LowPrice                    float64 `json:"low_price"`
	HeighPrice                  float64 `json:"heigh_price"`
	CurrentPrice                float64 `json:"current_price"`
	ZuiDaZhangFuLv              float64 `json:"zui_da_zhang_fu_lv"`
	ZuiDiDaoXianZaiShangZhangLv float64 `json:"zui_di_dao_xian_zai_shang_zhang_lv"`
	ZuiGaoDaoXianZaiHuiTiaoLv   float64 `json:"zui_gao_dao_xian_zai_hui_tiao_lv"`
}

type GoodCompanyConfig struct {
	YearNums                          int
	DayNums                           int
	StartDay                          string
	EndDay                            string
	ZuiDiDaoXianZaiShangZhangLvXiaoYu float64
	ZuiGaoDaoXianZaiHuiTiaoLvDaYu     float64
	ZuiDaZhangFuLvXiaoYu              float64
	YearIncome                        float64
	Roe                               float64
	Grossprofit_margins               float64 //毛
	Netprofit_margins                 float64 //净
}

var configGoodComany GoodCompanyConfig = GoodCompanyConfig{
	StartDay:                          "20100101",
	EndDay:                            "Now",
	YearNums:                          5,
	DayNums:                           60,
	ZuiDiDaoXianZaiShangZhangLvXiaoYu: 0.6,
	ZuiGaoDaoXianZaiHuiTiaoLvDaYu:     0.1,
	ZuiDaZhangFuLvXiaoYu:              0.9,
	YearIncome:                        500000000,
	Roe:                               10,
	Netprofit_margins:                 10,
	Grossprofit_margins:               30,
}

type GoodCompanyManager struct {
	Companys map[string]*GoodCompany

	SendEmailDay string
}

func (g *GoodCompanyManager) Start() {
	go g.loop()
}

func (g *GoodCompanyManager) loop() {
	//stocks := g.GetGoodCompanys2MySQL()
	//g.StartCalculate(stocks)
	//g.SaveOneDayData()

	//更新所有的股票当天价格,返回手上有的股票
	stocks := UpdateDayTradeData()
	g.StartCalculate(stocks)
	g.SaveGoodCompanys2MySQL()

	for range time.NewTicker(time.Hour * 2).C {
		hour := time.Now().Hour()
		nowDay := time.Now().Format("20160101")
		if hour >= 20 && g.SendEmailDay != nowDay {
			g.SendEmailDay = nowDay

			//更新所有的股票当天价格,返回手上有的股票
			stocks := UpdateDayTradeData()
			g.StartCalculate(stocks)
			g.SaveGoodCompanys2MySQL()
		} else {
			stocks := g.GetGoodCompanys2MySQL()
			g.StartCalculate(stocks)
			g.SaveOneDayData()
		}
	}
}

//找到符合要求的公司
func (g *GoodCompanyManager) StartCalculate(stocks []*structs.StockBasicInfo) {
	//var err error
	//if len(g.Stocks) == 0 {
	//	g.Stocks, err = GetAllStocksInfo()
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	var now string = ""
	if strings.Contains(configGoodComany.EndDay, "now") {
		now = time.Now().Format("20060101")
	} else {
		now = configGoodComany.EndDay
	}

	for _, v := range stocks {
		//if len(g.Companys) > 6 {
		//	break
		//}

		finas := GetFinaIndicatorFromMySQL(v.TsCode, configGoodComany.StartDay, now)
		incomes := QueryMySQLStockIncome(v.TsCode)

		oldGC, ok := g.Companys[v.TsCode]
		if ok {
			oldGC.InitFinaIndicator(finas)
			oldGC.InitIncome(incomes)
			g.Companys[v.TsCode] = oldGC
		} else {
			newGC := &GoodCompany{
				YearNums: configGoodComany.YearNums,
				DayNums:  configGoodComany.DayNums,
			}
			newGC.Init(v.TsCode, v.Name)
			newGC.InitFinaIndicator(finas)
			newGC.InitIncome(incomes)
			g.Companys[v.TsCode] = newGC
		}

		oldGC, _ = g.Companys[v.TsCode]
		oldGC.InitOtherData()

		if oldGC.IsGoodCompany() == false {
			fmt.Println(fmt.Sprintf("                 delete %s!", v.TsCode))
			delete(g.Companys, v.TsCode)
		} else {
			fmt.Println("%s  is good company!", v.TsCode)
		}
	}
}

func (g *GoodCompanyManager) SaveGoodCompanys2MySQL() {
	sql := "replace into good_company(`ts_code`,`name`)VALUES"
	tempSql := ""
	mapLen := len(g.Companys)
	index := 0
	for _, v := range g.Companys {
		if mapLen == index+1 {
			tempSql = fmt.Sprintf("(\"%s\",\"%s\");", v.TsCode, v.NameCode)
		} else {
			tempSql = fmt.Sprintf("(\"%s\",\"%s\"),", v.TsCode, v.NameCode)
		}

		index++
		sql += tempSql
	}

	_, err := Engine.Exec(sql)
	if err != nil {
		panic(fmt.Sprintf("Engine.Exec %v Error %v", sql, err))
	}

}

func (g *GoodCompanyManager) GetGoodCompanys2MySQL() []*structs.StockBasicInfo {
	sql := "select * from good_company;"
	res, err := Engine.QueryString(sql)
	if err != nil {
		panic("Error select * from good_company;")
	}

	sbi := make([]*structs.StockBasicInfo, 0)
	for _, v := range res {
		si := &structs.StockBasicInfo{}
		commons.DataToStruct(v, si)
		sbi = append(sbi, si)
	}
	return sbi
}




//输出
func (g *GoodCompanyManager) SaveOneDayData() {

	desOuts := []interface{}{}
	for _, v := range g.Companys {
		temp := DesOut{
			TsCode:                      v.TsCode,
			NameCode:                    v.NameCode,
			DayNums:                     configGoodComany.DayNums,
			CurrentPE:                   v.CurrentPE,
			LowPrice:                    v.LowPrice,
			HeighPrice:                  v.HeighPrice,
			CurrentPrice:                v.CurrentPrice,
			ZuiDaZhangFuLv:              v.ZuiDaZhangFuLv,
			ZuiDiDaoXianZaiShangZhangLv: v.ZuiDiDaoXianZaiShangZhangLv,
			ZuiGaoDaoXianZaiHuiTiaoLv:   v.ZuiGaoDaoXianZaiHuiTiaoLv,
		}
		desOuts = append(desOuts, temp)
	}

	commons.SortBody(desOuts, func(p, q *interface{}) bool {
		v := reflect.ValueOf(*p)
		i := v.FieldByName("CurrentPE")
		v = reflect.ValueOf(*q)
		j := v.FieldByName("CurrentPE")
		return i.Float() < j.Float()
	})

	saveDes := ""
	for i := 0; i < len(desOuts); i++ {
		v := desOuts[i].(DesOut)
		des := fmt.Sprintf("股票%s  %s  PE=%f 在最近的%d天内，最低价格%f，最高价格%f，当前价格%f,其中最大上涨幅度为：%f,最低价格到当前价格上涨率为：%f,最高价格到当前价格回调率：%f",
			v.TsCode, v.NameCode, v.CurrentPE, configGoodComany.DayNums, v.LowPrice, v.HeighPrice, v.CurrentPrice, v.ZuiDaZhangFuLv, v.ZuiDiDaoXianZaiShangZhangLv, v.ZuiGaoDaoXianZaiHuiTiaoLv)

		saveDes += "\n \n" + des

	}

	saveDes += "\n \n  \n \n\n \n  \n \n ------------------------------------------------------------------"
	commons.SortBody(desOuts, func(p, q *interface{}) bool {
		v := reflect.ValueOf(*p)
		i := v.FieldByName("ZuiDiDaoXianZaiShangZhangLv")
		v = reflect.ValueOf(*q)
		j := v.FieldByName("ZuiDiDaoXianZaiShangZhangLv")
		return i.Float() < j.Float()
	})

	for i := 0; i < len(desOuts); i++ {
		v := desOuts[i].(DesOut)
		des := fmt.Sprintf("股票%s  %s  PE=%f 在最近的%d天内，最低价格%f，最高价格%f，当前价格%f,其中最大上涨幅度为：%f,最低价格到当前价格上涨率为：%f,最高价格到当前价格回调率：%f",
			v.TsCode, v.NameCode, v.CurrentPE, configGoodComany.DayNums, v.LowPrice, v.HeighPrice, v.CurrentPrice, v.ZuiDaZhangFuLv, v.ZuiDiDaoXianZaiShangZhangLv, v.ZuiGaoDaoXianZaiHuiTiaoLv)

		saveDes += "\n \n" + des

	}

	saveDes += "\n \n  \n \n\n \n  \n \n ------------------------------------------------------------------"
	commons.SortBody(desOuts, func(p, q *interface{}) bool {
		v := reflect.ValueOf(*p)
		i := v.FieldByName("ZuiGaoDaoXianZaiHuiTiaoLv")
		v = reflect.ValueOf(*q)
		j := v.FieldByName("ZuiGaoDaoXianZaiHuiTiaoLv")
		return i.Float() < j.Float()
	})

	for i := 0; i < len(desOuts); i++ {
		v := desOuts[i].(DesOut)
		des := fmt.Sprintf("股票%s  %s  PE=%f 在最近的%d天内，最低价格%f，最高价格%f，当前价格%f,其中最大上涨幅度为：%f,最低价格到当前价格上涨率为：%f,最高价格到当前价格回调率：%f",
			v.TsCode, v.NameCode, v.CurrentPE, configGoodComany.DayNums, v.LowPrice, v.HeighPrice, v.CurrentPrice, v.ZuiDaZhangFuLv, v.ZuiDiDaoXianZaiShangZhangLv, v.ZuiGaoDaoXianZaiHuiTiaoLv)

		saveDes += "\n \n" + des

	}



	timeN := time.Now().Format("20060102")
	fileName :=fmt.Sprintf("GOOD_COMPANY_%s.txt", timeN)
	commons.WriteWithIoutil(fileName, saveDes)

	ec := &email.EmailContent{
		NickName:     "文健",
		Subject:      fileName,
		BodyContent:  saveDes,
		NoticeEmails: []string{emailUrl},
	}

	email.SendEmailTo(emailServer, ec)
}


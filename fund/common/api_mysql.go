package common

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
	"io/ioutil"
	"time"
	"strconv"
	. "fund/common/structs"
		"github.com/go-xorm/xorm"
	"common"
		"stock_statistics/constant"
	"github.com/axgle/mahonia"
	"encoding/csv"
	"strings"
	"io"

	commons "common"
	)

var Engine *xorm.Engine

func init() {
	logs.Info("Start MYSQL")
	var err error
	Engine, err = xorm.NewEngine("mysql", SQLParams)
	if err != nil {
		panic(err)
	}
}


func SaveIncomeDatasToMySQL() {
	stocks, err := GetAllStocksInfo()
	if err != nil {
		panic(err)
	}
	now := time.Now().Format("20060102")
	startDate := "20100101"
	for _, v := range stocks {
		_, _, itemsMap, err := GetIncome(v.TsCode, startDate, now)
		if err != nil {
			logs.Error("Not SUCCESS ",v.TsCode)
			continue
		}
		stockItems := make([]*StockIncome, 0)
		for _, vv := range itemsMap {
			temp := &StockIncome{}
			commons.DataToStruct(vv, temp)

			stockItems = append(stockItems, temp)
		}

		SaveIncomeDatasToMySQL2(stockItems)
		time.Sleep(1000*time.Millisecond)
	}
}

func SaveIncomeDatasToMySQL2(stocks []*StockIncome) {
	sql := "replace into stock_income(`ts_code`,`ann_date`,`f_ann_date`,`end_date`,`report_type`,`comp_type`,`basic_eps`,`diluted_eps`,`total_revenue`,`revenue`,`int_income`,`prem_earned`,`comm_income`,`n_commis_income`,`n_oth_income`,`n_oth_b_income`,`prem_income`,`out_prem`,`une_prem_reser`,`reins_income`,`n_sec_tb_income`,`n_sec_uw_income`,`n_asset_mg_income`,`oth_b_income`,`fv_value_chg_gain`,`invest_income`,`ass_invest_income`,`forex_gain`,`total_cogs`,`oper_cost`,`int_exp`,`comm_exp`,`biz_tax_surchg`,`sell_exp`,`admin_exp`,`fin_exp`,`assets_impair_loss`,`prem_refund`,`compens_payout`,`reser_insur_liab`,`div_payt`,`reins_exp`,`oper_exp`,`compens_payout_refu`,`insur_reser_refu`,`reins_cost_refund`,`other_bus_cost`,`operate_profit`,`non_oper_income`,`non_oper_exp`,`nca_disploss`,`total_profit`,`income_tax`,`n_income`,`n_income_attr_p`,`minority_gain`,`oth_compr_income`,`t_compr_income`,`compr_inc_attr_p`,`compr_inc_attr_m_s`,`ebit`,`ebitda`,`insurance_exp`,`undist_profit`,`distable_profit`)VALUES"

	for i := 0; i < len(stocks); i++ {
		v := stocks[i]

		if len(stocks)-1 == i {
			sql += fmt.Sprintf("(\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f);", v.Ts_code, v.Ann_date, v.F_ann_date, v.End_date, v.Report_type, v.Comp_type, v.Basic_eps, v.Diluted_eps, v.Total_revenue, v.Revenue, v.Int_income, v.Prem_earned, v.Comm_income, v.N_commis_income, v.N_oth_income, v.N_oth_b_income, v.Prem_income, v.Out_prem, v.Une_prem_reser, v.Reins_income, v.N_sec_tb_income, v.N_sec_uw_income, v.N_asset_mg_income, v.Oth_b_income, v.Fv_value_chg_gain, v.Invest_income, v.Ass_invest_income, v.Forex_gain, v.Total_cogs, v.Oper_cost, v.Int_exp, v.Comm_exp, v.Biz_tax_surchg, v.Sell_exp, v.Admin_exp, v.Fin_exp, v.Assets_impair_loss, v.Prem_refund, v.Compens_payout, v.Reser_insur_liab, v.Div_payt, v.Reins_exp, v.Oper_exp, v.Compens_payout_refu, v.Insur_reser_refu, v.Reins_cost_refund, v.Other_bus_cost, v.Operate_profit, v.Non_oper_income, v.Non_oper_exp, v.Nca_disploss, v.Total_profit, v.Income_tax, v.N_income, v.N_income_attr_p, v.Minority_gain, v.Oth_compr_income, v.T_compr_income, v.Compr_inc_attr_p, v.Compr_inc_attr_m_s, v.Ebit, v.Ebitda, v.Insurance_exp, v.Undist_profit, v.Distable_profit)
		} else {
			sql += fmt.Sprintf("(\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f),", v.Ts_code, v.Ann_date, v.F_ann_date, v.End_date, v.Report_type, v.Comp_type, v.Basic_eps, v.Diluted_eps, v.Total_revenue, v.Revenue, v.Int_income, v.Prem_earned, v.Comm_income, v.N_commis_income, v.N_oth_income, v.N_oth_b_income, v.Prem_income, v.Out_prem, v.Une_prem_reser, v.Reins_income, v.N_sec_tb_income, v.N_sec_uw_income, v.N_asset_mg_income, v.Oth_b_income, v.Fv_value_chg_gain, v.Invest_income, v.Ass_invest_income, v.Forex_gain, v.Total_cogs, v.Oper_cost, v.Int_exp, v.Comm_exp, v.Biz_tax_surchg, v.Sell_exp, v.Admin_exp, v.Fin_exp, v.Assets_impair_loss, v.Prem_refund, v.Compens_payout, v.Reser_insur_liab, v.Div_payt, v.Reins_exp, v.Oper_exp, v.Compens_payout_refu, v.Insur_reser_refu, v.Reins_cost_refund, v.Other_bus_cost, v.Operate_profit, v.Non_oper_income, v.Non_oper_exp, v.Nca_disploss, v.Total_profit, v.Income_tax, v.N_income, v.N_income_attr_p, v.Minority_gain, v.Oth_compr_income, v.T_compr_income, v.Compr_inc_attr_p, v.Compr_inc_attr_m_s, v.Ebit, v.Ebitda, v.Insurance_exp, v.Undist_profit, v.Distable_profit)
		}
	}

	_, err := Engine.Exec(sql)
	if err != nil {
		panic(err)
	}
	if len(stocks)> 0 {
		logs.Info("SUCCESS",stocks[0].Ts_code)
	}

}

//插入所有股票信息到stock_basic数据库
func UpdateAllStocksAndInsertToSQL() {
	Engine, err := xorm.NewEngine("mysql", SQLParams)
	if err != nil {
		panic(err)
	}
	defer Engine.Close()

	fmt.Println("start  ", time.Now().Format("2006-01-02 15:04:05.000"))

	_, contect, _ := GetAllStocksCode()

	allStocks := make([]*StockBasicInfo, 0)
	sql := "replace into stock_basic(`ts_code`,`symbol`,`name`,`area`,`industry`,`fullname`,`enname`,`market`,`exchange`,`curr_type`,`list_status`,`list_date`,`delist_date`,`is_hs`) VALUES"
	insertSql := make([]string, 0)
	for _, v := range contect {
		if v[0] == "" {
			logs.Error("Error ")
			continue
		}

		sb := &StockBasicInfo{
			TsCode:     v[0],
			Symbol:     v[1],
			Name:       v[2],
			Area:       v[3],
			Industry:   v[4],
			Fullname:   v[5],
			Enname:     v[6],
			Market:     v[7],
			Exchange:   v[8],
			CurrType:   v[9],
			ListStatus: v[10],
			ListDate:   v[11],
			DelistDate: v[12],
			IsHs:       v[13],
		}
		allStocks = append(allStocks, sb)
		sqlChild := fmt.Sprintf("(\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\")", sb.TsCode, sb.Symbol, sb.Name, sb.Area, sb.Industry, sb.Fullname, sb.Enname, sb.Market, sb.Exchange, sb.CurrType, sb.ListStatus, sb.ListDate, sb.DelistDate, sb.IsHs)
		insertSql = append(insertSql, sqlChild)
		//fmt.Print(res.RowsAffected())
	}

	for i := 0; i < len(insertSql)-1; i++ {
		sql += insertSql[i] + ","
	}

	sql += insertSql[len(insertSql)-1] + ";"

	effct, err := Engine.Exec(sql)
	if err != nil {
		panic(err)
	}

	en, _ := effct.RowsAffected()

	fmt.Println("SQL RowsAffected", en)
}

//从mysql数据库中获取所有股票
func GetAllStocksInfo() ([]*StockBasicInfo, error) {
	resultS := make([]*StockBasicInfo, 0)

	//Engine, err := xorm.NewEngine("mysql", SQLParams)
	//if err != nil {
	//	panic(err)
	//}
	//defer Engine.Close()

	sqlS := fmt.Sprintf("SELECT * from stock_basic;")
	queryR, err := Engine.QueryString(sqlS)
	if err != nil {
		logs.Error(fmt.Sprintf("QueryString %s error :%s", sqlS, err.Error()))
		return nil, err
	}

	for _, v := range queryR {

		temp := &StockBasicInfo{}
		common.DataToStruct(v, temp)
		//logs.Info(temp)
		resultS = append(resultS, temp)
	}

	return resultS, nil
}

//trade_his  1  从网络获取个股股价等信息，插入表trade_his中
func UpdateDayTradeData() {
	stocks, err := GetAllStocksInfo()
	if err != nil {
		panic(err)
	}

	sem := make(chan int, 15)
	for _, stock := range stocks {
		go func(s *StockBasicInfo) {
			UpdateDayTradeDataNext1(s, sem)
		}(stock)
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("========do time.After(1) ======")
			//delete(sem,chan)
		case <-sem:
			fmt.Println("========do self <-sem======")
		}

		//UpdateDayTradeDataNext1(stock, sem)
	}
}




//trade_his  2 插入操作
func UpdateDayTradeDataNext1(stock *StockBasicInfo, sem chan int) {
	defer commons.RecoverPanic()
	lastestStocksInfos := GetLatestDayStock(stock.Symbol,1)
	lastestStocksInfo :=lastestStocksInfos[0]

	now := time.Now().Format("20060102")
	var added []*DialyStockInfo
	if lastestStocksInfo == nil || lastestStocksInfo.Date == "" {
		before := time.Unix(time.Now().Unix()-3600*24*360*5, 0).Format("20060102")
		added = GetDailTradeFromCSV(stock, before, now)
	} else {
		last, _ := time.Parse("2006-01-02", lastestStocksInfo.Date)
		// init today
		from := time.Unix(last.Unix()+3600*24*1, 0).Format("20060102")
		//if from == now {
		//	log.Println("not data need insert ", lastestStocksInfo.Code)
		//	sem <- 1
		//	return
		//}

		added = GetDailTradeFromCSV(stock, from, now)
	}

	logs.Info("InsertTradeHis  ", stock.TsCode)

	if len(added) != 0 {
		InsertTradeHis(added)
	}

	logs.Info("SUCESS ", stock.TsCode)

	sem <- 1
}

//trade_his  3 获取最近一天记录到数据库的数据，为了更新数据用
func GetLatestDayStock(codeStr string ,dayNums int) []*DialyStockInfo {
	sqlStr := `
	SELECT * FROM trade_his WHERE code=%d ORDER BY date DESC LIMIT %d;
	`
	code, err := strconv.Atoi(codeStr)
	if err != nil {
		panic(fmt.Sprintf("%v strconv.Atoi error %v", codeStr, err))
	}

	sql := fmt.Sprintf(sqlStr, code,dayNums)
	sqlResult, err := Engine.QueryString(sql)
	if err != nil || sqlResult == nil {
		return nil
	}

	ret := make([]*DialyStockInfo,0)
	for i:=0;i<len(sqlResult);i++ {
		temp:=&DialyStockInfo{}
		commons.DataToStruct(sqlResult[i], temp)
		ret = append(ret,temp)
	}

	return ret
}

//trade_his  4 从163网上下载csv格式数据，然后保存到数据库中
func GetDailTradeFromCSV(stock *StockBasicInfo, begin, end string) []*DialyStockInfo {
	typeStock := []byte(stock.Symbol)[0]
	var url string
	if typeStock == []byte("0")[0] || typeStock == []byte("3")[0] {
		url = fmt.Sprintf(constant.DAY_TRADE_API, "1"+stock.Symbol, begin, end)
	} else if typeStock == []byte("6")[0] {
		url = fmt.Sprintf(constant.DAY_TRADE_API, "0"+stock.Symbol, begin, end)
	} else {
		panic(fmt.Sprintf("%v  error code id %v", "GetDailTradeFromCSV", stock.Symbol, ))
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil
	}
	enc := mahonia.NewDecoder("gbk")
	_, utf8Body, _ := enc.Translate(body, true)

	days := make([]*DialyStockInfo, 0)

	r := csv.NewReader(strings.NewReader(string(utf8Body)))
	r.Read()
	for {
		cols, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil
		}

		Close, _ := strconv.ParseFloat(cols[3], 64)
		Open, _ := strconv.ParseFloat(cols[6], 64)
		Low, _ := strconv.ParseFloat(cols[5], 64)
		High, _ := strconv.ParseFloat(cols[4], 64)
		Volume, _ := strconv.Atoi(cols[11])
		Money, _ := strconv.ParseFloat(cols[12], 64)
		Code, _ := strconv.Atoi(strings.TrimLeft(cols[1], "'"))

		if Open == 0 {
			continue
		}

		days = append(days, &DialyStockInfo{
			Date:   cols[0],
			Code:   Code,
			Close:  Close,
			High:   High,
			Low:    Low,
			Open:   Open,
			Volume: Volume,
			Money:  int(Money),
		})

	}

	return days
}

//trade_his  5 把数据插入到数据库中
func InsertTradeHis(stockInfos []*DialyStockInfo) {
	sql := "replace into trade_his(`code`,`date`,`open`,`close`,`high`,`low`,`volume`,`money`)VALUES"

	for i := 0; i < len(stockInfos); i++ {
		v := stockInfos[i]
		var tempStr string
		if i == len(stockInfos)-1 {
			tempStr = fmt.Sprintf("(%d,\"%s\",%f,%f,%f,%f,%d,%d);", v.Code, v.Date, v.Open, v.Close, v.High, v.Low, v.Volume, v.Money)
		} else {
			tempStr = fmt.Sprintf("(%d,\"%s\",%f,%f,%f,%f,%d,%d),", v.Code, v.Date, v.Open, v.Close, v.High, v.Low, v.Volume, v.Money)
		}
		sql += tempStr
	}

	_, err := Engine.Exec(sql)

	if err != nil {
		logs.Error(sql)
		panic(fmt.Sprintf("%v ERROR %v", "Engine.Exec", err))
	}
}

//根据中文的名字去查找所属概率的codeID
//reres:=common.FindConceptCode([]string{"水泥","医药","酒业","环保","钢铁","银行","建筑","装饰"})
func FindConceptCode(findStr []string) []*StockConcept  {
	ret:=make([]*StockConcept , 0)

	for _,v:=range findStr{
		sqls:="SELECT * FROM stock_concept WHERE name LIKE"+"'%"+v+"%';"
		qresut,err:= Engine.QueryString(sqls)
		if err!=nil {
			logs.Error("++++++++++++++++++++",v)
			continue
		}
		for _,vv:=range qresut{
			temp:=&StockConcept{}
			commons.DataToStruct(vv,temp)

			ret = append(ret,temp)
		}

	}

	return ret
}

func GetConceptInfosFromMySQL() []*StockConcept  {
	ret:=make([]*StockConcept,0)

	sql:="SELECT * FROM stock_concept;"

	sqlRes,err:=Engine.QueryString(sql)
	if err != nil {
		panic(fmt.Sprintf("GetConceptInfosFromMySQL ERROR %v",err))
	}

	for _,v:=range sqlRes{
		temp:=&StockConcept{}
		commons.DataToStruct(v,temp)
		ret = append(ret,temp)
	}

	return ret
}

//通过给定的概念ID找到对应的股票
func GetConceptDetailInfosByIdFromMySQL(conceptID string) []*StockConceptDetail  {
	ret:=make([]*StockConceptDetail,0)

	sql:=fmt.Sprintf("SELECT * FROM stock_concept_detail WHERE concept_code=\"%s\" ;",conceptID)

	sqlRes,err:=Engine.QueryString(sql)
	if err != nil {
		panic(fmt.Sprintf("GetConceptInfosFromMySQL ERROR %v",err))
	}

	for _,v:=range sqlRes{
		temp:=&StockConceptDetail{}
		commons.DataToStruct(v,temp)
		ret = append(ret,temp)
	}

	return ret
}

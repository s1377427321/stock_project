package common

import (
	. "fund/common/structs"
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
	"errors"
	"reflect"
	"fmt"
	"strconv"
	commons "common"
	"github.com/astaxie/beego/logs"
	"time"
)

//获取收盘价格
func GetDateClosePrice(code string, date string) float64 {
	fields, content, err := GetDailyBasic(code, date)
	if err != nil {
		panic("GetDateClosePrice  " + err.Error())
	}

	var closePriceIndex int
	for i := 0; i < len(fields); i++ {
		if fields[i] == "close" {
			closePriceIndex = i
		}
	}

	v, err := strconv.ParseFloat(content[0][closePriceIndex], 64)

	return v
}

//获取TuShare的数据
func PostToUrl(apiName string, fields string, ps *Params, ) ([]string, [][]string, []map[string]string, error) {
	//ps := common.Params{
	//	Ts_Code:    "002594.SZ",
	//	Trade_Date: "20180726",
	//}

	//rp := common.ReqParams{
	//	ApiName: "daily_basic",
	//	Token:   common.TOKEN,
	//	Params:  *ps,
	//	Fields:  "ts_code,pe,pb,close,total_share,float_share,free_share,total_mv",
	//}

	rp := ReqParams{
		ApiName: apiName,
		Token:   TOKEN,
		Params:  *ps,
		Fields:  fields,
	}

	jsb, err := json.Marshal(rp)
	if err != nil {
		panic("json error " + err.Error())
	}

	logs.Info("----------", string(jsb))

	req := bytes.NewBuffer(jsb)

	body_type := "application/json;charset=utf-8"
	resp, err := http.Post(HTTP_URL, body_type, req)
	body, _ := ioutil.ReadAll(resp.Body)
	logs.Info("----------", string(body))

	var dat map[string]interface{}
	json.Unmarshal(body, &dat)
	if v, ok := dat["code"]; ok {
		if v.(float64) != 0 {
			return nil, nil, nil, errors.New("request error :" + err.Error())
		}
	}

	for _, v := range dat {
		if v != nil {
			//logs.Info(reflect.TypeOf(v).Kind())
			if reflect.TypeOf(v).Kind() == reflect.Map {

				f := GetFields(v.(map[string]interface{})["fields"].([]interface{}))
				items, itmsMap := GetContent(v.(map[string]interface{})["fields"].([]interface{}), v.(map[string]interface{})["items"].([]interface{}))
				//logs.Info(f)
				//logs.Info(items)
				return f, items, itmsMap, nil
			}
		}
	}
	return nil, nil, nil, errors.New("NOT FIND")
}

func GetFields(fields []interface{}) []string {
	res := make([]string, 0)
	for _, v := range fields {
		//logs.Info(k,v)
		res = append(res, v.(string))
	}

	return res
}

func GetContent(fields []interface{}, con []interface{}) ([][]string, []map[string]string) {
	res := make([][]string, 0)
	resMaps := make([]map[string]string, 0)
	for _, v := range con {
		items := make([]string, 0)
		for _, vv := range v.([]interface{}) {
			//logs.Info("+++++++++++++++++")
			//logs.Info(kk,vv,reflect.TypeOf(vv))
			if vv == nil {
				items = append(items, "null")
				continue
			}
			kind := reflect.TypeOf(vv).Kind()
			var value string
			switch kind {
			case reflect.Float64:
				convv := vv.(float64)
				value = fmt.Sprintf("%f", convv)
				//value = strconv.FormatFloat(convv,'E',-1,64)
				//float,_ := strconv.ParseFloat(value,64)
				//logs.Info("***** ",value)
			case reflect.String:
				value = vv.(string)
				if value == "" {
					panic("AAAAAAAAAAAAAA")
				}
				//logs.Info("***** ",value)
			case reflect.Int:
				convv := vv.(int)
				value = fmt.Sprintf("%d", convv)
				if value == "" {
					panic("AAAAAAAAAAAAAA")
				}
			default:
				panic("GetContent GetContent Error")
			}
			items = append(items, value)

		}

		resMap := make(map[string]string, 0)
		for i := 0; i < len(items); i++ {
			key := fields[i].(string)
			resMap[key] = items[i]
		}

		res = append(res, items)
		resMaps = append(resMaps, resMap)

		//fmt.Println(reflect.TypeOf(v))
		//logs.Info(k,v)
	}
	return res, resMaps
}

/*
ts_code	str	TS股票代码
trade_date	str	交易日期
close	float	当日收盘价
turnover_rate	float	换手率（%）
turnover_rate_f	float	换手率（自由流通股）
volume_ratio	float	量比
pe	float	市盈率（总市值/净利润）
pe_ttm	float	市盈率（TTM）
pb	float	市净率（总市值/净资产）
ps	float	市销率
ps_ttm	float	市销率（TTM）
total_share	float	总股本 （万）
float_share	float	流通股本 （万）
free_share	float	自由流通股本 （万）
total_mv	float	总市值 （万元）
circ_mv	float	流通市值（万元）
 */
func GetDailyBasic(code string, date string) ([]string, [][]string, error) {
	//filds := "ts_code,pe,pb,close,total_share,float_share,free_share,total_mv"
	filds := "ts_code,trade_date,close,turnover_rate,turnover_rate_f,volume_ratio,pe,pe_ttm,pb,ps,ps_ttm,total_share,float_share,free_share,total_mv,circ_mv"
	//ps := &Params{
	//	Ts_Code:    "002594.SZ",
	//	Trade_Date: "20180726",
	//}
	ps := &Params{
		Ts_Code:    code,
		Trade_Date: date,
	}
	fields, items, _, err := PostToUrl("daily_basic", filds, ps)

	if err != nil {
		return nil, nil, err
	}

	return fields, items, nil
}

/*
ts_code	str	TS股票代码
ann_date	str	公告日期
f_ann_date	str	实际公告日期，即发生过数据变更的最终日期
end_date	str	报告期
report_type	str	报告类型： 参考下表说明
comp_type	str	公司类型：1一般工商业 2银行 3保险 4证券
basic_eps	float	基本每股收益
diluted_eps	float	稀释每股收益
total_revenue	float	营业总收入 (元，下同)
revenue	float	营业收入
int_income	float	利息收入
prem_earned	float	已赚保费
comm_income	float	手续费及佣金收入
n_commis_income	float	手续费及佣金净收入
n_oth_income	float	其他经营净收益
n_oth_b_income	float	加:其他业务净收益
prem_income	float	保险业务收入
out_prem	float	减:分出保费
une_prem_reser	float	提取未到期责任准备金
reins_income	float	其中:分保费收入
n_sec_tb_income	float	代理买卖证券业务净收入
n_sec_uw_income	float	证券承销业务净收入
n_asset_mg_income	float	受托客户资产管理业务净收入
oth_b_income	float	其他业务收入
fv_value_chg_gain	float	加:公允价值变动净收益
invest_income	float	加:投资净收益
ass_invest_income	float	其中:对联营企业和合营企业的投资收益
forex_gain	float	加:汇兑净收益
total_cogs	float	营业总成本
oper_cost	float	减:营业成本
int_exp	float	减:利息支出
comm_exp	float	减:手续费及佣金支出
biz_tax_surchg	float	减:营业税金及附加
sell_exp	float	减:销售费用
admin_exp	float	减:管理费用
fin_exp	float	减:财务费用
assets_impair_loss	float	减:资产减值损失
prem_refund	float	退保金
compens_payout	float	赔付总支出
reser_insur_liab	float	提取保险责任准备金
div_payt	float	保户红利支出
reins_exp	float	分保费用
oper_exp	float	营业支出
compens_payout_refu	float	减:摊回赔付支出
insur_reser_refu	float	减:摊回保险责任准备金
reins_cost_refund	float	减:摊回分保费用
other_bus_cost	float	其他业务成本
operate_profit	float	营业利润
non_oper_income	float	加:营业外收入
non_oper_exp	float	减:营业外支出
nca_disploss	float	其中:减:非流动资产处置净损失
total_profit	float	利润总额
income_tax	float	所得税费用
n_income	float	净利润(含少数股东损益)
n_income_attr_p	float	净利润(不含少数股东损益)
minority_gain	float	少数股东损益
oth_compr_income	float	其他综合收益
t_compr_income	float	综合收益总额
compr_inc_attr_p	float	归属于母公司(或股东)的综合收益总额
compr_inc_attr_m_s	float	归属于少数股东的综合收益总额
ebit	float	息税前利润
ebitda	float	息税折旧摊销前利润
insurance_exp	float	保险业务支出
undist_profit	float	年初未分配利润
distable_profit	float	可分配利润
 */
func GetIncome(code, startDate, endDate string) ([]string, [][]string, []map[string]string, error) {
	filds := "ts_code,ann_date,f_ann_date,end_date,report_type,comp_type,basic_eps,diluted_eps,total_revenue,revenue,int_income,prem_earned,comm_income,n_commis_income,n_oth_income,n_oth_b_income,prem_income,out_prem,une_prem_reser,reins_income,n_sec_tb_income,n_sec_uw_income,n_asset_mg_income,oth_b_income,fv_value_chg_gain,invest_income,ass_invest_income,forex_gain,total_cogs,oper_cost,int_exp,comm_exp,biz_tax_surchg,sell_exp,admin_exp,fin_exp,assets_impair_loss,prem_refund,compens_payout,reser_insur_liab,div_payt,reins_exp,oper_exp,compens_payout_refu,insur_reser_refu,reins_cost_refund,other_bus_cost,operate_profit,non_oper_income,non_oper_exp,nca_disploss,total_profit,income_tax,n_income,n_income_attr_p,minority_gain,oth_compr_income,t_compr_income,compr_inc_attr_p,compr_inc_attr_m_s,ebit,ebitda,insurance_exp,undist_profit,distable_profit"
	//ps := &Params{
	//	Ts_Code:    "002594.SZ",
	//	Trade_Date: "20180726",
	//}
	ps := &Params{
		Ts_Code:    code,
		Start_Date: startDate,
		End_Date:   endDate,
	}
	fields, items, itemsMap, err := PostToUrl("income", filds, ps)

	if err != nil {
		return nil, nil, nil, err
	}

	return fields, items, itemsMap, nil
}

//获取所有股票的代码，及相关信息
func GetAllStocksCode() ([]string, [][]string, error) {
	filds := "ts_code,symbol,name,area,industry,fullname,enname,market,exchange,curr_type,list_status,list_date,delist_date,is_hs"

	ps := &Params{
		List_Status: "L",
		Exchange:    "",
	}
	fields, items, _, err := PostToUrl("stock_basic", filds, ps)

	if err != nil {
		return nil, nil, err
	}

	return fields, items, nil
}

//获取股票概念信息
func SaveConceptInfosToMySQL() {

	ps := &Params{}

	_, _, mapRes, err := PostToUrl("concept", "", ps)

	if err != nil {
		return
	}

	concepts := make([]*StockConcept, 0)

	for i := 0; i < len(mapRes); i++ {
		temp := &StockConcept{}
		commons.DataToStruct(mapRes[i], temp)
		concepts = append(concepts, temp)
	}

	sql := "replace into stock_concept(`concept_code`,`name`,`src`)VALUES"

	for i := 0; i < len(concepts); i++ {
		v := concepts[i]
		var tempStr string
		if i == len(concepts)-1 {
			tempStr = fmt.Sprintf("(\"%s\",\"%s\",\"%s\");", v.Code, v.Name, v.Src)
		} else {
			tempStr = fmt.Sprintf("(\"%s\",\"%s\",\"%s\"),", v.Code, v.Name, v.Src)
		}
		sql += tempStr
	}

	_, err = Engine.Exec(sql)
	if err != nil {
		logs.Error(err)
	}
}

//获取股票概念信息对应的股票，并且存储到数据库中
func SaveConceptDetailInfosToMySQL() {

	allConcepts := GetConceptInfosFromMySQL()

	for _, v := range allConcepts {
		ps := &Params{
			Id: v.Code,
		}

		_, _, mapRes, err := PostToUrl("concept_detail", "ts_code,name,in_date,out_date", ps)

		if err != nil {
			return
		}

		logs.Info(mapRes)

		var tempSql string = ""
		sql := "replace into stock_concept_detail(`concept_code`,`ts_code`,`name`,`out_date`,`in_date`)VALUES"
		//mapConceptDetails := make([]*StockConceptDetail, 0)
		for i := 0; i < len(mapRes); i++ {
			vv := mapRes[i]
			temp := &StockConceptDetail{}
			commons.DataToStruct(vv, temp)
			//mapConceptDetails = append(mapConceptDetails,temp)

			if i == len(mapRes)-1 {
				tempSql = fmt.Sprintf("(\"%s\",\"%s\",\"%s\",\"%s\",\"%s\");", v.Code, temp.TsCode, temp.Name, temp.OutDate, temp.InDate)
			} else {
				tempSql = fmt.Sprintf("(\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"),", v.Code, temp.TsCode, temp.Name, temp.OutDate, temp.InDate)
			}
			sql += tempSql

		}

		_, err = Engine.Exec(sql)
		if err != nil {
			panic(fmt.Sprintf("Engine.Exec %v Error %v", sql, err))
		}

		time.Sleep(1 * time.Second)

	}

}

func GetStockBalanceSheetFromTuShare(code, startDate, endDate string) ([]string, [][]string, []map[string]string, error) {

	ps := &Params{
		Ts_Code:    code,
		Start_Date:   startDate,
		End_Date: endDate,
	}

	fields, items, itemsMap, err := PostToUrl("balancesheet", "ts_code,f_ann_date,end_date,total_share", ps)

	if err != nil {
		return nil, nil, nil, err
	}

	return fields, items, itemsMap, err

}

//type StockBalanceSheet struct {
//	Ts_code     string  `json:"ts_code"`
//	F_ann_date  string  `json:"f_ann_date"`
//	End_date    string  `json:"end_date"`
//	Total_share float64 `json:"total_share"`
//}


func GetFinaIndicatorFromTuShare(code, startDate, endDate string) ([]string, [][]string, []map[string]string, error) {

	ps := &Params{
		Ts_Code:    code,
		Start_Date:   startDate,
		End_Date: endDate,
	}

	fields, items, itemsMap, err := PostToUrl("fina_indicator", "", ps)

	if err != nil {
		return nil, nil, nil, err
	}

	return fields, items, itemsMap, err

}

func SaveFinaIndicatorFromTuShare() {
	now := time.Now().Format("20060102")
	befor:="20000101"
	logs.Info(now)

	stocks, err := GetAllStocksInfo()
	if err != nil {
		panic(err)
	}

	for _, v := range stocks {
		_, _, finaIndicatordatas, err := GetFinaIndicatorFromTuShare(v.TsCode, befor, now)
		if err != nil {
			fmt.Errorf(err.Error())
			continue
		}

		_, _, balanceSheetdatas, err := GetStockBalanceSheetFromTuShare(v.TsCode, befor, now)
		if err != nil {
			fmt.Errorf(err.Error())
			continue
		}

		stockBalanceSheets:=make(map[string]*StockBalanceSheet,0)
		for _, vv := range balanceSheetdatas {
			temp := &StockBalanceSheet{}

			commons.DataToStruct(vv, temp)

			stockBalanceSheets[temp.End_date] = temp
		}

		stockDatas := make([]*StockFinaIndicator, 0)
		for _, vv := range finaIndicatordatas {
			temp := &StockFinaIndicator{}

			commons.DataToStruct(vv, temp)

			var total float64 = 0
			if sbs,ok:= stockBalanceSheets[temp.End_date];ok{
				total = sbs.Total_share
			}


			//totol:=stockBalanceSheets[temp.End_date].Total_share
			temp.Total_share = total
			stockDatas = append(stockDatas, temp)
		}

		//fmt.Println(stockDatas)

		var tempSql string = ""
		sql := "replace into stock_fina_indicator(`ts_code`,`ann_date`,`end_date`,`roe`,`roe_waa`,`roe_dt`,`roe_yearly`,`roe_avg`,`q_roe`,`q_dt_roe`,`roe_yoy`,`netprofit_margin`,`grossprofit_margin`,`profit_to_gr`,`op_of_gr`,`q_netprofit_margin`,`q_gsprofit_margin`,`q_profit_to_gr`,`q_op_to_gr`,`total_share`)VALUES"
		//mapConceptDetails := make([]*StockConceptDetail, 0)
		for i := 0; i < len(stockDatas); i++ {
			value := stockDatas[i]
			//temp := &StockConceptDetail{}
			//commons.DataToStruct(vv, temp)
			//mapConceptDetails = append(mapConceptDetails,temp)

			if i == len(stockDatas)-1 {
				tempSql = fmt.Sprintf("(\"%s\",\"%s\",\"%s\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\");", value.Ts_code, value.Ann_date, value.End_date, value.Roe,value.Roe_waa,value.Roe_dt,value.Roe_yearly,value.Roe_avg,value.Q_roe,value.Q_dt_roe,value.Roe_yoy,value.Netprofit_margin,value.Grossprofit_margin,value.Profit_to_gr,value.Op_of_gr,value.Q_netprofit_margin,value.Q_gsprofit_margin,value.Q_profit_to_gr,value.Q_op_to_gr,value.Total_share)
			} else {
				tempSql = fmt.Sprintf("(\"%s\",\"%s\",\"%s\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\",\"%f\"),", value.Ts_code, value.Ann_date, value.End_date, value.Roe,value.Roe_waa,value.Roe_dt,value.Roe_yearly,value.Roe_avg,value.Q_roe,value.Q_dt_roe,value.Roe_yoy,value.Netprofit_margin,value.Grossprofit_margin,value.Profit_to_gr,value.Op_of_gr,value.Q_netprofit_margin,value.Q_gsprofit_margin,value.Q_profit_to_gr,value.Q_op_to_gr,value.Total_share)
			}
			sql += tempSql

		}

		_, err = Engine.Exec(sql)
		if err != nil {
			panic(fmt.Sprintf("Engine.Exec %v Error %v", sql, err))
		}
		time.Sleep(1 * time.Second)
	}
}

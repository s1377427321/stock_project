package structs

import "fmt"

//163获取股票历史记录
var DAY_TRADE_API = "http://quotes.money.163.com/service/chddata.html?code=%s&start=%s&end=%s"

var HTTP_URL = "http://api.tushare.pro"

var TOKEN = "6ad2bcae0a39b5feab53acb555a149080db53e4f0640492485cbb8ce"

var SQLParams = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456", "120.79.154.53:3306", "stock")

type ReqParams struct {
	ApiName string `json:"api_name"`
	Token   string `json:"token"`
	Params  Params `json:"params"`
	Fields  string `json:"fields"`
}

type Params struct {
	Ts_Code     string `json:"ts_code"`
	Trade_Date  string `json:"trade_date"`
	Start_Date  string `json:"start_date"`
	End_Date    string `json:"end_date"`
	Ann_Date    string `json:"ann_date"`
	Period      string `json:"period"`
	Report_Type string `json:"report_type"`
	Comp_Type   string `json:"comp_type"`
	List_Status string `json:"list_status"`
	Exchange    string `json:"exchange"`
	Test        string `json:"test"`
}

type Respons struct {
	Request_Id string `json:"request_id"`
	Code       int    `json:"code"`
	Msg        bool   `json:"msg"`
	Data       string `json:"data"`
}

type StockItem struct {
	Code    string
	Name    string
	BuyNums float64
}

//buy_statics数据库处理字段
type BuyStaticsInfo struct {
	Code                     string  `json:"code"`
	CodeName                 string  `json:"code_name"`
	FixedInvestmentMoney     float64 `json:"fixed_investment_money"`
	FirstProfitabilityRate   float64 `json:"first_profitability_rate"`
	CurrentProfitabilityRate float64 `json:"current_profitability_rate"`
	Pe                       float64 `json:"pe"`
	Pb                       float64 `json:"pb"`
	DividentYield            float64 `json:"divident_yield"`
	Roe                      float64 `json:"roe"`
	SecondPower              float64 `json:"second_power"`
}

type StockBasicInfo struct {
	Id         int    `json:"id"`
	TsCode     string `json:"ts_code"`     //TS代码
	Symbol     string `json:"symbol"`      //股票代码
	Name       string `json:"name"`        //股票名称
	Area       string `json:"area"`        //所在地域
	Industry   string `json:"industry"`    //str	所属行业
	Fullname   string `json:"fullname"`    //str	股票全称
	Enname     string `json:"enname"`      //英文全称
	Market     string `json:"market"`      //市场类型 （主板/中小板/创业板）
	Exchange   string `json:"exchange"`    //str	交易所代码
	CurrType   string `json:"curr_type"`   //str	交易货币
	ListStatus string `json:"list_status"` //str	上市状态： L上市 D退市 P暂停上市
	ListDate   string `json:"list_date"`   //str	上市日期
	DelistDate string `json:"delist_date"` //str	退市日期
	IsHs       string `json:"is_hs"`       //是否沪深港通标的，N否 H沪股通 S深股通
}

//每日数据
type DialyStockInfo struct {
	Id     int     `json:"id"`
	Code   int     `json:"code"`
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Volume int     `json:"volume"`
	Money  int     `json:"money"`
}

type StockIncome struct {
	Ts_code             string  `json:"ts_code"`
	Ann_date            string  `json:"ann_date"`
	F_ann_date          string  `json:"f_ann_date"`
	End_date            string  `json:"end_date"`
	Report_type         string  `json:"report_type"`
	Comp_type           string  `json:"comp_type"`
	Basic_eps           float64 `json:"basic_eps"`
	Diluted_eps         float64 `json:"diluted_eps"`
	Total_revenue       float64 `json:"total_revenue"`
	Revenue             float64 `json:"revenue"`
	Int_income          float64 `json:"int_income"`
	Prem_earned         float64 `json:"prem_earned"`
	Comm_income         float64 `json:"comm_income"`
	N_commis_income     float64 `json:"n_commis_income"`
	N_oth_income        float64 `json:"n_oth_income"`
	N_oth_b_income      float64 `json:"n_oth_b_income"`
	Prem_income         float64 `json:"prem_income"`
	Out_prem            float64 `json:"out_prem"`
	Une_prem_reser      float64 `json:"une_prem_reser"`
	Reins_income        float64 `json:"reins_income"`
	N_sec_tb_income     float64 `json:"n_sec_tb_income"`
	N_sec_uw_income     float64 `json:"n_sec_uw_income"`
	N_asset_mg_income   float64 `json:"n_asset_mg_income"`
	Oth_b_income        float64 `json:"oth_b_income"`
	Fv_value_chg_gain   float64 `json:"fv_value_chg_gain"`
	Invest_income       float64 `json:"invest_income"`
	Ass_invest_income   float64 `json:"ass_invest_income"`
	Forex_gain          float64 `json:"forex_gain"`
	Total_cogs          float64 `json:"total_cogs"`
	Oper_cost           float64 `json:"oper_cost"`
	Int_exp             float64 `json:"int_exp"`
	Comm_exp            float64 `json:"comm_exp"`
	Biz_tax_surchg      float64 `json:"biz_tax_surchg"`
	Sell_exp            float64 `json:"sell_exp"`
	Admin_exp           float64 `json:"admin_exp"`
	Fin_exp             float64 `json:"fin_exp"`
	Assets_impair_loss  float64 `json:"assets_impair_loss"`
	Prem_refund         float64 `json:"prem_refund"`
	Compens_payout      float64 `json:"compens_payout"`
	Reser_insur_liab    float64 `json:"reser_insur_liab"`
	Div_payt            float64 `json:"div_payt"`
	Reins_exp           float64 `json:"reins_exp"`
	Oper_exp            float64 `json:"oper_exp"`
	Compens_payout_refu float64 `json:"compens_payout_refu"`
	Insur_reser_refu    float64 `json:"insur_reser_refu"`
	Reins_cost_refund   float64 `json:"reins_cost_refund"`
	Other_bus_cost      float64 `json:"other_bus_cost"`
	Operate_profit      float64 `json:"operate_profit"`
	Non_oper_income     float64 `json:"non_oper_income"`
	Non_oper_exp        float64 `json:"non_oper_exp"`
	Nca_disploss        float64 `json:"nca_disploss"`
	Total_profit        float64 `json:"total_profit"`
	Income_tax          float64 `json:"income_tax"`
	N_income            float64 `json:"n_income"`
	N_income_attr_p     float64 `json:"n_income_attr_p"`
	Minority_gain       float64 `json:"minority_gain"`
	Oth_compr_income    float64 `json:"oth_compr_income"`
	T_compr_income      float64 `json:"t_compr_income"`
	Compr_inc_attr_p    float64 `json:"compr_inc_attr_p"`
	Compr_inc_attr_m_s  float64 `json:"compr_inc_attr_m_s"`
	Ebit                float64 `json:"ebit"`
	Ebitda              float64 `json:"ebitda"`
	Insurance_exp       float64 `json:"insurance_exp"`
	Undist_profit       float64 `json:"undist_profit"`
	Distable_profit     float64 `json:"distable_profit"`
}

type IncomeS struct {
	Data      []string
	Time      int
	PerInCome float64
}

type IncomeSS []*IncomeS

func (s IncomeSS) Len() int {
	return len(s)
}

func (s IncomeSS) Less(i, j int) bool {
	return s[i].Time > s[j].Time
}

//Swap()
func (s IncomeSS) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

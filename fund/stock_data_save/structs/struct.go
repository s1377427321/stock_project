package structs

import (
	"fmt"
	"strconv"
)

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
	Id          string `json:"id"` //concept id
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

/*
名称	类型	描述
ts_code					str	          	TS股票代码
ann_date	          	str	公告日期
f_ann_date	          	str	实际公告日期，即发生过数据变更的最终日期
end_date	          	str	报告期
report_type	          	str	报告类型： 参考下表说明
comp_type	          	str	公司类型：1一般工商业 2银行 3保险 4证券
basic_eps	          	float	基本每股收益
diluted_eps	          	float	稀释每股收益
total_revenue	 	  	float	营业总收入 (元，下同)
revenue				  	float	 	  营业收入
int_income			  	float		利息收入
prem_earned			  	float		已赚保费
comm_income			  	float		手续费及佣金收入
n_commis_income	float	手续费及佣金净收入
n_oth_income	float	其他经营净收益
n_oth_b_income	float	加:其他业务净收益
prem_income	float		保险业务收入
out_prem	float		减:分出保费
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

/*
fina_indicator

ts_code	str	TS代码
ann_date	str	公告日期
end_date	str	报告期

roe	float	净资产收益率
roe_waa	float	加权平均净资产收益率
roe_dt	float	净资产收益率(扣除非经常损益)
roe_yearly	float	年化净资产收益率
roe_avg	float	平均净资产收益率(增发条件)
q_roe	float	净资产收益率(单季度)
q_dt_roe	float	净资产单季度收益率(扣除非经常损益)
roe_yoy	float	净资产收益率(摊薄)同比增长率(%)

netprofit_margin	float	销售净利率
grossprofit_margin	float	销售毛利率
profit_to_gr	float	净利润/营业总收入
op_of_gr	float	营业利润/营业总收入
q_netprofit_margin	float	销售净利率(单季度)
q_gsprofit_margin	float	销售毛利率(单季度)
q_profit_to_gr	float	净利润／营业总收入(单季度)
q_op_to_gr	float	营业利润／营业总收入(单季度)



roic	float	投入资本回报率
eps	float	基本每股收益
dt_eps	float	稀释每股收益
total_revenue_ps	float	每股营业总收入
revenue_ps	float	每股营业收入
capital_rese_ps	float	每股资本公积
surplus_rese_ps	float	每股盈余公积
undist_profit_ps	float	每股未分配利润
extra_item	float	非经常性损益
profit_dedt	float	扣除非经常性损益后的净利润
gross_margin	float	毛利
current_ratio	float	流动比率
quick_ratio	float	速动比率
cash_ratio	float	保守速动比率
invturn_days	float	存货周转天数
arturn_days	float	应收账款周转天数
inv_turn	float	存货周转率
ar_turn	float	应收账款周转率
ca_turn	float	流动资产周转率
fa_turn	float	固定资产周转率
assets_turn	float	总资产周转率
op_income	float	经营活动净收益
valuechange_income	float	价值变动净收益
interst_income	float	利息费用
daa	float	折旧与摊销
ebit	float	息税前利润
ebitda	float	息税折旧摊销前利润
fcff	float	企业自由现金流量
fcfe	float	股权自由现金流量
current_exint	float	无息流动负债
noncurrent_exint	float	无息非流动负债
interestdebt	float	带息债务
netdebt	float	净债务
tangible_asset	float	有形资产
working_capital	float	营运资金
networking_capital	float	营运流动资本
invest_capital	float	全部投入资本
retained_earnings	float	留存收益
diluted2_eps	float	期末摊薄每股收益
bps	float	每股净资产
ocfps	float	每股经营活动产生的现金流量净额
retainedps	float	每股留存收益
cfps	float	每股现金流量净额
ebit_ps	float	每股息税前利润
fcff_ps	float	每股企业自由现金流量
fcfe_ps	float	每股股东自由现金流量

cogs_of_sales	float	销售成本率
expense_of_sales	float	销售期间费用率


saleexp_to_gr	float	销售费用/营业总收入
adminexp_of_gr	float	管理费用/营业总收入
finaexp_of_gr	float	财务费用/营业总收入
impai_ttm	float	资产减值损失/营业总收入
gc_of_gr	float	营业总成本/营业总收入
ebit_of_gr	float	息税前利润/营业总收入


roa	float	总资产报酬率
npta	float	总资产净利润

roa2_yearly	float	年化总资产报酬率

opincome_of_ebt	float	经营活动净收益/利润总额
investincome_of_ebt	float	价值变动净收益/利润总额
n_op_profit_of_ebt	float	营业外收支净额/利润总额
tax_to_ebt	float	所得税/利润总额
dtprofit_to_profit	float	扣除非经常损益后的净利润/净利润
salescash_to_or	float	销售商品提供劳务收到的现金/营业收入
ocf_to_or	float	经营活动产生的现金流量净额/营业收入
ocf_to_opincome	float	经营活动产生的现金流量净额/经营活动净收益
capitalized_to_da	float	资本支出/折旧和摊销
debt_to_assets	float	资产负债率
assets_to_eqt	float	权益乘数
dp_assets_to_eqt	float	权益乘数(杜邦分析)
ca_to_assets	float	流动资产/总资产
nca_to_assets	float	非流动资产/总资产
tbassets_to_totalassets	float	有形资产/总资产
int_to_talcap	float	带息债务/全部投入资本
eqt_to_talcapital	float	归属于母公司的股东权益/全部投入资本
currentdebt_to_debt	float	流动负债/负债合计
longdeb_to_debt	float	非流动负债/负债合计
ocf_to_shortdebt	float	经营活动产生的现金流量净额/流动负债
debt_to_eqt	float	产权比率
eqt_to_debt	float	归属于母公司的股东权益/负债合计
eqt_to_interestdebt	float	归属于母公司的股东权益/带息债务
tangibleasset_to_debt	float	有形资产/负债合计
tangasset_to_intdebt	float	有形资产/带息债务
tangibleasset_to_netdebt	float	有形资产/净债务
ocf_to_debt	float	经营活动产生的现金流量净额/负债合计
ocf_to_interestdebt	float	经营活动产生的现金流量净额/带息债务
ocf_to_netdebt	float	经营活动产生的现金流量净额/净债务
ebit_to_interest	float	已获利息倍数(EBIT/利息费用)
longdebt_to_workingcapital	float	长期债务与营运资金比率
ebitda_to_debt	float	息税折旧摊销前利润/负债合计
turn_days	float	营业周期
roa_yearly	float	年化总资产净利率
roa_dp	float	总资产净利率(杜邦分析)
fixed_assets	float	固定资产合计
profit_prefin_exp	float	扣除财务费用前营业利润
non_op_profit	float	非营业利润
op_to_ebt	float	营业利润／利润总额
nop_to_ebt	float	非营业利润／利润总额
ocf_to_profit	float	经营活动产生的现金流量净额／营业利润
cash_to_liqdebt	float	货币资金／流动负债
cash_to_liqdebt_withinterest	float	货币资金／带息流动负债
op_to_liqdebt	float	营业利润／流动负债
op_to_debt	float	营业利润／负债合计
roic_yearly	float	年化投入资本回报率
profit_to_op	float	利润总额／营业收入
q_opincome	float	经营活动单季度净收益
q_investincome	float	价值变动单季度净收益
q_dtprofit	float	扣除非经常损益后的单季度净利润
q_eps	float	每股收益(单季度)


q_exp_to_sales	float	销售期间费用率(单季度)

q_saleexp_to_gr	float	销售费用／营业总收入 (单季度)
q_adminexp_to_gr	float	管理费用／营业总收入 (单季度)
q_finaexp_to_gr	float	财务费用／营业总收入 (单季度)
q_impair_to_gr_ttm	float	资产减值损失／营业总收入(单季度)
q_gc_to_gr	float	营业总成本／营业总收入 (单季度)



q_npta	float	总资产净利润(单季度)
q_opincome_to_ebt	float	经营活动净收益／利润总额(单季度)
q_investincome_to_ebt	float	价值变动净收益／利润总额(单季度)
q_dtprofit_to_profit	float	扣除非经常损益后的净利润／净利润(单季度)
q_salescash_to_or	float	销售商品提供劳务收到的现金／营业收入(单季度)
q_ocf_to_sales	float	经营活动产生的现金流量净额／营业收入(单季度)
q_ocf_to_or	float	经营活动产生的现金流量净额／经营活动净收益(单季度)
basic_eps_yoy	float	基本每股收益同比增长率(%)
dt_eps_yoy	float	稀释每股收益同比增长率(%)
cfps_yoy	float	每股经营活动产生的现金流量净额同比增长率(%)
op_yoy	float	营业利润同比增长率(%)
ebt_yoy	float	利润总额同比增长率(%)
netprofit_yoy	float	归属母公司股东的净利润同比增长率(%)
dt_netprofit_yoy	float	归属母公司股东的净利润-扣除非经常损益同比增长率(%)
ocf_yoy	float	经营活动产生的现金流量净额同比增长率(%)

bps_yoy	float	每股净资产相对年初增长率(%)
assets_yoy	float	资产总计相对年初增长率(%)
eqt_yoy	float	归属母公司的股东权益相对年初增长率(%)
tr_yoy	float	营业总收入同比增长率(%)
or_yoy	float	营业收入同比增长率(%)
q_gr_yoy	float	营业总收入同比增长率(%)(单季度)
q_gr_qoq	float	营业总收入环比增长率(%)(单季度)
q_sales_yoy	float	营业收入同比增长率(%)(单季度)
q_sales_qoq	float	营业收入环比增长率(%)(单季度)
q_op_yoy	float	营业利润同比增长率(%)(单季度)
q_op_qoq	float	营业利润环比增长率(%)(单季度)
q_profit_yoy	float	净利润同比增长率(%)(单季度)
q_profit_qoq	float	净利润环比增长率(%)(单季度)
q_netprofit_yoy	float	归属母公司股东的净利润同比增长率(%)(单季度)
q_netprofit_qoq	float	归属母公司股东的净利润环比增长率(%)(单季度)
equity_yoy	float	净资产同比增长率
rd_exp	float	研发费用
 */
//财务指标数据
type StockFinaIndicator struct {
	Ts_code    string  `json:"ts_code"`
	Ann_date   string  `json:"ann_date"`
	End_date   string  `json:"end_date"`
	Roe        float64 `json:"roe"`
	Roe_waa    float64 `json:"roe_waa"`
	Roe_dt     float64 `json:"roe_dt"`
	Roe_yearly float64 `json:"roe_yearly"`
	Roe_avg    float64 `json:"roe_avg"`
	Q_roe      float64 `json:"q_roe"`
	Q_dt_roe   float64 `json:"q_dt_roe"`
	Roe_yoy    float64 `json:"roe_yoy"`

	Netprofit_margin   float64 `json:"netprofit_margin"`
	Grossprofit_margin float64 `json:"grossprofit_margin"`
	Profit_to_gr       float64 `json:"profit_to_gr"`
	Op_of_gr           float64 `json:"op_of_gr"`
	Q_netprofit_margin float64 `json:"q_netprofit_margin"`
	Q_gsprofit_margin  float64 `json:"q_gsprofit_margin"`
	Q_profit_to_gr     float64 `json:"q_profit_to_gr"`
	Q_op_to_gr         float64 `json:"q_op_to_gr"`

	Total_share float64 `json:"total_share"`
}

type StockInComes []*StockIncome

func (s StockInComes) Len() int {
	return len(s)
}

func (s StockInComes) Less(i, j int) bool {
	ii, err := strconv.Atoi(s[i].End_date)
	if err != nil {
		panic("StockInComes Less ERROR" + err.Error())
	}

	jj, err := strconv.Atoi(s[j].End_date)
	if err != nil {
		panic("StockInComes Less ERROR" + err.Error())
	}

	return ii > jj
}

//Swap()
func (s StockInComes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
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

/*
code	str	Y	概念分类ID
name	str	Y	概念分类名称
src	str	Y	来源
 */
type StockConcept struct {
	Code string `json:"concept_code"`
	Name string `json:"name"`
	Src  string `json:"src"`
}

/*
id	str	Y	概念代码
ts_code	str	Y	股票代码
name	str	Y	股票名称
in_date	str	N	纳入日期
out_date	str	N	剔除日期
接口使用


pro = ts.pro_api()

#取5G概念明细
df = pro.concept_detail(id='TS2', fields='ts_code,name')
 */
type StockConceptDetail struct {
	ConceptCode string `json:"concept_code"`
	TsCode      string `json:"ts_code"`
	Name        string `json:"name"`
	InDate      string `json:"in_date"`
	OutDate     string `json:"out_date"`
}

type IncomeGoodS struct {
	Des   string
	Order float64
}

type IncomeGoodSS []*IncomeGoodS

func (s IncomeGoodSS) Len() int {
	return len(s)
}

func (s IncomeGoodSS) Less(i, j int) bool {
	return s[i].Order > s[j].Order
}

//Swap()
func (s IncomeGoodSS) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type StockBalanceSheet struct {
	Ts_code     string  `json:"ts_code"`
	F_ann_date  string  `json:"f_ann_date"`
	End_date    string  `json:"end_date"`
	Total_share float64 `json:"total_share"`
}


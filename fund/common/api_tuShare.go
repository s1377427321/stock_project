package common

import (
	. "fund/common/structs"
)

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

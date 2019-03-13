package common

import (
	"github.com/go-xorm/xorm"
	"fmt"
	"strconv"
	"math"
	. "fund/common/structs"
)

//获取数据库表buy_statics数据，整理出投资金额
func GetDealBuyStaticsInfos() {
	var err error
	Engine, err := xorm.NewEngine("mysql", SQLParams)
	if err != nil {
		panic(err)
	}
	defer Engine.Close()

	sqls := fmt.Sprintf("select * from %s where code=%s", "buy_statics", "510880")
	result, err := Engine.Query(sqls)
	if err != nil {
		fmt.Println(err)
	}

	CodeInfos := make([]*BuyStaticsInfo, 0)

	for _, v := range result {
		newCf := &BuyStaticsInfo{}
		newCf.Code = string(v["code"])
		newCf.CodeName = string(v["code_name"])
		fim, err := strconv.ParseFloat(string(v["fixed_investment_money"]), 64)
		if err != nil {
			//panic(err)
			newCf.FixedInvestMentMoney = 0
		} else {
			newCf.FixedInvestMentMoney = fim
		}

		fpr, err := strconv.ParseFloat(string(v["first_profitability_rate"]), 64)
		if err != nil {
			//panic(err)
			newCf.FirstProfitabilityRate = 0
		} else {
			newCf.FirstProfitabilityRate = fpr
		}

		cpr, err := strconv.ParseFloat(string(v["current_profitability_rate"]), 64)
		if err != nil {
			//panic(err)
			newCf.CurretProfitabilityRate = 0
		} else {
			newCf.CurretProfitabilityRate = cpr
		}

		pe, err := strconv.ParseFloat(string(v["pe"]), 64)
		if err != nil {
			//panic(err)
			newCf.Pe = 0
		} else {
			newCf.Pe = pe
		}

		pb, err := strconv.ParseFloat(string(v["pb"]), 64)
		if err != nil {
			//panic(err)
			newCf.Pb = 0
		} else {
			newCf.Pb = pb
		}

		dy, err := strconv.ParseFloat(string(v["dividend_yield"]), 64)
		if err != nil {
			//panic(err)
			newCf.DividentYield = 0
		} else {
			newCf.DividentYield = dy
		}

		roe, err := strconv.ParseFloat(string(v["roe"]), 64)
		if err != nil {
			//panic(err)
			newCf.Roe = 0
		} else {
			newCf.Roe = roe
		}

		//
		//
		//for kk, vv := range v {
		//	fmt.Println(kk)
		//	fmt.Println(string(vv))
		//}
		CodeInfos = append(CodeInfos, newCf)
	}

	for _, v := range CodeInfos {
		DealProfitabilityRate(v)
	}

}

//盈利收益率算法
func DealProfitabilityRate(v *BuyStaticsInfo) {
	var coefficient float64
	if v.CurretProfitabilityRate > v.FirstProfitabilityRate {
		coefficient = math.Pow(v.CurretProfitabilityRate/v.FirstProfitabilityRate, 1.5)
	} else {
		coefficient = math.Pow(v.CurretProfitabilityRate/v.FirstProfitabilityRate, 1.5)
	}
	buyMoney := v.FixedInvestMentMoney * coefficient

	showPrint := fmt.Sprintf("%s %s 买入金额：%v,买入的盈利收益率：%v,当前盈利收益率：%v，现在买入金额：%v,", v.CodeName, v.Code, v.FixedInvestMentMoney,
		v.FirstProfitabilityRate, v.CurretProfitabilityRate, buyMoney)
	fmt.Println(showPrint)
}
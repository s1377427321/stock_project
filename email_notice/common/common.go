package common

import (
	"common"
)

/*
华宝添益
120.79.154.53:4661/add?code=sh511990&height=100.05&low=99.905
白云机场
120.79.154.53:4661/addstock?code=sh600004&money=130000&copies=20&price=14.844
伊利股份
120.79.154.53:4661/addstock?code=sh600887&money=130000&copies=10&price=25.18
招商银行
120.79.154.53:4661/addstock?code=sh600036&money=130000&copies=10&price=29.32
复星医药
120.79.154.53:4661/addstock?code=sh600196&money=130000&copies=10&price=25.76
贵州茅台
120.79.154.53:4661/addstock?code=sh600519&money=130000&copies=10&price=717.92
*/

type InitNoticeStocksS struct {
	Name   string
	Code   string
	Height float64
	Low    float64
	Money  float64
}

type InitBuyStocksS struct {
	Name   string
	Code   string
	Money  float64
	Copies int
	Price  float64
}

func IntToFloat64(n int) float64 {
	r := common.Round64(float64(n), 0)
	return r
}

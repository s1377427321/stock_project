package common

type StockRange struct {
	*StockCommon
	IntervalHigh float64 //区间高点
	IntervalLow  float64 //区间低点

	AllMoney           float64
	Part               int     //份数
	FinalStopLossPrice float64 //最终止损
	FinalStopLossMoney float64 //止损亏的钱
	HoldDayNums        int     //持有天数
}

func (this *StockRange) Init(allMoney, StopLose float64, part int, code, name string, buyMoney float64, intervalHigh, intervalLow float64, ) {

}

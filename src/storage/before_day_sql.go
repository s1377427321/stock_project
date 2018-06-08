package storage


func InsertResult(code int,start_day string ,befor_day int ,total_money ,win_money,one_money,remain_money,stock_money,stock_price,share_holding float64,
	tactics string ,do_tactics_num int ,tactics_win float64,tactics_change string,updata_time string )  {
	sqlString:=`
	INSERT INTO befor_day
	(
		code,
		start_day,
		befor_day,
		total_money,
		win_money,
		one_money,
		remain_money,
		stock_money,
		stock_price,
		share_holding,
		tactics,
		do_tactics_num,
		tactics_win,
		tactics_change,
		updata_time
	) VALUES (
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?
	) 
	ON DUPLICATE KEY UPDATE 
	start_day=?,befor_day=?,total_money=?,win_money=?,one_money=?,remain_money=?,stock_money=?,stock_price=?,share_holding=?,tactics=?,do_tactics_num=?,
	tactics_win=?,tactics_change=?,updata_time=?
`
	//,befor_day,total_money,one_money,remain_money,stock_money,stock_price,share_holding,processes`
	//INSERT INTO pro_realtimeprice_rank (id, id_s, plat,d_m) VALUES (1, '5445', '5454', 5) ON DUPLICATE KEY UPDATE d_m=d_m+5;
	stmt, err := db.Prepare(sqlString)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(code,start_day,befor_day,total_money,win_money,one_money,remain_money,stock_money,stock_price,share_holding,tactics,do_tactics_num,tactics_win,tactics_change,updata_time,
		start_day,befor_day,total_money,win_money,one_money,remain_money,stock_money,stock_price,share_holding,tactics,do_tactics_num,tactics_win,tactics_change,updata_time)
	if err != nil {
		panic(err.Error())
	}
}




//
//func GetLatestResultForBeforeDay(code int) *model.BeforeDayStruct {
//	sqlstr:=`
//		select code,start_day,befor_day,total_money,win_money,one_money,remain_money,
//       stock_money,stock_price,share_holding,stactics,processesNum,processes,
//       processesWin,upDataTime FROM befor_day WHERE code=?
//	`
//	stmt,err:=db.Prepare(sqlstr)
//
//	if err != nil{
//		panic(err.Error())
//	}
//
//	defer stmt.Close()
//
//	needStruct:=&model.BeforeDayStruct{
//	}
//
//	row:= stmt.QueryRow(code)
//	if row == nil {
//		return needStruct
//	}
//
//	err = row.Scan(
//		&needStruct.Code,
//		&needStruct.StartDay,
//		&needStruct.BeforeDay,
//		&needStruct.TotalMoney,
//		&needStruct.WinMoney,
//		&needStruct.OneMoney,
//		&needStruct.RemainMoney,
//		&needStruct.StockMoney,
//		&needStruct.StockPrice,
//		&needStruct.ShareHolding,
//		&needStruct.Tactics,
//		&needStruct.DoTacticsNums,
//		&needStruct.TacticsChange,
//		&needStruct.TacticsWin,
//		&needStruct.UpdateDay,
//	)
//
//	if err != nil {
//		return needStruct
//	}
//
//	return  needStruct
//}

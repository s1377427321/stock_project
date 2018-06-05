package storage

import (
	//"tactics/percentTen"
	//"tactics/percentTen"
)

func InsertResult(code int,start_day string ,befor_day int ,total_money ,one_money,remain_money,stock_money,stock_price,share_holding float64,strStactics,strProcess string)  {
	sqlString:=`
	INSERT INTO befor_day
	(
		code,
		start_day,
		befor_day,
		total_money,
		one_money,
		remain_money,
		stock_money,
		stock_price,
		share_holding,
		stactics,
		processes
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
		?
	) 
	ON DUPLICATE KEY UPDATE 
	start_day=?,befor_day=?,total_money=?,one_money=?,remain_money=?,stock_money=?,stock_price=?,share_holding=?,stactics=?,processes=?
`
//,befor_day,total_money,one_money,remain_money,stock_money,stock_price,share_holding,processes`
	//INSERT INTO pro_realtimeprice_rank (id, id_s, plat,d_m) VALUES (1, '5445', '5454', 5) ON DUPLICATE KEY UPDATE d_m=d_m+5;
	stmt, err := db.Prepare(sqlString)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(code,start_day,befor_day,total_money,one_money,remain_money,stock_money,stock_price,share_holding,strStactics,strProcess,
		start_day,befor_day,total_money,one_money,remain_money,stock_money,stock_price,share_holding,strStactics,strProcess)
	if err != nil {
		panic(err.Error())
	}
}

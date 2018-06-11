package storage

import (
	"model"
	"sync"
)

var befor_day_sql_lock *sync.Mutex = &sync.Mutex{}

var allInsertDayStruct []*model.BeforeDayStruct=make([]*model.BeforeDayStruct,0)

func AddInsertDay(day *model.BeforeDayStruct)  {
	befor_day_sql_lock.Lock()
	allInsertDayStruct=append(allInsertDayStruct,day)
	befor_day_sql_lock.Unlock()
}

func InsertAll()  {
	if len(allInsertDayStruct) == 0 {
		return
	}

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
	stmt, err := db.Prepare(sqlString)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	begin, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}

	defer begin.Rollback()

	for _, dayStruct := range allInsertDayStruct {
		if dayStruct == nil {
			continue
		}

		_, err :=begin.Stmt(stmt).Exec(dayStruct.Code,dayStruct.StartDay,dayStruct.BeforeDay,dayStruct.TotalMoney,dayStruct.WinMoney,dayStruct.OneMoney,dayStruct.RemainMoney,dayStruct.StockMoney,dayStruct.StockPrice,dayStruct.ShareHolding,dayStruct.Tactics,dayStruct.DoTacticsNums,dayStruct.TacticsWin,dayStruct.TacticsChange,dayStruct.UpdateDay,
			dayStruct.StartDay,dayStruct.BeforeDay,dayStruct.TotalMoney,dayStruct.WinMoney,dayStruct.OneMoney,dayStruct.RemainMoney,dayStruct.StockMoney,dayStruct.StockPrice,dayStruct.ShareHolding,dayStruct.Tactics,dayStruct.DoTacticsNums,dayStruct.TacticsWin,dayStruct.TacticsChange,dayStruct.UpdateDay)


		if err != nil {
			panic(err.Error())
		}
	}

	err = begin.Commit()
	if err != nil {
		panic(err.Error())
	}
}



//获取从今天到前面的几天
func GetLatestSomeDataFromDay(code ,dayNum int) map[string]*model.DayTrade  {
	sql:=`
	SELECT 
		code, 
		date,
		open,
		close,
		high,
		low,
		volume,
		money
	FROM trade_his
	WHERE code = ?
	ORDER BY date DESC
	LIMIT ?
	`
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	days :=  make(map[string]*model.DayTrade,1)
	rows,err := stmt.Query(code,dayNum)

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		day := &model.DayTrade{}
		err = rows.Scan(
			&day.Code,
			&day.Date,
			&day.Open,
			&day.Close,
			&day.High,
			&day.Low,
			&day.Volume,
			&day.Money,
		)
		if err != nil {
			return nil
		}
		days[day.Date] = day
	}
	//log.Println("-----return--------GetLatestSomeDataFromDay----------------")
	return days
}


func IsAllRecordIsOver() bool  {
	ret:=false


	return ret

}

func GetBeforDayCount() int  {
	ret:=0

	return ret
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

//
///获取从今天到前面的几天
//func GetAllLatestSomeDataFromDay(stocks  []*model.Stock,dayNum int) map[int]*model.DayTrade  {
//	//code ,dayNum int
//	sql:=`
//	SELECT
//		code,
//		date,
//		open,
//		close,
//		high,
//		low,
//		volume,
//		money
//	FROM trade_his
//	WHERE code = ?
//	ORDER BY date DESC
//	LIMIT ?
//	`
//	stmt, err := db.Prepare(sql)
//
//	if err != nil {
//		panic(err.Error())
//	}
//	defer stmt.Close()
//
//	begin, err := db.Begin()
//	if err != nil {
//		panic(err.Error())
//	}
//	defer begin.Rollback()
//
//	for _,v:=range stocks{
//		begin.Stmt(stmt).Query(v.Id,dayNum)
//	}
//
//	err= begin.Commit()
//	if err != nil{
//		panic(err.Error())
//	}
//
//
//
//
//
//	days :=  make(map[string]*model.DayTrade,1)
//	rows,err := stmt.Query(code,dayNum)
//
//	if err != nil {
//		panic(err.Error())
//	}
//
//	for rows.Next() {
//		day := &model.DayTrade{}
//		err = rows.Scan(
//			&day.Code,
//			&day.Date,
//			&day.Open,
//			&day.Close,
//			&day.High,
//			&day.Low,
//			&day.Volume,
//			&day.Money,
//		)
//		if err != nil {
//			return nil
//		}
//		days[day.Date] = day
//	}
//	//log.Println("-----return--------GetLatestSomeDataFromDay----------------")
//	return days
//}
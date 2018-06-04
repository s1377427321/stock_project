package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"model"
	"fmt"
	"constant"
	"log"
)

var db = NewMysql()

func NewMysql() *sql.DB {
	dbData:= fmt.Sprintf("%s:%s@tcp(%s)/%s", constant.DB_USER, constant.DB_PASSWORD,constant.DB_IP ,constant.DB_NAME)
	db, err := sql.Open("mysql", dbData)
	//db, err := sql.Open("mysql", "root:@tcp(120.79.154.53:3306)/stock?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func InsertTradeHis(days []*model.DayTrade) error {
	//log.Println("day.Code------------start--------------------")
	sql := `
	INSERT INTO trade_his
	(
		code,
		date,
		open,
		close,
		high,
		low,
		volume,
		money
	) VALUES (
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?
	)`

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}

	begin, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}


	defer stmt.Close()



	for _, day := range days {
		if day == nil {
			continue
		}
		log.Println(day)

		_, err :=begin.Stmt(stmt).Exec(day.Code, day.Date,
			day.Open, day.Close, day.High, day.Low,
			day.Volume, day.Money)


		if err != nil {
			panic(err.Error())
		}
	}

	err = begin.Commit()
	if err != nil {
		panic(err.Error())
	}
	//log.Println("!!!!!!!!!!!!!!!!!!end!!!!!!!!!!!!!!!!!!")

	return nil
}

//func InsertTradeHis(days []*model.DayTrade) error {
//	sql := `
//	INSERT INTO trade_his
//	(
//		code,
//		date,
//		open,
//		close,
//		high,
//		low,
//		volume,
//		money
//	) VALUES (
//		?,
//		?,
//		?,
//		?,
//		?,
//		?,
//		?,
//		?
//	)`
//
//	stmt, err := db.Prepare(sql)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	begin, err := db.Begin()
//	if err != nil {
//		panic(err.Error())
//	}
//
//
//	defer stmt.Close()
//
//
//
//	for _, day := range days {
//		if day == nil {
//			continue
//		}
//		log.Println(day)
//
//		_, err :=begin.Stmt(stmt).Exec(day.Code, day.Date,
//			day.Open, day.Close, day.High, day.Low,
//			day.Volume, day.Money)
//
//
//		if err != nil {
//			panic(err.Error())
//		}
//	}
//
//	err = begin.Commit()
//	if err != nil {
//		panic(err.Error())
//	}
//
//	return nil
//}

func GetTradeHis(stock *model.Stock, begin string, end string) []*model.DayTrade {
	sql := `
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
	AND date >= ?
	AND date <= ?
	ORDER BY DATE ASC
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	days := make([]*model.DayTrade, 0)
	rows, err := stmt.Query(stock.Code, begin, end)
	defer rows.Close()
	if rows == nil || err != nil {
		return days
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
			continue
		}
		days = append(days, day)
	}
	return days
}

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

	return days
}

/**
 * 某只股票最新一天的交易数据
 */
func GetLatestDayStock(stock *model.Stock) *model.DayTrade {
	sql := `
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
	LIMIT 1
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	day := &model.DayTrade{}
	row := stmt.QueryRow(stock.Code)
	if row == nil {
		return day
	}
	err = row.Scan(
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
		return day
	}
	return day
}

func GetAllStocks() []*model.Stock {
	sql := `
		SELECT code, cn_name FROM stock
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	rows, _ := stmt.Query()
	if rows == nil {
		return nil
	}
	stocks := make([]*model.Stock, 0)
	for rows.Next() {
		stock := &model.Stock{}
		err = rows.Scan(&stock.Code, &stock.CnName)
		if err != nil {
			return nil
		}
		stock.Type = model.Code2Type(stock.Code)
		stocks = append(stocks, stock)
	}
	return stocks
}

func GetStockInfo(stock *model.Stock) bool {
	sql := `
		SELECT cn_name FROM stock WHERE code = ?
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	row := stmt.QueryRow(stock.Code)
	if row == nil {
		return false
	}
	err = row.Scan(&stock.CnName)
	if err != nil {
		return false
	}
	stock.Type = model.Code2Type(stock.Code)
	return true

}

// 更新股票实体信息
func UpSertStockInfo(stocks []*model.Stock) error {
	sql := `
	INSERT INTO stock
	(
		code,
		cn_name
	)
	VALUES
	(
		?,
		?
	)
	ON DUPLICATE KEY UPDATE
	cn_name = ?
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	for _, stock := range stocks {
		if stock == nil {
			continue
		}
		_, err = stmt.Exec(stock.Code, stock.CnName, stock.CnName)
		if err != nil {
			panic(err.Error())
		}
	}
	return nil
}



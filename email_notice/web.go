package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
	"strconv"
	"net/http"
	. "email_notice/common"
)

func AddNoticeStock(code string, heightPrice, lowPrice float64) {
	mx.Lock()
	var s *Stock
	var isNew = false
	if oldStock, ok := NoticeStockS[code]; ok {
		s = oldStock
	} else {
		s = &Stock{
			Code:           code,
			NoticeCallBack: NoticeEmail,
			NoticeLimit:    2,
		}
		NoticeStockS[code] = s
		isNew = true
	}

	s.Url = fmt.Sprintf(mainUrl, code)
	s.LowPrice = lowPrice
	s.HeightPrice = heightPrice
	s.Count = 0
	mx.Unlock()
	if isNew {
		s.Start()
	}
}

func DeleteNoticeStock(code string) bool {
	if old, ok := NoticeStockS[code]; ok {
		mx.Lock()
		old.Close()
		delete(NoticeStockS, code)
		mx.Unlock()
		return true
	} else {
		return false
	}
}

//localhost:4661/delete?code=sh511990
func deleteNoticeStock(c echo.Context) error {
	code := c.QueryParam("code")

	isDelete := DeleteNoticeStock(code)

	if isDelete {
		return c.String(http.StatusOK, "Delete OK")
	} else {
		return c.String(http.StatusOK, "Delete Object Not Exist")
	}
}

//120.79.154.53:4661/add?code=sh511990&height=100.05&low=99.905
//netstat -aon|findstr "40051"
//localhost:4661/add?code=sh511990&height=100.05&low=99
//localhost:4661/add?code=sh511990&height=100.05&low=99.905
func addNoticeStock(c echo.Context) error {
	code := c.QueryParam("code")
	heightPrice, _ := strconv.ParseFloat(c.QueryParam("height"), 64)
	lowPrice, _ := strconv.ParseFloat(c.QueryParam("low"), 64)

	AddNoticeStock(code, heightPrice, lowPrice)

	return c.String(http.StatusOK, "Add OK")
}

//localhost:4661/addstock?code=sh600004&money=130000&copies=20&price=14.844
func addStock(c echo.Context) error {
	code := c.QueryParam("code")
	price, _ := strconv.ParseFloat(c.QueryParam("price"), 64)
	money, _ := strconv.ParseFloat(c.QueryParam("money"), 64)
	copies, _ := strconv.Atoi(c.QueryParam("copies"))

	bmx.Lock()
	var s *BuyStock
	var isNew = false
	if oldStock, ok := BuyStocks[code]; ok {
		s = oldStock
	} else {
		s = &BuyStock{
			StockName:                  code,
			BuyPrice:                   price,
			AllMoney:                   money,
			NumberOfCopies:             copies,
			NumberOfCopiesPice:         make(map[float64]int, 0),
			OrderNumberOfCopiesPiceKey: make([]float64, 0),
			AddNoticeFunc:              AddNoticeStock,
			DeleteNoticeFunc:           DeleteNoticeStock,
		}
		BuyStocks[code] = s
		isNew = true
	}

	s.StockUrl = fmt.Sprintf(mainUrl, code)

	bmx.Unlock()
	if isNew {
		s.Start()
	}

	return c.String(http.StatusOK, "Add Stock OK")
}

//localhost:4661/deletestock?code=sh600004
func deleteStock(c echo.Context) error {
	code := c.QueryParam("code")

	if old, ok := BuyStocks[code]; ok {
		mx.Lock()
		old.Close()
		delete(BuyStocks, code)
		mx.Unlock()
		return c.String(http.StatusOK, "Delete OK")
	} else {
		return c.String(http.StatusOK, "Delete Object Not Exist")
	}

}

func RunHttpServer() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("1M"))

	e.Static("/", "./assets")

	e.GET("/add", addNoticeStock)
	e.GET("/delete", deleteNoticeStock)
	e.GET("/addstock", addStock)
	e.GET("/deletestock", deleteStock)
	fmt.Println("RunHttpServer ----------------- ")
	err := e.Start(httpPort)
	fmt.Println(err)
	panic(err.Error())

}

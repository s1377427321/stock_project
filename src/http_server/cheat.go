package http_server

import (
	"constant"
	//	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	//	"sort"
	//"storage"
	"strconv"
	"tactics"
	"tactics/percentTen"

)

type H map[string]interface{}

func RunHttpServer() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("1M"))

	e.Static("/", "./assets")

	e.GET("/test", test)
	e.GET("/tactics1", tactics1)
	e.GET("/beforeDay", beforeDay)
	e.GET("/beforeAllDay", beforeAllDay)
	e.Start(constant.HTTP_PORT)
	fmt.Println("RunHttpServer -----------------")
}

func beforeDay(c echo.Context) error {
	c.Response().CloseNotify()
	//code ,_:= strconv.Atoi(c.QueryParam("code"))
	//
	//beforeDay ,_:=strconv.Atoi(c.QueryParam("beforeDay"))

	//menoy,_:=strconv.ParseFloat(c.QueryParam("momey"), 64)

	code, _ := strconv.Atoi("002923") //0409

	beforeDay, _ := strconv.Atoi("100")

	menoy, _ := strconv.ParseFloat("200000", 64)

	percentTen.Start(code, beforeDay, menoy)
	//beforeStruct:= percentTen.NewBeforeDayStruct(code,beforeDay,menoy)
	//beforeStruct.Do()

	return nil
}

func beforeAllDay(c echo.Context) error {
	c.Response().CloseNotify()


	//beforeDay ,_:=strconv.Atoi(c.QueryParam("beforeDay"))
	//menoy,_:=strconv.ParseFloat(c.QueryParam("momey"), 64)
	beforeDay, _ := strconv.Atoi("30")
	menoy, _ := strconv.ParseFloat("200000", 64)

	percentTen.BeforeAllDay(beforeDay,menoy)

	return nil
}



func tactics1(c echo.Context) error {
	c.Response().CloseNotify()
	code, _ := strconv.Atoi(c.QueryParam("code"))
	origPrice, _ := strconv.ParseFloat(c.QueryParam("origPrice"), 64)
	bearLose, _ := strconv.ParseFloat(c.QueryParam("bearLose"), 64)
	haveMoney, _ := strconv.ParseFloat(c.QueryParam("haveMoney"), 64)
	divide, _ := strconv.ParseFloat(c.QueryParam("divide"), 64)

	fmt.Println(code, origPrice, bearLose, haveMoney, divide)

	tac := tactics.NewTactics1(code, origPrice, bearLose, haveMoney, divide)
	tac.Do()
	//fmt.Println(tac.Grids)
	//	result := make(map[string]string)

	//	sortTmep := make([]float64, 0)

	//	for key, _ := range tac.Grids {
	//		sortTmep = append(sortTmep, key)
	//	}

	//	sort.Float64s(sortTmep)

	//	for _, val := range sortTmep {
	//		t := strconv.FormatFloat(val, 'f', 2, 64)
	//		result[t] = tac.Grids[val]
	//	}

	//	fmt.Println(result)

	//	result := make(map[string]string)
	//return c.JSON(http.StatusOK, result)
	//return c.JSON(http.StatusOK, H{"id": tac.ToString()})

	return c.JSON(http.StatusOK, tac.Grids)

}

func test(c echo.Context) error {
	c.Response().CloseNotify()

	fmt.Println("ehco jjjjjjjjjjjjjjjjjjjjjjj")
	return nil
}

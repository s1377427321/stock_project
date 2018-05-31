package http_server

import (
	"constant"
	//	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	//	"sort"
	"strconv"
	"tactics"
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
	e.Start(constant.HTTP_PORT)
	fmt.Println("RunHttpServer -----------------")
}

func tactics1(c echo.Context) error {
	c.Response().CloseNotify()
	code := c.QueryParam("code")
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

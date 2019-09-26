package main

import (
	"bufio"
	"golang.org/x/net/context"
	"strconv"
	"os"
	"path/filepath"
	"fmt"
	"strings"
	"github.com/go-vgo/robotgo"
	"time"
)

var reader *bufio.Reader
var trainStock *TrainStock
var ctx context.Context
var cancel context.CancelFunc

func main() {
	//stopCh = make(chan int, 0)
	reader = bufio.NewReader(os.Stdin)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Errorf("errror", err)
		return
	}
	dir = strings.Replace(dir, "\\", "/", -1)
	for {
		fmt.Println("1 恢复  2 重新开始")
		if strBytes, _, err := reader.ReadLine(); err == nil {
			if strings.Contains(string(strBytes), "2") == true {

				fmt.Println("请输入起始数据：OriginMoney,buyCode,buyMoney,buyDay")
				if strBytes, _, err := reader.ReadLine(); err == nil {
					buyData := strings.Split(string(strBytes), ",")
					trainStock = NewTrain(buyData)
					ctx, cancel = context.WithCancel(context.TODO())
					//fmt.Println("1 keep  2 oldbuy 3rangebuy 4sell 5退出")
					Do()

				} else {
					return
				}
			} else {
				trainStock = LoadSaveFileRecover()
				ctx, cancel = context.WithCancel(context.TODO())
				Do()
			}
		}
	}
}

func Do() {
	//keep
	keepCallback := func() {

		for true {
			select {
			case <-ctx.Done():
				return
			default:
			}
			time.Sleep(200 * time.Millisecond)
			//fmt.Println("AAAAAAAAAAAA")
			keve := robotgo.AddEvent("space")
			if keve {
				//fmt.Println("you press... ", "space")
				trainStock.NextDay(KEEP)
			}
		}
	}

	go keepCallback()

	for {
		fmt.Println("1 keep  2 oldbuy 3rangebuy 4sell 5退出 ")
		if strBytes, _, err := reader.ReadLine(); err == nil {
			if strings.Contains(string(strBytes), "1") == true {
				trainStock.NextDay(KEEP)
			} else if strings.Contains(string(strBytes), "2") == true {
				fmt.Println("BuyDay:")
				if strBytes, _, err := reader.ReadLine(); err == nil {
					isOk := trainStock.OldBuy(string(strBytes))
					if isOk {
						go keepCallback()
					} else {
						fmt.Println("重新买入失败")
					}
				}
			} else if strings.Contains(string(strBytes), "3") == true {
				fmt.Println("endDay:")
				if strBytes, _, err := reader.ReadLine(); err == nil {

					isOk := trainStock.KeepRange(string(strBytes))
					if isOk {
						go keepCallback()
					} else {
						fmt.Println("重新买入范围失败")
					}
				}
			} else if strings.Contains(string(strBytes), "4") == true {
				trainStock.NextDay(SELL)
				cancel()
			} else if strings.Contains(string(strBytes), "5") == true {
				break
			}
		}
	}

	fmt.Println("All OVER")
}

func NewTrain(data []string) (*TrainStock) {
	my, _ := strconv.ParseFloat(data[0], 64)
	bc := data[1]
	buym, _ := strconv.ParseFloat(data[2], 64)
	bd := data[3]
	opt := &BuyOpt{
		BuyMoney: buym,
		BuyDay:   bd,
		BuyCode:  bc,
	}

	ts := &TrainStock{}
	ts.NewBuy(my, opt)
	return ts
}

//func main() {
//	keve := robotgo.AddEvent("alt")
//	if keve {
//		fmt.Println("you press... ", "alt")
//	}
//}

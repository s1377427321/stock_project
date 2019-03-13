package main

import "fmt"

func main()  {
	lv:=0.15
	var result float64 = 150000

	for i:=1;i<50 ;i++  {
		result = result*(1+lv)
		fmt.Println(fmt.Sprintf("%d  ---  %f",i,result))
	}
}
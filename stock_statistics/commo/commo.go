package commo

import (
	"fmt"
	"strconv"
	"sort"
)

func Round64(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}


func GetMinMax(a,b,c float64) (aret float64 , bret float64) {
	temp:=[]float64{a,b,c}

	sort.Float64s(temp)

	return temp[0],temp[len(temp)-1]
}
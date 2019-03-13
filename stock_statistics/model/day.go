package model

import (
	"fmt"
	"strconv"
)

type DayTrade struct {
	Code   int
	Date   string
	Open   float32
	Close  float32
	Volume int
	Money  int
	High   float32
	Low    float32
}

func (this *DayTrade) String() string {
	return fmt.Sprintf("%s [%d] 开:%f 收:%f 高:%f 低:%f 量:%d 金:%d\n", this.Date, this.Code,
		this.Open, this.Close, this.High, this.Low, this.Volume, this.Money)
}

type DayTradeSlice []*DayTrade

func (c DayTradeSlice) Len() int {
	return len(c)
}
func (c DayTradeSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c DayTradeSlice) Less(i, j int) bool {
	ii:=c[i]
	jj:=c[j]

	cmpi,_:=strconv.Atoi(ii.Date)

	cmpj,_:=strconv.Atoi(jj.Date)

	return  cmpi > cmpj
}






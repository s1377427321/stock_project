package model

import "fmt"

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

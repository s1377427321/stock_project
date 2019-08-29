package cook

import (
	"common"
	"encoding/json"
	"io/ioutil"
	"github.com/astaxie/beego"
	"math/rand"
	"strings"
	"fmt"
)

type CookItem struct {
	Maincourses []string
	Vegetables  [] string
	Soups       []string
}

func (c *CookItem) InitCookItem() {
	c.Maincourses = make([]string, 0)
	c.Vegetables = make([]string, 0)
	c.Soups = make([]string, 0)
}

type RandomCook struct {
	SelectCooks    *CookItem
	NotSelectCooks *CookItem
	CookBook       *CookItem
}

func (r *RandomCook) SetCookBook(name string) {
	fpath := common.GetApplicationPosition()
	fpath += "/" + name

	readBinary, err := ioutil.ReadFile(fpath)
	if err != nil {
		panic(err)
	}

	r.CookBook.InitCookItem()

	json.Unmarshal(readBinary, r.CookBook)

	mtemp := make([]string, 0)
	for _, v := range r.CookBook.Maincourses {
		if strings.Contains(v, "no") {
			continue
		}
		mtemp = append(mtemp, v)
	}
	r.CookBook.Maincourses = mtemp

	vtemp := make([]string, 0)
	for _, v := range r.CookBook.Vegetables {
		if strings.Contains(v, "no") {
			continue
		}
		vtemp = append(vtemp, v)
	}
	r.CookBook.Vegetables = vtemp

	stemp := make([]string, 0)
	for _, v := range r.CookBook.Soups {
		if strings.Contains(v, "no") {
			continue
		}
		stemp = append(stemp, v)
	}
	r.CookBook.Soups = stemp

	beego.Info(r.CookBook)
}

func (r *RandomCook) DeepCopy(target *[]string, source *[]string) {
	for _, v := range *source {
		*target = append(*target, v)
	}
}

func (r *RandomCook) RandomCooks(mNums, vNums, sNums int) string {
	var ms, vs, ss string
	for i := 0; i < mNums;{
		mlen := len(r.NotSelectCooks.Maincourses)
		if mlen > 0 {
			rIndex := rand.Intn(mlen)
			ms += "\t" + r.NotSelectCooks.Maincourses[rIndex]
			r.NotSelectCooks.Maincourses = append(r.NotSelectCooks.Maincourses[0:rIndex], r.NotSelectCooks.Maincourses[rIndex+1:]...)
			i++
		} else {
			r.DeepCopy(&r.NotSelectCooks.Maincourses, &r.CookBook.Maincourses)
		}
	}

	for i := 0; i < vNums;{
		mlen := len(r.NotSelectCooks.Vegetables)
		if mlen > 0 {
			rIndex := rand.Intn(mlen)
			vs += "\t" + r.NotSelectCooks.Vegetables[rIndex]
			r.NotSelectCooks.Vegetables = append(r.NotSelectCooks.Vegetables[0:rIndex], r.NotSelectCooks.Vegetables[rIndex+1:]...)
			i++
		} else {
			r.DeepCopy(&r.NotSelectCooks.Vegetables, &r.CookBook.Vegetables)
		}
	}

	for i := 0; i < sNums;{
		mlen := len(r.NotSelectCooks.Soups)
		if mlen > 0 {
			rIndex := rand.Intn(mlen)
			ss += "\t" + r.NotSelectCooks.Soups[rIndex]
			r.NotSelectCooks.Soups = append(r.NotSelectCooks.Soups[0:rIndex], r.NotSelectCooks.Soups[rIndex+1:]...)
			i++
		} else {
			r.DeepCopy(&r.NotSelectCooks.Soups, &r.CookBook.Soups)
		}
	}

	des:=fmt.Sprintf("===主菜：%v,===青菜：%v,===汤：%v", ms, vs, ss)

	beego.Info(des)

	return des

}

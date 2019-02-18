package main

import (
	"time"
	"strconv"
)

type LFUNode struct {
	value       int
	useCount    int
	lastGetTime time.Time
}

type LFUCach struct {
	cache     map[string]*LFUNode
	capacity int
	used     int
}

func (this *LFUCach) get(k int) {

}

func (this *LFUCach) set(k int, v int) {
	m,ok:=this.cache[strconv.Itoa(k)]
	if ok {
		m.useCount ++
		m.lastGetTime = time.Now()
		m.value = v
		return
	}else {
		ln:=&LFUNode{}
		ln.value = v
		ln.lastGetTime = time.Now()
		ln.useCount = 0
		if this.capacity == 0 {
			return
		}
		if this.used < this.capacity  {
			this.used ++
		}else {
			this.removeLast()
		}
		this.cache[strconv.Itoa(k)] = ln
	}
}

func (this *LFUCach) removeLast() {
	var minCount  = 9999999
	var getTime time.Time = time.Now()
	var rmS string

	keyS:=make([]string,0)
	for k,_:=range this.cache {
		keyS =  append(keyS,k)
	}

	for i:=0;i<len(keyS) ;i++  {
		v:=this.cache[keyS[i]]
		if v.useCount < minCount || (v.useCount == minCount && v.lastGetTime.Before(getTime)) {
			minCount = v.useCount
			getTime = v.lastGetTime
			rmS = keyS[i]
		}

	}

	delete(this.cache,rmS)
}

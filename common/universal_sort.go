package common

import "sort"

type UniversalSort struct {
	Bodys []interface{}
	do    func(p, q *interface{}) bool
}

type SortBodyDo func(p, q *interface{}) bool

func (s UniversalSort) Len() int {
	return len(s.Bodys)
}

func (s UniversalSort) Less(i, j int) bool {
	return s.do(&s.Bodys[i], &s.Bodys[j])
}

func (s UniversalSort) Swap(i, j int) {
	s.Bodys[i], s.Bodys[j] = s.Bodys[j], s.Bodys[i]
}

func SortBody(bodys []interface{}, do SortBodyDo) {
	sort.Sort(UniversalSort{bodys, do})
}

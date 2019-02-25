package main

func (this *Solution30) insert(intervals []*Interval30, interval *Interval30) []*Interval30 {
	ret := make([]*Interval30, 0)
	var i = 0
	size := len(intervals)
	for ; i < size && intervals[i].End < interval.Start; i++ {
		ret = append(ret, intervals[i])
	}

	if i == size {
		ret = append(ret, interval)
		return ret
	}

	start := this.min(intervals[i].Start, interval.Start)

	for ; i < size && intervals[i].Start <= interval.End; {
		i++
	}

	var end int
	if i > 0 {
		end = this.max(intervals[i-1].End, interval.End)
	} else {
		end = interval.End
	}
	newI := &Interval30{Start: start, End: end}
	ret = append(ret, newI)

	for ; i < size; i++ {
		ret = append(ret, newI)
	}

	return ret
}

func (this *Solution30) min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}

func (this *Solution30) max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

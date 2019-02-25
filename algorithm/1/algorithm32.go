package main

//https://www.cnblogs.com/Tang-tangt/p/9038033.html

type algorithm32 struct {
}

/**
 * @param source : A string
 * @param target: A string
 * @return: A string denote the minimum window, return "" if there is no such a string
 */
func (this *algorithm32) minWindow(source string, target string) string {
	// write your code here
	s := []byte(source)
	t := []byte(target)
	ssize := len(s)
	tsize := len(t)

	tcount:=make(map[int]int,128)
	scount:=make(map[int]int,128)

	for i := 0; i < 128; i++ {
		tcount[i] = 0
		scount[i] = 0
	}

	for i:=0;i<tsize ;i++  {
		index:= int(t[i])
		tcount[index]++
	}

	begin:= -1
	found:=0
	minLen:=ssize
	start:=0
	for i:=0 ;i<ssize;i++ {
		scount[int(s[i])]++
		if scount[int(s[i])] <= tcount[int(s[i])]{
			found++
		}

		if found == tsize {
			for ;start<i && scount[int(s[start])]>tcount[int(s[start])]; {
				scount[int(s[start])]--
				start++
			}

			if i-start<minLen {
				minLen = i-start
				begin = start
			}

			scount[int(s[start])]--
			found --
			start ++
		}
	}

	if begin == -1 {
		return ""
	}else {
		return string(s[begin:minLen+1])
	}
}

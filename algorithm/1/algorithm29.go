package main

import "fmt"

type IsInterleaveS struct {

}

func (this *IsInterleaveS) IsInterleave(s1, s2, s3 string) bool {
	a := []byte(s1)
	b := []byte(s2)
	c := []byte(s3)

	rowLen := len(a)
	colLen := len(b)
	if (rowLen + colLen) != len(c) {
		return false
	}
	dp := make([][]bool, 0)
	for i := 0; i < rowLen; i++ {
		r := make([]bool, colLen)
		dp = append(dp, r)
	}

	for i:=0;i<rowLen ;i++  {
		for j:=0;j<colLen;j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			}else if i==0 {
				dp[i][j] = dp[i][j-1]&&b[j-1] ==c[i+j-1]
			}else if j == 0 {
				dp[i][j] = dp[i-1][j]&&a[i-1] == c[i+j-1]
			}else {
				dp[i][j] = (dp[i-1][j]&&a[i-1]==c[i+j-1]) || (dp[i][j-1]&&b[j-1] ==c[i+j-1])
			}
		}
	}

	return dp[rowLen-1][colLen-1]
}


func main() {
	t:=&IsInterleaveS{}
	r:=t.IsInterleave("aabcc","dbbca","aadbbbaccc")
	fmt.Println(r)
}

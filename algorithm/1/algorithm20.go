package main

import (
	"math"
	"fmt"
)

/*
20. 骰子求和
中文
English
扔 n 个骰子，向上面的数字之和为 S。给定 Given n，请列出所有可能的 S 值及其相应的概率。

样例
给定 n = 1，返回 [ [1, 0.17], [2, 0.17], [3, 0.17], [4, 0.17], [5, 0.17], [6, 0.17]]。

注意事项
你不需要关心结果的准确性，我们会帮你输出结果。

 */

var resultList = make([]map[int]float64, 0)

func dicesSum(n int) {
	dp := make([][]float64, n+1)
	po := math.Pow(6, float64(n))
	for i := 0; i <= n; i++ {
		r := make([]float64, 6*n+1)
		dp[i]=r
	}

	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := i; j <= 6*n; j++ {
			dp[i][j] = 0
			var k int
			if j <= 6 {
				k = 1
			} else {
				k = j - 6
			}
			for ; k < j; k++ {
				dp[i][j] += dp[i-1][k]
			}
		}
	}

	for i := n; i <= 6*n; i++ {
		result:=dp[n][i]/po
		rm:=make(map[int]float64,0)
		rm[i] = result
		resultList = append(resultList,rm)
	}
}

func main() {
	dicesSum(3)
	fmt.Println(resultList)
}

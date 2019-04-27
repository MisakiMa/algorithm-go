package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("vim-go")
	var N, W int
	fmt.Scan(&N, &W)
	w := make([]int, N+1)
	v := make([]int, N+1)

	for i := 0; i < N; i++ {
		fmt.Scan(&w[i], &v[i])
	}

	// 二次元スライスを作る
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}

	// dp
	for i := 0; i < N; i++ {
		for j := 0; j < W+1; j++ {
			if j >= w[i] {
				dp[i+1][j] = max(dp[i][j-w[i]]+v[i], dp[i][j])
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}

	// ans
	fmt.Println("ans", dp[N][W])

	fmt.Println("→ W")
	fmt.Println("↓ N")
	// dp
	fmt.Println("dp")
	for i := 0; i < N+1; i++ {
		// fmt.Println(dp[i])
		fmt.Println("\n")
		for j := 0; j < W+1; j++ {
			fmt.Printf("%3d", dp[i][j])
		}
	}
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

/*
参考にしました
https://atcoder.jp/contests/dp/tasks/dp_d
https://atcoder.jp/contests/dp/submissions/5112197
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var minv int
	fmt.Scanf("%d", &minv)
	maxd := -1 << 30
	for i := 1; i < n; i++ {
		var v int
		fmt.Scanf("%d", &v)
		maxd = max(maxd, v-minv)
		minv = min(minv, v)
	}
	fmt.Println(maxd)
}

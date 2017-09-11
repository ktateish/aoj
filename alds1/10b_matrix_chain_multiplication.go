package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func solve(a []int) int {
	n := len(a) - 1
	dp := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		dp[i] = make([]int, len(a))
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 0
	}
	for k := 1; k < n; k++ {
		for i := 0; i+k < n; i++ {
			j := i + k
			dp[i][j] = 1 << 31
			for l := i + 1; l <= j; l++ {
				dp[i][j] = min(
					dp[i][j],
					dp[i][l-1]+dp[l][j]+a[i]*a[l]*a[j+1],
				)
			}
		}
	}
	return dp[0][n-1]
}

func main() {
	defer stdout.Flush()
	n := readInt()
	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		a[i] = readInt()
		a[i+1] = readInt()
	}
	println(solve(a))
}

var (
	readString func() string
	stdout     *bufio.Writer
)

func init() {
	readString = newReadString(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}

func newReadString(ior io.Reader) func() string {
	r := bufio.NewScanner(ior)
	r.Split(bufio.ScanWords)

	return func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
}

func readInt() int {
	i, err := strconv.Atoi(readString())
	if err != nil {
		panic(err.Error())
	}
	return i
}

func readIntSlice(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = readInt()
	}
	return b
}

func readLengthAndSlice() []int {
	n := readInt()
	return readIntSlice(n)
}

func printf(f string, args ...interface{}) (int, error) {
	return fmt.Fprintf(stdout, f, args...)
}

func println(args ...interface{}) (int, error) {
	return fmt.Fprintln(stdout, args...)
}

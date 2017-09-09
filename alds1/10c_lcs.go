package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func lcs(x, y string) int {
	rx := []rune{}
	for _, r := range x {
		rx = append(rx, r)
	}
	ry := []rune{}
	for _, r := range y {
		ry = append(ry, r)
	}
	dp := make([][]int, len(rx)+1)
	for i := 0; i < len(rx)+1; i++ {
		dp[i] = make([]int, len(ry)+1)
	}
	for i := 1; i <= len(rx); i++ {
		for j := 1; j <= len(ry); j++ {
			v := dp[i-1][j-1]
			if rx[i-1] == ry[j-1] {
				v++
			}
			if dp[i-1][j] > v {
				v = dp[i-1][j]
			}
			if dp[i][j-1] > v {
				v = dp[i][j-1]
			}
			dp[i][j] = v
		}
	}
	return dp[len(rx)][len(ry)]
}

func main() {
	defer stdout.Flush()
	q := readInt()
	for i := 0; i < q; i++ {
		x := readString()
		y := readString()
		println(lcs(x, y))
	}
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

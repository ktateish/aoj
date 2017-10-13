package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var INF int64 = math.MaxInt64

func min(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	defer stdout.Flush()
	n := readInt()
	g := make([][]int64, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int64, n)
		for j := 0; j < n; j++ {
			g[i][j] = INF
		}
		g[i][i] = 0
	}
	q := readInt()
	for i := 0; i < q; i++ {
		s := readInt()
		t := readInt()
		g[s][t] = int64(readInt())
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if g[i][k] == INF {
				continue
			}
			for j := 0; j < n; j++ {
				if g[k][j] == INF {
					continue
				}
				g[i][j] = min(g[i][j], g[i][k]+g[k][j])
			}
		}
	}
	for i := 0; i < n; i++ {
		if g[i][i] < 0 {
			println("NEGATIVE CYCLE")
			return
		}
	}
	for i := 0; i < n; i++ {
		sp := ""
		for j := 0; j < n; j++ {
			c := fmt.Sprintf("%d", g[i][j])
			if g[i][j] == INF {
				c = "INF"
			}
			printf("%s%s", sp, c)
			sp = " "
		}
		println()
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

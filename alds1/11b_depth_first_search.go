package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func dfs(g, o [][]int, v, t int) int {
	if len(o[v]) != 0 {
		return t
	}
	t++
	o[v] = append(o[v], t)
	for _, u := range g[v] {
		t = dfs(g, o, u, t)
	}
	t++
	o[v] = append(o[v], t)
	return t
}

func main() {
	defer stdout.Flush()
	n := readInt()
	g := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		u := readInt()
		k := readInt()
		for j := 0; j < k; j++ {
			g[u] = append(g[u], readInt())
		}
	}
	o := make([][]int, n+1)
	t := 0
	for i := 1; i <= n; i++ {
		t = dfs(g, o, i, t)
	}
	for i := 1; i <= n; i++ {
		printf("%d", i)
		for _, t := range o[i] {
			printf(" %d", t)
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

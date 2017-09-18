package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func mst(a [][]int, root int) []int {
	n := len(a)
	visited := make([]bool, n)
	d := make([]int, n) // cost from current MST
	p := make([]int, n) // parents of vertex i in the MST
	for i := 0; i < n; i++ {
		d[i] = math.MaxInt32
		p[i] = -1
	}

	d[root] = 0 // root vertex is in the MST from the begining
	for {
		mincost := math.MaxInt32
		u := -1
		// choose the vertex of minimum cost
		for i := 0; i < n; i++ {
			if !visited[i] && d[i] < mincost {
				mincost = d[i]
				u = i
			}
		}

		if u == -1 {
			break
		}
		visited[u] = true

		// update cost of adjacent verteces
		for i := 0; i < n; i++ {
			if !visited[i] && a[u][i] < d[i] {
				d[i] = a[u][i]
				p[i] = u
			}
		}
	}
	return p
}

func main() {
	defer stdout.Flush()
	n := readInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			w := readInt()
			if w == -1 {
				w = math.MaxInt32
			}
			a[i][j] = w
		}
	}
	p := mst(a, 0) // root vetex is 0
	total := 0
	for i := 1; i < n; i++ {
		total += a[i][p[i]]
	}
	println(total)
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

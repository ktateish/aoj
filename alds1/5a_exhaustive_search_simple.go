package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func check(a []int, m int) bool {
	var rec func(int, int) bool
	rec = func(i, m int) bool {
		if i == len(a) {
			return m == 0
		}
		return rec(i+1, m-a[i]) || rec(i+1, m)
	}
	return rec(0, m)
}

func main() {
	n := readInt()
	a := readIntSlice(n)
	q := readInt()
	m := readIntSlice(q)

	for _, p := range m {
		if check(a, p) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

var (
	readInt      func() int
	readIntSlice func(int) []int
)

type FastScanner struct {
	r *bufio.Reader
}

func NewFastScanner(r io.Reader) *FastScanner {
	return &FastScanner{bufio.NewReader(r)}
}

func (fs *FastScanner) Scanf(f string, args ...interface{}) (int, error) {
	n, err := fmt.Fscanf(fs.r, f, args...)
	for {
		r, _, err := fs.r.ReadRune()
		if err != nil || !unicode.IsSpace(r) {
			fs.r.UnreadRune()
			break
		}
	}
	return n, err
}

func init() {
	fs := NewFastScanner(os.Stdin)
	readInt = func() int {
		var n int
		fs.Scanf("%d", &n)
		return n
	}

	readIntSlice = func(n int) []int {
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fs.Scanf("%d", &a[i])
		}
		return a
	}
}

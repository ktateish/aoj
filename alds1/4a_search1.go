package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func Lsearch(a []int, x int) int {
	for _, i := range a {
		if i == x {
			return 1
		}
	}
	return 0
}

func main() {
	n := readInt()
	S := readIntSlice(n)
	q := readInt()
	T := readIntSlice(q)

	total := 0
	for _, i := range T {
		total += Lsearch(S, i)
	}

	fmt.Println(total)
}

// vim:set ft=go:

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

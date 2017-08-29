package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func partition(a []int, p, r int) int {
	i := p - 1
	for j := p; j < r; j++ {
		if a[j] <= a[r] {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	n, a := readIntAndSlice()
	q := partition(a, 0, n-1)
	sp := ""
	for i := 0; i < n; i++ {
		if i == q {
			fmt.Fprintf(w, "%s[%d]", sp, a[i])
		} else {
			fmt.Fprintf(w, "%s%d", sp, a[i])
		}
		sp = " "
	}
	fmt.Fprintln(w)
	w.Flush()
}

var (
	readInt         func() int
	readIntSlice    func(int) []int
	readIntAndSlice func() (int, []int)
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

	readIntAndSlice = func() (int, []int) {
		n := readInt()
		return n, readIntSlice(n)
	}
}

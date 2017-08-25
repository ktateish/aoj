package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func loadable(w []int, k, p int) bool {
	l := 0
	q := p
	for i := 0; i < len(w); i++ {
		if q+w[i] > p {
			l++
			q = 0
			if l > k {
				// p is too small
				return false
			}
		}
		q += w[i]
	}
	return true
}

func search(w []int, k, minp, maxp int) int {
	if minp >= maxp {
		return maxp
	}
	p := (minp + maxp) / 2
	if loadable(w, k, p) {
		return search(w, k, minp, p)
	} else {
		return search(w, k, p+1, maxp)
	}
}

func main() {
	n := readInt()
	k := readInt()
	w := make([]int, n)
	maxw := 0
	wsum := 0
	for i := 0; i < n; i++ {
		w[i] = readInt()
		if w[i] > maxw {
			maxw = w[i]
		}
		wsum += w[i]
	}

	fmt.Println(search(w, k, maxw, wsum))
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

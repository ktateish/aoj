package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func mergeSort(a []int) int {
	count := 0
	tmp := make([]int, len(a))
	var rec func(int, int)
	rec = func(l, r int) {
		if l+1 >= r {
			return
		}
		m := (l + r) / 2
		rec(l, m)
		rec(m, r)
		i := l
		k := l
		j := m
		for k < r {
			count++
			if j >= r {
				tmp[k] = a[i]
				i++
			} else if i >= m {
				tmp[k] = a[j]
				j++
			} else if a[i] <= a[j] {
				tmp[k] = a[i]
				i++
			} else {
				tmp[k] = a[j]
				j++
			}
			k++
		}
		for i := l; i < r; i++ {
			a[i] = tmp[i]
		}
	}
	rec(0, len(a))
	return count
}

func main() {
	n := readInt()
	s := readIntSlice(n)
	c := mergeSort(s)
	sp := ""
	for _, x := range s {
		fmt.Printf("%s%d", sp, x)
		sp = " "
	}
	fmt.Println()
	fmt.Println(c)
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

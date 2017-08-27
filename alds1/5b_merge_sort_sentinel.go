package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"unicode"
)

func mergeSort(a []int) int {
	ltmp := make([]int, len(a)/2+2)
	rtmp := make([]int, len(a)/2+2)
	var merge func(int, int, int) int
	var rec func(int, int) int

	merge = func(l, m, r int) int {
		for i := 0; i < m-l; i++ {
			ltmp[i] = a[l+i]
		}
		for i := 0; i < r-m; i++ {
			rtmp[i] = a[m+i]
		}
		ltmp[m-l] = math.MaxInt32
		rtmp[r-m] = math.MaxInt32
		i := 0
		j := 0
		count := 0
		for k := l; k < r; k++ {
			count++
			if ltmp[i] <= rtmp[j] {
				a[k] = ltmp[i]
				i++
			} else {
				a[k] = rtmp[j]
				j++
			}
		}
		return count
	}
	rec = func(l, r int) int {
		if l+1 >= r {
			return 0
		}
		m := (l + r) / 2
		count := rec(l, m)
		count += rec(m, r)
		count += merge(l, m, r)
		return count
	}
	return rec(0, len(a))
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

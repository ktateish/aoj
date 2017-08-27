package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func mergeSort(a []int) int {
	tmp := make([]int, len(a))
	merge := func(l, m, r int) int {
		count := 0
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
		return count
	}

	n := len(a) / 8
	var rec func(int, int) int
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
	var gorec func(int, int, chan int)
	gorec = func(l, r int, done chan int) {
		if l+1 >= r {
			done <- 0
			return
		}
		m := (l + r) / 2

		var count int
		if r-l < 128 || r-l < n {
			count += rec(l, m)
			count += rec(m, r)
		} else {
			donel := make(chan int)
			doner := make(chan int)
			go gorec(l, m, donel)
			go gorec(m, r, doner)
			count = <-donel
			close(donel)
			count += <-doner
			close(doner)
		}
		count += merge(l, m, r)
		done <- count
	}
	done := make(chan int)
	go gorec(0, len(a), done)
	count := <-done
	return count
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	n := readInt()
	s := readIntSlice(n)
	c := mergeSort(s)
	sp := ""
	for _, x := range s {
		fmt.Fprintf(w, "%s%d", sp, x)
		sp = " "
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, c)
	w.Flush()
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

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func countingSort(a, b []int, k int) {
	c := make([]int, k+1)
	n := len(a) - 1
	for i := 1; i <= n; i++ {
		c[a[i]]++
	}

	for i := 1; i <= k; i++ {
		c[i] = c[i] + c[i-1]
	}

	for i := n; i > 0; i-- {
		b[c[a[i]]] = a[i]
		c[a[i]]--
	}
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	a, k := readLengthAndSlice()
	a = append([]int{0}, a...)
	b := make([]int, len(a))
	countingSort(a, b, k)
	sp := ""
	for i := 1; i < len(a); i++ {
		fmt.Fprintf(w, "%s%d", sp, b[i])
		sp = " "
	}
	fmt.Fprintln(w)
}

var (
	Stdin *FastScanner
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

func readInt() int {
	var n int
	Stdin.Scanf("%d", &n)
	return n
}

func readIntSlice(n int) ([]int, int) {
	a := make([]int, n)
	k := 0
	for i := 0; i < n; i++ {
		Stdin.Scanf("%d", &a[i])
		if a[i] > k {
			k = a[i]
		}
	}
	return a, k
}

func readLengthAndSlice() ([]int, int) {
	n := readInt()
	return readIntSlice(n)
}

func init() {
	Stdin = NewFastScanner(os.Stdin)
}

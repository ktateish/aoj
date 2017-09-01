package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

type Element struct {
	Value int
	Index int
	Done  bool
}

type ElementSlice []Element

func (e ElementSlice) Len() int {
	return len(e)
}

func (e ElementSlice) Less(i, j int) bool {
	return e[i].Value < e[j].Value
}

func (e ElementSlice) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func calch(a []Element, s, val int) int {
	total := 0
	for i := a[s].Index; i != s; i = a[i].Index {
		total += val + a[i].Value
		a[i].Done = true
	}
	return total
}

func calc(a []Element) int {
	sort.Sort(ElementSlice(a))

	total := 0
	for i := 0; i < len(a); i++ {
		if a[i].Index != i && !a[i].Done {
			x := calch(a, i, a[i].Value)
			y := calch(a, i, a[0].Value) + 2*(a[i].Value+a[0].Value)
			if x < y {
				total += x
			} else {
				total += y
			}
		}
	}

	return total
}

func main() {
	defer stdout.Flush()
	n := readInt()
	a := make([]Element, n)
	for i := 0; i < len(a); i++ {
		scanf("%d", &a[i].Value)
		a[i].Index = i
	}
	println(calc(a))
}

var (
	stdin  *bufio.Reader
	stdout *bufio.Writer
)

func init() {
	stdin = bufio.NewReader(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}

func scanf(f string, args ...interface{}) (int, error) {
	n, err := fmt.Fscanf(stdin, f, args...)
	for {
		r, _, err := stdin.ReadRune()
		if err != nil || !unicode.IsSpace(r) {
			stdin.UnreadRune()
			break
		}
	}
	return n, err
}

func readInt() int {
	var n int
	scanf("%d", &n)
	return n
}

func readIntSlice(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d", &a[i])
	}
	return a
}

func readLengthAndSlice() []int {
	n := readInt()
	return readIntSlice(n)
}

func readString() string {
	var s string
	scanf("%s", &s)
	return s
}

func printf(f string, args ...interface{}) (int, error) {
	return fmt.Fprintf(stdout, f, args...)
}

func println(args ...interface{}) (int, error) {
	return fmt.Fprintln(stdout, args...)
}

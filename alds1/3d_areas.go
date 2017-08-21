package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	a []interface{}
}

func NewStack() *Stack {
	return &Stack{make([]interface{}, 0)}
}

func (s *Stack) Push(x interface{}) {
	s.a = append(s.a, x)
}

func (s *Stack) Pop() (interface{}, bool) {
	n := len(s.a)
	if n == 0 {
		return 0, false
	}
	res := s.a[n-1]
	s.a = s.a[:n-1]
	return res, true
}

func (s *Stack) Peek() (interface{}, bool) {
	n := len(s.a)
	if n == 0 {
		return nil, false
	}
	return s.a[n-1], true
}

func (s *Stack) Empty() bool {
	return len(s.a) == 0
}

func fatal(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

type Area struct {
	Position int
	Volume   int
}

func main() {
	r := bufio.NewScanner(os.Stdin)

	r.Scan()
	levels := r.Text()

	s := NewStack()
	push := func(x int) {
		s.Push(x)
	}
	pop := func() int {
		x, ok := s.Pop()
		if !ok {
			return -1
		}
		return x.(int)
	}

	areas := NewStack()
	apush := func(x Area) {
		areas.Push(x)
	}
	apop := func() (Area, bool) {
		x, ok := areas.Pop()
		if !ok {
			return Area{}, false
		}
		return x.(Area), true
	}

	for i, r := range levels {
		switch r {
		case '\\':
			push(i)
		case '/':
			j := pop()
			if j == -1 {
				break
			}
			x := i - j
			for {
				a, ok := apop()
				if !ok {
					break
				}
				if a.Position < j {
					apush(a)
					break
				}
				x += a.Volume
			}
			apush(Area{i, x})
		case '_':
			// do nothing
		default:
			fatal("unkown level symbol: %c", r)
		}
	}

	total := 0
	as := []Area{}
	for a, ok := apop(); ok; a, ok = apop() {
		total += a.Volume
		as = append(as, a)
	}

	fmt.Printf("%d\n%d", total, len(as))
	for i := len(as) - 1; i >= 0; i-- {
		fmt.Printf(" %d", as[i].Volume)
	}
	fmt.Println()

	return
}

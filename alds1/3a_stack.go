package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func fatal(s string) {
	fmt.Fprintln(os.Stderr, s)
	os.Exit(1)
}

func main() {
	r := bufio.NewScanner(os.Stdin)

	r.Scan()
	args := strings.Fields(r.Text())

	s := NewStack()
	push := func(x int) {
		s.Push(x)
	}
	pop := func() int {
		x, ok := s.Pop()
		if !ok {
			fatal("stack empty")
		}
		return x.(int)
	}

	for _, arg := range args {
		if arg == "+" {
			y := pop()
			x := pop()
			push(x + y)
		} else if arg == "-" {
			y := pop()
			x := pop()
			push(x - y)
		} else if arg == "*" {
			y := pop()
			x := pop()
			push(x * y)
		} else {
			x, _ := strconv.Atoi(arg)
			push(x)
		}
	}

	fmt.Println(pop())
}

package stringcompress443

import (
	"iter"
	"strconv"
)

type Pair struct {
	rep int
	val byte
}

func (p *Pair) SetRep(rep int) {
	p.rep = rep
}

func (p Pair) GetRep() int {
	return p.rep
}

func (p Pair) GetVal() byte {
	return p.val
}

type Stack struct {
	items []Pair
}

// This is complex iterator
func (s *Stack) Items() iter.Seq[Pair] {
	return func(yeild func(Pair) bool) {
		for _, v := range s.items {
			if !yeild(v) {
				return
			}
		}
	}
}

func (s *Stack) Pop() Pair {
	retPair := s.Peek()
	s.items = s.items[:len(s.items)-1]
	return retPair
}

func (s *Stack) Peek() Pair {
	return s.items[len(s.items)-1]
}

func (s *Stack) Push(p Pair) {
	s.items = append(s.items, p)
}

func (s Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func compress(chars []byte) int {
	// Can change the value in slice not append/remove
	stack := NewStack()

	for _, c := range chars {
		if stack.IsEmpty() {
			stack.Push(NewPair(c, 1))
		} else if stack.Peek().GetVal() == c {
			p := stack.Pop()
			p.SetRep(p.GetRep() + 1)
			stack.Push(p)
		} else {
			stack.Push(NewPair(c, 1))
		}
	}

	index := 0
	for p := range stack.Items() {
		chars[index] = p.GetVal()
		index++
		rep := p.GetRep()
		if rep == 1 {
			continue
		}
		repArr := strconv.Itoa(rep)
		for _, repChar := range repArr {
			chars[index] = byte(repChar)
			index++
		}
	}
	return index
}

func NewPair(c byte, rep int) Pair {
	return Pair{
		val: c,
		rep: rep,
	}
}

func NewStack() *Stack {
	return &Stack{
		items: make([]Pair, 0),
	}
}

package decodestring394

import (
	"iter"
	"unicode"
)

type StackElement struct {
	repetitions int
	ch          rune
	isChar      bool
}

func (s *StackElement) getRepetitions() int {
	return s.repetitions
}

func (s *StackElement) getChar() rune {
	return s.ch
}

type Stack struct {
	elements []StackElement
}

func (s *Stack) size() int {
	return len(s.elements)
}

func (s *Stack) pop() *StackElement {
	ret := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return &ret
}

func (s *Stack) push(element interface{}) {
	var stackElement *StackElement
	var topStackElement *StackElement
	if !s.isEmpty() {
		topStackElement = s.peek()
	}
	switch e := element.(type) {
	case rune:
		stackElement = &StackElement{
			ch:     e,
			isChar: true,
		}
	case int:
		if topStackElement != nil && !topStackElement.isChar {
			topStackElement.repetitions = topStackElement.repetitions*10 + e
		} else {
			stackElement = &StackElement{
				repetitions: e,
				isChar:      false,
			}
		}
	}
	if stackElement != nil {
		s.elements = append(s.elements, *stackElement)
	}
}

func (s *Stack) peek() *StackElement {
	return &s.elements[len(s.elements)-1]
}

func (s *Stack) isEmpty() bool {
	return len(s.elements) == 0
}

func NewStack() *Stack {
	stack := &Stack{
		elements: []StackElement{},
	}
	return stack
}

func (s *Stack) Items() iter.Seq[StackElement] {
	return func(yield func(StackElement) bool) {
		for _, sElement := range s.elements {
			if !yield(sElement) {
				return
			}
		}
	}
}
func decodeString(s string) string {
	elementStack := NewStack()

	for _, ch := range s {
		if unicode.IsDigit(ch) {
			elementStack.push(int(ch - '0'))
		} else if ch == ']' {
			var poppedElements []rune
			for elementStack.peek().getChar() != '[' {
				poppedElements = append(poppedElements, elementStack.pop().getChar())
			}
			elementStack.pop() // '['
			count := elementStack.pop().getRepetitions()
			for range count {
				for i := len(poppedElements) - 1; i >= 0; i-- {
					elementStack.push(poppedElements[i])
				}
			}
		} else {
			elementStack.push(ch)
		}
	}

	retElements := make([]rune, elementStack.size())
	index := 0
	for sItem := range elementStack.Items() {
		retElements[index] = sItem.ch
		index++
	}
	return string(retElements)
}

package stack

import "errors"

// Stack is the base data structure.
// It's internal implementation is empty interface.
type Stack struct {
	data []interface{}
}

// New returns an initialized Stack.
func New() *Stack {
	return new(Stack)
}

// Push a new value onto the stack.
func (s *Stack) Push(value interface{}) {
	s.data = append(s.data, value)
}

// Pop some data from this stack.
func (s *Stack) Pop() (interface{}, error) {
	length := len(s.data)
	if length > 0 {
		v := s.data[length-1]
		s.data = s.data[:length-1]
		return v, nil
	}
	return nil, errors.New("no more items in stack.")
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	return len(s.data)
}

// Empty returns if the stack is empty or not.
func (s *Stack) Empty() bool {
	if len(s.data) > 0 {
		return false
	}
	return true
}

// Stack with an implementation of ints.
type IntStack struct {
	s Stack
}

// New returns an initialized IntStack.
func NewInt() *IntStack {
	return new(IntStack)
}

// Pushes a new int value onto the stack.
func (is *IntStack) Push(value int) {
	is.s.Push(value)
}

// Pops some int data from this stack.
func (is *IntStack) Pop() (int, error) {
	val, err := is.s.Pop()
	if err != nil {
		return 0, err
	}
	return val.(int), nil
}

// Size returns the number of elements in the stack.
func (is *IntStack) Size() int {
	return is.s.Size()
}

// Empty returns if the Stack is empty or not.
func (is *IntStack) Empty() bool {
	return is.s.Empty()
}

// Stack with an implementation of strings.
type StringStack struct {
	s Stack
}

// New returns an initialized StringStack.
func NewString() *StringStack {
	return new(StringStack)
}

// Pushes a new string value onto the stack.
func (ss *StringStack) Push(value string) {
	ss.s.Push(value)
}

// Pops some string data from this Stack.
func (ss *StringStack) Pop() (string, error) {
	val, err := ss.s.Pop()
	if err != nil {
		return "", err
	}
	return val.(string), nil
}

// Size returns the number of elements in the stack.
func (ss *StringStack) Size() int {
	return ss.s.Size()
}

// Empty returns if the Stack is empty or not.
func (ss *StringStack) Empty() bool {
	return ss.s.Empty()
}

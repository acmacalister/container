package stack

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	s.Push("Scotty")
}

func TestPush(t *testing.T) {
	s := Stack{}
	s.Push("Kirk")
	s.Push("Spock")
	s.Push("Red Shirt")
}

func TestPop(t *testing.T) {
	s := New()
	s.Push("Red Shirt")
	s.Push("Scotty")
	s.Push("Spock")
	val, _ := s.Pop()
	if val != "Spock" {
		t.Error("Pop failed.")
	}
}

func TestSize(t *testing.T) {
	s := New()
	s.Push("Red Shirt")
	s.Push("Scotty")
	s.Push("Spock")
	if s.Size() != 3 {
		t.Error("Size is wrong.")
	}
}

func TestEmpty(t *testing.T) {
	s := New()
	s.Push("Red Shirt")
	s.Push("Scotty")
	s.Push("Spock")
	if s.Empty() {
		t.Error("Empty is incorrect.")
	}
}

func TestNewInt(t *testing.T) {
	s := NewInt()
	s.Push(1)
}

func TestIntPush(t *testing.T) {
	s := IntStack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
}

func TestIntPop(t *testing.T) {
	s := NewInt()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	val, _ := s.Pop()
	if val != 3 {
		t.Error("IntPop failed.")
	}
}

func TestNewString(t *testing.T) {
	s := NewString()
	s.Push("Scotty")
}

func TestStringPush(t *testing.T) {
	s := StringStack{}
	s.Push("Kirk")
	s.Push("Spock")
	s.Push("Red Shirt")
}

func TestStringPop(t *testing.T) {
	s := NewString()
	s.Push("Red Shirt")
	s.Push("Scotty")
	s.Push("Spock")
	val, _ := s.Pop()
	if val != "Spock" {
		t.Error("StringPop failed.")
	}
}

func BenchmarkPush(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		s.Push(b.N)
	}
}

func BenchmarkPop(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		s.Push(b.N)
		s.Pop()
	}
}

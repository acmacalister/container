package queue

import (
	"testing"
)

func TestNew(t *testing.T) {
	q := New()
	q.Push("Scotty")
}

func TestPush(t *testing.T) {
	q := Queue{}
	q.Push("Kirk")
	q.Push("Spock")
	q.Push("Red Shirt")
}

func TestPop(t *testing.T) {
	q := New()
	q.Push("Red Shirt")
	q.Push("Scotty")
	q.Push("Spock")
	val, _ := q.Pop()
	if val != "Red Shirt" {
		t.Errorf("Pop failed.")
	}
}

func TestNewInt(t *testing.T) {
	q := NewInt()
	q.Push(1)
}

func TestIntPush(t *testing.T) {
	q := IntQueue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
}

func TestIntPop(t *testing.T) {
	q := NewInt()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	val, _ := q.Pop()
	if val != 1 {
		t.Errorf("IntPop failed.")
	}
}

func TestNewString(t *testing.T) {
	q := NewString()
	q.Push("Scotty")
}

func TestStringPush(t *testing.T) {
	q := StringQueue{}
	q.Push("Kirk")
	q.Push("Spock")
	q.Push("Red Shirt")
}

func TestStringPop(t *testing.T) {
	q := NewString()
	q.Push("Red Shirt")
	q.Push("Scotty")
	q.Push("Spock")
	val, _ := q.Pop()
	if val != "Red Shirt" {
		t.Errorf("StringPop failed.")
	}
}

func BenchmarkPush(b *testing.B) {
	q := New()
	for i := 0; i < b.N; i++ {
		q.Push(b.N)
	}
}

func BenchmarkPop(b *testing.B) {
	q := New()
	for i := 0; i < b.N; i++ {
		q.Push(b.N)
		q.Pop()
	}
}

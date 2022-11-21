package LgoRand

import (
	"testing"
)

func TestSafeRand(t *testing.T) {
	rd := New()
	expect := rd.Intn(1)
	for i := 0; i < 100; i++ {
		got := rd.Intn(1)
		if got != expect {
			t.Fatalf("expected %d but got %d", expect, got)
		}
	}
}

func TestUnsafeRand(t *testing.T) {
	rd := NewUnsafe()
	expect := rd.Intn(1)
	for i := 0; i < 100; i++ {
		got := rd.Intn(1)
		if got != expect {
			t.Fatalf("expected %d but got %d", expect, got)
		}
	}
}

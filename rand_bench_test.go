package LgoRand

import (
	"math/rand"
	"testing"
)

func BenchmarkStandardRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(1000)
	}
}

func BenchmarkLgoRand(b *testing.B) {
	rd := New()
	for i := 0; i < b.N; i++ {
		rd.Intn(1000)
	}
}

func BenchmarkUnsafeLgoRand(b *testing.B) {
	rd := NewUnsafe()
	for i := 0; i < b.N; i++ {
		rd.Intn(1000)
	}
}

func BenchmarkStandardRandWithConcurrent(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Intn(1000)
		}
	})
}

func BenchmarkLgoRandWithConcurrent(b *testing.B) {
	rd := New()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rd.Intn(1000)
		}
	})
}
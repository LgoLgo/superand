package superand

import (
	"math/rand"
	"sync"
	"time"
)

type randPool struct {
	p *sync.Pool
}

// Int63 Seed Uint64 implements the Source interface in the standard library
func (r *randPool) Int63() int64 {
	v := r.p.Get()
	defer r.p.Put(v)
	return v.(rand.Source).Int63()
}

func (r *randPool) Seed(seed int64) {
	v := r.p.Get()
	defer r.p.Put(v)
	v.(rand.Source).Seed(seed)
}

func (r *randPool) Uint64() uint64 {
	v := r.p.Get()
	defer r.p.Put(v)
	return v.(rand.Source64).Uint64()
}

func newPoolSource() *randPool {
	s := &randPool{}
	p := &sync.Pool{New: func() any {
		return rand.NewSource(time.Now().Unix())
	}}
	s.p = p
	return s
}

// New create a thread-safe rand.Rand object
func New() *rand.Rand {
	return rand.New(newPoolSource())
}

// NewUnsafe create a non-thread-safe rand.Rand object with current time seed
func NewUnsafe() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().Unix()))
}

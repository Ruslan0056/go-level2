package RWMutex

import (
	"sync"
)

type Set struct {
	sync.RWMutex
	mm map[int]struct{}
}

const count = 100

func RWMutexOne() *Set {
	for i := 0; i <= 0.1*count; i++ {
		go func() {
			NewSet().Add(i)
		}()
	}
	for i := 0; i <= 0.9*count; i++ {
		go func() {
			NewSet().Has(i)
		}()
	}
	return &Set{}
}

func RWMutexTwo() *Set {
	for i := 0; i <= 0.5*count; i++ {
		go func() {
			NewSet().Add(i)
		}()
	}
	for i := 0; i <= 0.5*count; i++ {
		go func() {
			NewSet().Has(i)
		}()
	}
	return &Set{}
}

func RWMutexThree() *Set {
	for i := 0; i <= 0.9*count; i++ {
		go func() {
			NewSet().Add(i)
		}()
	}
	for i := 0; i <= 0.1*count; i++ {
		go func() {
			NewSet().Has(i)
		}()
	}
	return &Set{}
}

func NewSet() *Set {
	return &Set{
		mm: map[int]struct{}{},
	}
}

func (s *Set) Add(i int) {
	s.RLock()
	s.mm[i] = struct{}{}
	s.RUnlock()
}

func (s *Set) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}

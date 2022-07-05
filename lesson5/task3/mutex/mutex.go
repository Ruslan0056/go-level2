package Mutex

import (
	"sync"
)

type Set struct {
	sync.Mutex
	mm map[int]struct{}
}

const count = 100

func MutexOne() *Set {
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

func MutexTwo() *Set {
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

func MutexThree() *Set {
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
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *Set) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

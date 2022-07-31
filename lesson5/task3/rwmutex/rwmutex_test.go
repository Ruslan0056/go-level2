package RWMutex

import "testing"

func BenchmarkRWMutexOne(b *testing.B) {
	for i:=0; i < b.N; i++{
		RWMutexOne()
	}
}

func BenchmarkRWMutexTwo(b *testing.B) {
	for i:=0; i < b.N; i++{
		RWMutexTwo()
	}
}

func BenchmarkRWMutexThree(b *testing.B) {
	for i:=0; i < b.N; i++{
		RWMutexThree()
	}
}
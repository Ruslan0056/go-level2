package Mutex

import "testing"

func BenchmarkMutexOne(b *testing.B) {
	for i:=0; i < b.N; i++{
		MutexOne()
	}
}

func BenchmarkMutexTwo(b *testing.B) {
	for i:=0; i < b.N; i++{
		MutexTwo()
	}
}

func BenchmarkMutexThree(b *testing.B) {
	for i:=0; i < b.N; i++{
		MutexThree()
	}
}
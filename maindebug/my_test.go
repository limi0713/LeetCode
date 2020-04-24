package main

import "testing"

// BenchmarkCopy copy
func BenchmarkCopy(b *testing.B) {
	numsO := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numsT := make([]int, len(numsO))
	for i := 0; i < b.N; i++ {
		copy(numsT, numsO)
	}
}

// BenchmarkAppend append
func BenchmarkAppend(b *testing.B) {
	numsO := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numsT := make([]int, 0)
	for i := 0; i < b.N; i++ {
		numsT = append(numsT, numsO...)
	}
}
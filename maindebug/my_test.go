package main

import "testing"

// BenchmarkCopy copy
func BenchmarkCopy(b *testing.B) {
	numsO := make([]int, 10000)
	// numsT := make([]int, len(numsO), len(numsO))
	for i := 0; i < b.N; i++ {
		numsT := make([]int, len(numsO), len(numsO))
		copy(numsT, numsO)
	}
}

// BenchmarkAppend append
func BenchmarkAppend(b *testing.B) {
	numsO := make([]int, 10000)
	numsT := make([]int, 0, len(numsO))
	for i := 0; i < b.N; i++ {
		numsT = append(numsT[:0:0], numsO...)
	}
}
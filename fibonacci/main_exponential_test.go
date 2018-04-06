package main

import (
	"testing"
)

// The Go testing package contains a benchmarking facility that can be used to examine the performance of your Go code.
// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
func benchmarkRecursive(i int, b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		exponentialFibonacci(i)
	}
}

func BenchmarkRecursive1(b *testing.B)  { benchmarkRecursive(25, b) }
func BenchmarkRecursive2(b *testing.B)  { benchmarkRecursive(25, b) }
func BenchmarkRecursive3(b *testing.B)  { benchmarkRecursive(25, b) }
func BenchmarkRecursive10(b *testing.B) { benchmarkRecursive(25, b) }
func BenchmarkRecursive20(b *testing.B) { benchmarkRecursive(25, b) }
func BenchmarkRecursive40(b *testing.B) { benchmarkRecursive(25, b) }

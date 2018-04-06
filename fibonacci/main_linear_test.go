package main

import (
	"testing"
)

// The Go testing package contains a benchmarking facility that can be used to examine the performance of your Go code.
// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
func benchmarkLinear(i int, b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		linearFibonacci(i)
	}
}

func BenchmarkLinear1(b *testing.B)  { benchmarkLinear(25, b) }
// func BenchmarkLinear2(b *testing.B)  { benchmarkLinear(25, b) }
// func BenchmarkLinear3(b *testing.B)  { benchmarkLinear(25, b) }
// func BenchmarkLinear10(b *testing.B) { benchmarkLinear(25, b) }
// func BenchmarkLinear20(b *testing.B) { benchmarkLinear(25, b) }
// func BenchmarkLinear40(b *testing.B) { benchmarkLinear(25, b) }

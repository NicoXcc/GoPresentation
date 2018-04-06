package main

import (
	"testing"
)

func TestGoogle(t *testing.T) {
	results := Google("go")
	if len(results) != 3 {
		t.Errorf("should return results for web image and video")
	}
}

// The Go testing package contains a benchmarking facility that can be used to examine the performance of your Go code.
// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
func BenchmarkGoogle(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Google("go")
	}
}

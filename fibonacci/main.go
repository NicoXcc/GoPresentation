package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func linearFibonacci(n int) int64 {
	// Create an int array of size n + 1
	v := make([]int64, n+1)

	// F(0) = 0
	v[0] = 0
	// F(1) = 1
	v[1] = 1

	// F(i) = F(i-1) + F(i-2)
	for i := 2; i <= n; i++ {
		v[i] = v[i-1] + v[i-2]
	}

	// F(n) - return the n-th Fibonacci number
	return v[n]
}

// https://blog.intelligentbee.com/2017/08/01/profiling-web-applications-golang/
func exponentialFibonacci(n int) int64 {
	// F(0) = 0
	if n == 0 {
		return 0
	}

	// F(1) = 1
	if n == 1 {
		return 1
	}

	// F(n) = F(n-1) + F(n-2) - return the n-th Fibonacci number
	return exponentialFibonacci(n-1) + exponentialFibonacci(n-2)
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	fmt.Println(exponentialFibonacci(50))

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}

# Basic Programming in Go

### [install Present tool](http://halyph.com/blog/2015/05/18/golang-presentation-tool.html)

After successfully installing the present tool, runn the following command to start presentation
```
$ cd GoPresentation
GoPresentation $ present
```

Running unit test cases:
```
$ cd GoPresentation/google
google $ go test
```

Running benchmarks:
```
$ cd GoPresentation/google
google $ go test -bench=.
```

getting the cpu performance profile out of benchmarks:
```
$ cd GoPresentation/fibonacci
fibonacci $ go test -bench=BenchmarkRecursive -cpuprofile=cpu-exponential.prof

$ cd GoPresentation/fibonacci
fibonacci$ go test -bench=BenchmarkLinear -cpuprofile=cpu-linear.prof

```
checking the profile generated out of benchmarks:
```
fibonacci $ go tool pprof cpu-exponential.prof
(pprof) $ web
(pprof) $ list exponentialFibonacci
```
```web``` : gives the flow diagram and cpu usage by the functions for various benchmark tests

```list exponentialFibonacci```: gives the granular cpu usage by the instructions given in the exponenetialFibonacci function.

### References:

#### For more info on profiling:
https://blog.golang.org/profiling-go-programs
https://blog.intelligentbee.com/2017/08/01/profiling-web-applications-golang/

#### go tools
https://github.com/campoy/go-tooling-workshop

#### GO programming & and its applications:

https://github.com/golang/talks
https://www.toptal.com/back-end/server-side-io-performance-node-php-java-go
https://blog.golang.org/cover
https://tqdev.com/2017-gophercon-2017-videos-online
https://vimeo.com/53221560

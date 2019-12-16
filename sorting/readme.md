```shell
ƒ(x): go test -bench=. -benchmem
goos: darwin
goarch: amd64
BenchmarkSelection100-8     	   82585	     13695 ns/op	     896 B/op	       1 allocs/op
BenchmarkSelection1000-8    	    2630	    426601 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelection10000-8   	       1	3648744594 ns/op	  802816 B/op	       1 allocs/op
BenchmarkInsertion100-8     	  105118	     11439 ns/op	     896 B/op	       1 allocs/op
BenchmarkInsertion1000-8    	    3817	    287612 ns/op	    8192 B/op	       1 allocs/op
BenchmarkInsertion10000-8   	       1	2575209135 ns/op	  802816 B/op	       1 allocs/op
```

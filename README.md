# Json marhal bench compare

```
goos: darwin
goarch: arm64
pkg: github.com/SealTV/jsonbench
BenchmarkStdJSONMarshal-8       	 2940991	       380.2 ns/op	     264 B/op	       5 allocs/op
BenchmarkCustromJSONMarshal-8   	 1829240	       659.6 ns/op	    1976 B/op	      26 allocs/op
BenchmarkBufferJsonMarshal-8    	 1495398	       829.8 ns/op	    2616 B/op	      27 allocs/op
BenchmarkGoJSONMarshal-8        	 7319736	       164.5 ns/op	     104 B/op	       4 allocs/op
BenchmarkEasyJSONMarshal-8      	39776092	        30.30 ns/op	     112 B/op	       1 allocs/op
PASS
coverage: 91.7% of statements
ok  	github.com/SealTV/jsonbench	8.325s
```
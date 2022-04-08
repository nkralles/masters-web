<img src="/assets/fancylog-icon.svg"  alt="fancylog" width="400px">


fancylog was built in an effort to create stylized logs that was fast and
usable across all the different tech stacks we are using at NASK INC. 
I have eventually decided to take a page out of uber's `bufferpool` package 
and use their buffered pool implementation of the buffer interface
***

### Benchmarks
```text 
nick@Nicks-MacBook-Pro-2 fancylog % go test -bench=. -benchtime=1s
goos: darwin
goarch: amd64
pkg: gitlab.cloud.n-ask.com/n-ask/fancylog
BenchmarkLogger_Trace-16                 1123214              1014 ns/op
BenchmarkLogger_Error-16                 1000000              1018 ns/op
BenchmarkLogger_Info-16                  1000000              1012 ns/op
BenchmarkLogger_InfoNoTs-16              1000000              1012 ns/op
BenchmarkLogger_InfoNoTsColor-16         1000000              1009 ns/op
BenchmarkLogger_Debug-16                  377385              2739 ns/op
BenchmarkLogger_DebugPlain-16             442108              2841 ns/op
BenchmarkLogger_Quiet-16                18950532                61.4 ns/op
BenchmarkLogger_InfoMap-16                647956              1867 ns/op
BenchmarkLogger_DebugMap-16               639224              1902 ns/op
BenchmarkLogger_Infof-16                 1000000              1077 ns/op
BenchmarkLogger_Debugf-16                 399288              2979 ns/op
PASS
ok      gitlab.cloud.n-ask.com/n-ask/fancylog   14.661s


nick@pop-os:~/work/fancylog$ go test -bench=. -benchtime=1s
goos: linux
goarch: amd64
pkg: gitlab.cloud.n-ask.com/n-ask/fancylog
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkLogger_Trace-12                 1767062               683.0 ns/op
BenchmarkLogger_Error-12                 1635012               721.2 ns/op
BenchmarkFmtPrintf-12                    2769434               458.1 ns/op
BenchmarkLogger_Info-12                  1615462               710.2 ns/op
BenchmarkLogger_InfoNoTs-12              1550862               715.0 ns/op
BenchmarkLogger_InfoNoTsColor-12         1808527               692.5 ns/op
BenchmarkLogger_Debug-12                  869169              2384 ns/op
BenchmarkLogger_DebugPlain-12             584377              2345 ns/op
BenchmarkLogger_Quiet-12                 9316052               142.2 ns/op
BenchmarkLogger_InfoMap-12                654340              1937 ns/op
BenchmarkLogger_DebugMap-12               579030              1849 ns/op
BenchmarkLogger_Infof-12                 1563890               699.5 ns/op
BenchmarkLogger_Debugf-12                 841011              2502 ns/op
PASS
ok      gitlab.cloud.n-ask.com/n-ask/fancylog   24.135s

```
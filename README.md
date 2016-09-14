# gobitmapbenchmark
A Go project to benchmark various bitmap implementations (this is not a library!)



### Dependencies

  - go get github.com/willf/bitset
  - go get github.com/RoaringBitmap/roaring
  - go get github.com/fzandona/goroar
  - go get github.com/RoaringBitmap/gocroaring

### Usage

Type

          go test -bench Benchmark -run -

### Sample result

```bash
$ go test -bench Benchmark -run -
BenchmarkIntersectionCRoaring-2                5000               351797 ns/op
BenchmarkIntersectionRoaring-2                 1000              1734584 ns/op
BenchmarkIntersectionGoroar-2                  1000              2017873 ns/op

BenchmarkIntersectionDenseCRoaring-2           5000               282847 ns/op
BenchmarkIntersectionDenseRoaring-2            3000               516019 ns/op
BenchmarkIntersectionDenseGoroar-2             3000               475892 ns/op

BenchmarkUnionCRoaring-2                       1000              1415311 ns/op
BenchmarkUnionRoaring-2                        1000              1799285 ns/op
BenchmarkUnionGoRoar-2                          500              2590627 ns/op

BenchmarkSizeCRoaring-2                    0.2 MB 0.2 MB 0.2 MB 0.2 MB 0.2 MB 0.2 MB 2000000000                0.00 ns/op
BenchmarkSizeRoaring-2                     0.2 MB 0.2 MB 0.2 MB 0.2 MB 0.2 MB 0.2 MB 2000000000                0.00 ns/op
BenchmarkSizeGoroar-2                      0.2 MB 0.2 MB 0.2 MB 0.2 MB 0.2 MB 0.2 MB 2000000000                0.00 ns/op

BenchmarkSetRCoaring-2                     10000000                  211 ns/op
BenchmarkSetRoaring-2                      20000000                   59.7 ns/op
BenchmarkSetGoroar-2                       20000000                   60.8 ns/op

BenchmarkGetTestCRoaring-2                  5000000                  257 ns/op
BenchmarkGetTestRoaring-2                  20000000                  115 ns/op
BenchmarkGetTestGoroar-2                   10000000                  124 ns/op

BenchmarkCountCRoaring-2                   10000000                  163 ns/op
BenchmarkCountRoaring-2                    30000000                   51.5 ns/op
BenchmarkCountGoroar-2                     20000000                   69.2 ns/op
```

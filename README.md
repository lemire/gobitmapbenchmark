# gobitmapbenchmark
A Go project to benchmark various bitmap implementations (this is not a library!)



### Dependencies

  - go get github.com/willf/bitset
  - go get github.com/RoaringBitmap/roaring
  - go get github.com/fzandona/goroar
  - go get github.com/RoaringBitmap/gocroaring

### Usage

For a comparison between gocroaring and roaring on real data, try:

          go test -bench BenchmarkReal -run -

Type

          go test -bench Benchmark -run -

### Sample result

```bash
$  go test -bench BenchmarkReal -run -
gocroaring version  0 . 2 . 38
 ======
read  200  arrays from  realdata/census-income

goos: linux
goarch: amd64
BenchmarkReal/roaring_creating_realdata/census-income-2         	      30	  45028561 ns/op
BenchmarkReal/gocroaring_creating_realdata/census-income-2      	     100	  21941683 ns/op

BenchmarkReal/roaring_clone_realdata/census-income-2            	    3000	    494532 ns/op
BenchmarkReal/gocroaring_clone_realdata/census-income-2         	    2000	   1066473 ns/op


BenchmarkReal/roaring_iterate_realdata/census-income-2          	      20	  81798450 ns/op
BenchmarkReal/gocroaring_iterate_realdata/census-income-2       	       2	 542632617 ns/op

BenchmarkReal/roaring_toArray_realdata/census-income-2          	     100	  12079516 ns/op
BenchmarkReal/gocroaring_toArray_realdata/census-income-2       	     200	   6028955 ns/op

BenchmarkReal/roaring_pairwise_AND_realdata/census-income-2     	    1000	   2384344 ns/op
BenchmarkReal/gocroaring_pairwise_AND_realdata/census-income-2  	    1000	   1365977 ns/op

BenchmarkReal/roaring_pairwise_OR_realdata/census-income-2      	     500	   3357336 ns/op
BenchmarkReal/gocroaring_pairwise_OR_realdata/census-income-2   	    1000	   2104553 ns/op

BenchmarkReal/roaring_fastOR_realdata/census-income-2           	    2000	    854944 ns/op
BenchmarkReal/gocroaring_fastOR_realdata/census-income-2        	   10000	    100659 ns/op


 ======
read  200  arrays from  realdata/census-income_srt

BenchmarkReal/roaring_creating_realdata/census-income_srt-2     	      30	  41720445 ns/op
BenchmarkReal/gocroaring_creating_realdata/census-income_srt-2  	     100	  20083310 ns/op

BenchmarkReal/roaring_clone_realdata/census-income_srt-2        	   10000	    186022 ns/op
BenchmarkReal/gocroaring_clone_realdata/census-income_srt-2     	    3000	    609062 ns/op


BenchmarkReal/roaring_iterate_realdata/census-income_srt-2      	      10	 111885653 ns/op
BenchmarkReal/gocroaring_iterate_realdata/census-income_srt-2   	       3	 468990768 ns/op

BenchmarkReal/roaring_toArray_realdata/census-income_srt-2      	     200	   6391405 ns/op
BenchmarkReal/gocroaring_toArray_realdata/census-income_srt-2   	     500	   3999921 ns/op

BenchmarkReal/roaring_pairwise_AND_realdata/census-income_srt-2 	    2000	   1062861 ns/op
BenchmarkReal/gocroaring_pairwise_AND_realdata/census-income_srt-2         	    2000	    842367 ns/op

BenchmarkReal/roaring_pairwise_OR_realdata/census-income_srt-2             	     500	   3330623 ns/op
BenchmarkReal/gocroaring_pairwise_OR_realdata/census-income_srt-2          	    1000	   1659235 ns/op

BenchmarkReal/roaring_fastOR_realdata/census-income_srt-2                  	    3000	    427755 ns/op
BenchmarkReal/gocroaring_fastOR_realdata/census-income_srt-2               	   10000	    184545 ns/op


 ======
read  200  arrays from  realdata/census1881

BenchmarkReal/roaring_creating_realdata/census1881-2                       	     200	   8533858 ns/op
BenchmarkReal/gocroaring_creating_realdata/census1881-2                    	     300	   5703671 ns/op

BenchmarkReal/roaring_clone_realdata/census1881-2                          	    2000	    577244 ns/op
BenchmarkReal/gocroaring_clone_realdata/census1881-2                       	    2000	   1129899 ns/op


BenchmarkReal/roaring_iterate_realdata/census1881-2                        	     200	   9006506 ns/op
BenchmarkReal/gocroaring_iterate_realdata/census1881-2                     	      20	  77737324 ns/op

BenchmarkReal/roaring_toArray_realdata/census1881-2                        	    1000	   1365105 ns/op
BenchmarkReal/gocroaring_toArray_realdata/census1881-2                     	    1000	   1256569 ns/op

BenchmarkReal/roaring_pairwise_AND_realdata/census1881-2                   	   20000	     71183 ns/op
BenchmarkReal/gocroaring_pairwise_AND_realdata/census1881-2                	    5000	    356198 ns/op

BenchmarkReal/roaring_pairwise_OR_realdata/census1881-2                    	    1000	   1640428 ns/op
BenchmarkReal/gocroaring_pairwise_OR_realdata/census1881-2                 	    2000	   2300969 ns/op

BenchmarkReal/roaring_fastOR_realdata/census1881-2                         	    1000	   1782322 ns/op
BenchmarkReal/gocroaring_fastOR_realdata/census1881-2                      	    2000	    787069 ns/op


 ======
read  200  arrays from  realdata/census1881_srt

BenchmarkReal/roaring_creating_realdata/census1881_srt-2                   	     200	   6782934 ns/op
BenchmarkReal/gocroaring_creating_realdata/census1881_srt-2                	     500	   4005242 ns/op

BenchmarkReal/roaring_clone_realdata/census1881_srt-2                      	    5000	    240649 ns/op
BenchmarkReal/gocroaring_clone_realdata/census1881_srt-2                   	    3000	    839457 ns/op


BenchmarkReal/roaring_iterate_realdata/census1881_srt-2                    	     100	  12708996 ns/op
BenchmarkReal/gocroaring_iterate_realdata/census1881_srt-2                 	      30	  52689646 ns/op

BenchmarkReal/roaring_toArray_realdata/census1881_srt-2                    	    2000	    964374 ns/op
BenchmarkReal/gocroaring_toArray_realdata/census1881_srt-2                 	    2000	    758920 ns/op

BenchmarkReal/roaring_pairwise_AND_realdata/census1881_srt-2               	   30000	     45061 ns/op
BenchmarkReal/gocroaring_pairwise_AND_realdata/census1881_srt-2            	    3000	    410062 ns/op

BenchmarkReal/roaring_pairwise_OR_realdata/census1881_srt-2                	    1000	   2053263 ns/op
BenchmarkReal/gocroaring_pairwise_OR_realdata/census1881_srt-2             	    2000	    932609 ns/op

BenchmarkReal/roaring_fastOR_realdata/census1881_srt-2                     	    2000	    982225 ns/op
BenchmarkReal/gocroaring_fastOR_realdata/census1881_srt-2                  	    2000	    726187 ns/op


 ======
read  200  arrays from  realdata/uscensus2000

BenchmarkReal/roaring_creating_realdata/uscensus2000-2                     	    3000	    475542 ns/op
BenchmarkReal/gocroaring_creating_realdata/uscensus2000-2                  	    2000	   1299098 ns/op

BenchmarkReal/roaring_clone_realdata/uscensus2000-2                        	   10000	    178872 ns/op
BenchmarkReal/gocroaring_clone_realdata/uscensus2000-2                     	    3000	    671385 ns/op


BenchmarkReal/roaring_iterate_realdata/uscensus2000-2                      	   10000	    154059 ns/op
BenchmarkReal/gocroaring_iterate_realdata/uscensus2000-2                   	    2000	    669725 ns/op

BenchmarkReal/roaring_toArray_realdata/uscensus2000-2                      	   30000	     42051 ns/op
BenchmarkReal/gocroaring_toArray_realdata/uscensus2000-2                   	   30000	     56623 ns/op

BenchmarkReal/roaring_pairwise_AND_realdata/uscensus2000-2                 	  100000	     17584 ns/op
BenchmarkReal/gocroaring_pairwise_AND_realdata/uscensus2000-2              	    5000	    425028 ns/op

BenchmarkReal/roaring_pairwise_OR_realdata/uscensus2000-2                  	    3000	    475342 ns/op
BenchmarkReal/gocroaring_pairwise_OR_realdata/uscensus2000-2               	    2000	    815871 ns/op

BenchmarkReal/roaring_fastOR_realdata/uscensus2000-2                       	    1000	   1875687 ns/op
BenchmarkReal/gocroaring_fastOR_realdata/uscensus2000-2                    	    2000	   1298083 ns/op


 ======
read  200  arrays from  realdata/weather_sept_85

BenchmarkReal/roaring_creating_realdata/weather_sept_85-2                  	      20	  88877032 ns/op
BenchmarkReal/gocroaring_creating_realdata/weather_sept_85-2               	      30	  45375611 ns/op

BenchmarkReal/roaring_clone_realdata/weather_sept_85-2                     	    1000	   1997008 ns/op
BenchmarkReal/gocroaring_clone_realdata/weather_sept_85-2                  	     500	   3450691 ns/op


BenchmarkReal/roaring_iterate_realdata/weather_sept_85-2                   	      10	 149616008 ns/op
BenchmarkReal/gocroaring_iterate_realdata/weather_sept_85-2                	       1	1007214720 ns/op

BenchmarkReal/roaring_toArray_realdata/weather_sept_85-2                   	      50	  24092363 ns/op
BenchmarkReal/gocroaring_toArray_realdata/weather_sept_85-2                	     100	  13618656 ns/op

BenchmarkReal/roaring_pairwise_AND_realdata/weather_sept_85-2              	     200	   8969499 ns/op
BenchmarkReal/gocroaring_pairwise_AND_realdata/weather_sept_85-2           	     500	   3339475 ns/op

BenchmarkReal/roaring_pairwise_OR_realdata/weather_sept_85-2               	     100	  11846355 ns/op
BenchmarkReal/gocroaring_pairwise_OR_realdata/weather_sept_85-2            	     300	   7185589 ns/op

BenchmarkReal/roaring_fastOR_realdata/weather_sept_85-2                    	     500	   3921765 ns/op
BenchmarkReal/gocroaring_fastOR_realdata/weather_sept_85-2                 	    3000	    574939 ns/op


 ======
read  200  arrays from  realdata/weather_sept_85_srt

BenchmarkReal/roaring_creating_realdata/weather_sept_85_srt-2              	      10	 110679813 ns/op
BenchmarkReal/gocroaring_creating_realdata/weather_sept_85_srt-2           	      30	  53385555 ns/op

BenchmarkReal/roaring_clone_realdata/weather_sept_85_srt-2                 	    3000	    399898 ns/op
BenchmarkReal/gocroaring_clone_realdata/weather_sept_85_srt-2              	    3000	   1194580 ns/op


BenchmarkReal/roaring_iterate_realdata/weather_sept_85_srt-2               	       5	 297901757 ns/op
BenchmarkReal/gocroaring_iterate_realdata/weather_sept_85_srt-2            	       1	1239943929 ns/op

BenchmarkReal/roaring_toArray_realdata/weather_sept_85_srt-2               	     100	  21274002 ns/op
BenchmarkReal/gocroaring_toArray_realdata/weather_sept_85_srt-2            	     100	  13973855 ns/op

BenchmarkReal/roaring_pairwise_AND_realdata/weather_sept_85_srt-2          	    1000	   1685735 ns/op
BenchmarkReal/gocroaring_pairwise_AND_realdata/weather_sept_85_srt-2       	    1000	   1290481 ns/op

BenchmarkReal/roaring_pairwise_OR_realdata/weather_sept_85_srt-2           	     200	   7564440 ns/op
BenchmarkReal/gocroaring_pairwise_OR_realdata/weather_sept_85_srt-2        	    1000	   2335059 ns/op

BenchmarkReal/roaring_fastOR_realdata/weather_sept_85_srt-2                	    5000	    347722 ns/op
BenchmarkReal/gocroaring_fastOR_realdata/weather_sept_85_srt-2             	   10000	    184497 ns/op


 ======
read  200  arrays from  realdata/wikileaks-noquotes

BenchmarkReal/roaring_creating_realdata/wikileaks-noquotes-2               	     300	   4995984 ns/op
BenchmarkReal/gocroaring_creating_realdata/wikileaks-noquotes-2            	     500	   3270301 ns/op

BenchmarkReal/roaring_clone_realdata/wikileaks-noquotes-2                  	    5000	    240132 ns/op
BenchmarkReal/gocroaring_clone_realdata/wikileaks-noquotes-2               	    2000	    763437 ns/op


BenchmarkReal/roaring_iterate_realdata/wikileaks-noquotes-2                	     300	   5567055 ns/op
BenchmarkReal/gocroaring_iterate_realdata/wikileaks-noquotes-2             	     100	  21921597 ns/op

BenchmarkReal/roaring_toArray_realdata/wikileaks-noquotes-2                	    2000	    667282 ns/op
BenchmarkReal/gocroaring_toArray_realdata/wikileaks-noquotes-2             	    2000	    725658 ns/op

BenchmarkReal/roaring_pairwise_AND_realdata/wikileaks-noquotes-2           	    3000	    474999 ns/op
BenchmarkReal/gocroaring_pairwise_AND_realdata/wikileaks-noquotes-2        	    2000	    715276 ns/op

BenchmarkReal/roaring_pairwise_OR_realdata/wikileaks-noquotes-2            	     500	   2280622 ns/op
BenchmarkReal/gocroaring_pairwise_OR_realdata/wikileaks-noquotes-2         	    2000	   1207009 ns/op

BenchmarkReal/roaring_fastOR_realdata/wikileaks-noquotes-2                 	    3000	    349567 ns/op
BenchmarkReal/gocroaring_fastOR_realdata/wikileaks-noquotes-2              	   10000	    210815 ns/op


 ======
read  200  arrays from  realdata/wikileaks-noquotes_srt

BenchmarkReal/roaring_creating_realdata/wikileaks-noquotes_srt-2           	     500	   3447683 ns/op
BenchmarkReal/gocroaring_creating_realdata/wikileaks-noquotes_srt-2        	     500	   3233452 ns/op

BenchmarkReal/roaring_clone_realdata/wikileaks-noquotes_srt-2              	   10000	    172535 ns/op
BenchmarkReal/gocroaring_clone_realdata/wikileaks-noquotes_srt-2           	    3000	    759000 ns/op


BenchmarkReal/roaring_iterate_realdata/wikileaks-noquotes_srt-2            	     300	   5334177 ns/op
BenchmarkReal/gocroaring_iterate_realdata/wikileaks-noquotes_srt-2         	     100	  22504487 ns/op

BenchmarkReal/roaring_toArray_realdata/wikileaks-noquotes_srt-2            	    3000	    403184 ns/op
BenchmarkReal/gocroaring_toArray_realdata/wikileaks-noquotes_srt-2         	    5000	    341512 ns/op

BenchmarkReal/roaring_pairwise_AND_realdata/wikileaks-noquotes_srt-2       	   10000	    130692 ns/op
BenchmarkReal/gocroaring_pairwise_AND_realdata/wikileaks-noquotes_srt-2    	    2000	    560423 ns/op

BenchmarkReal/roaring_pairwise_OR_realdata/wikileaks-noquotes_srt-2        	    1000	   1354152 ns/op
BenchmarkReal/gocroaring_pairwise_OR_realdata/wikileaks-noquotes_srt-2     	    2000	    893540 ns/op

BenchmarkReal/roaring_fastOR_realdata/wikileaks-noquotes_srt-2             	   10000	    192324 ns/op
BenchmarkReal/gocroaring_fastOR_realdata/wikileaks-noquotes_srt-2          	   10000	    139796 ns/op


PASS
ok  	_/home/dlemire/CVS/github/gobitmapbenchmark	254.456s
```

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

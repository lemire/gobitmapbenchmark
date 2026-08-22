# gobitmapbenchmark

A Go program to compare bitmap implementations. This is not a library.

It measures uncompressed bitsets and two Roaring Bitmap implementations:
a pure-Go port, and a cgo wrapper around [CRoaring](https://github.com/RoaringBitmap/CRoaring).

## Implementations

| Package | Notes |
| --- | --- |
| [github.com/bits-and-blooms/bitset](https://github.com/bits-and-blooms/bitset) | Uncompressed bitset (formerly `github.com/willf/bitset`) |
| [github.com/RoaringBitmap/roaring/v2](https://github.com/RoaringBitmap/roaring) | Pure Go Roaring bitmaps |
| [github.com/RoaringBitmap/gocroaring](https://github.com/RoaringBitmap/gocroaring) | cgo wrapper for CRoaring |

The older [goroar](https://github.com/fzandona/goroar) implementation is no
longer included. It has been unmaintained since 2014.

## Requirements

- Go 1.24 or later
- A C compiler (`gocroaring` uses cgo)

```
go test -count=1
```

## Usage

Synthetic microbenchmarks (intersection, union, set, contains, count, iterate):

```
go test -bench . -run '^$' -skip Real -benchmem
```

Real datasets under `realdata/` (census, weather, wikileaks, and sorted
variants):

```
go test -bench BenchmarkReal -run '^$' -benchmem
```

Serialized size of two random sets (`N1=650000` values in `[0, 15000000)` and
`N2=650000` values in `[0, 100000000)`):

```
go test -run TestSize -v
```

The real-data files are comma-separated 32-bit integer sets. See
[`realdata/README.md`](realdata/README.md).

## Sample results

Recorded with Go 1.24.3 on an Apple M4 Max (`darwin/arm64`), using
`bitset` v1.25.0, `roaring/v2` v2.25.0, and `gocroaring` v1.0.0 (CRoaring 4.5.0).
Numbers are illustrative; run the commands above on your machine.

### Serialized size (`TestSize`)

| Implementation | Size |
| --- | ---: |
| bitset | 13.709 MB |
| roaring | 2.462 MB |
| gocroaring | 2.463 MB |

### Synthetic

Time (`ns/op`):

| Benchmark | bitset | roaring | gocroaring |
| --- | ---: | ---: | ---: |
| Intersection | 209,236 | 745,625 | 719,899 |
| IntersectionDense | 26,519 | 63,484 | 45,144 |
| Union | 801,288 | 994,885 | 921,073 |
| Set | 3.320 | 18.84 | 49.41 |
| Contains | 3.291 | 51.57 | 81.81 |
| Count | 3,554 | 14.76 | 33.02 |
| Iterate | 161,471 | 246,312 | 1,714,062 |
| SparseIterate | 878,637 | 196,074 | 2,041,873 |

Memory (`B/op` / `allocs/op`):

| Benchmark | bitset | roaring | gocroaring |
| --- | ---: | ---: | ---: |
| Intersection | 1,876,006 / 2 | 224,169 / 481 | 8 / 1 |
| IntersectionDense | 188,449 / 2 | 113,560 / 60 | 8 / 1 |
| Union | 12,501,031 / 2 | 2,792,967 / 3,086 | 8 / 1 |
| Set | 0 / 0 | 0 / 0 | 4 / 1 |
| Contains | 0 / 0 | 0 / 0 | 0 / 0 |
| Count | 0 / 0 | 0 / 0 | 0 / 0 |
| Iterate | 0 / 0 | 112 / 1 | 16 / 1 |
| SparseIterate | 0 / 0 | 112 / 1 | 16 / 1 |

### Real data

Each table compares pure-Go `roaring` with cgo `gocroaring` on one dataset.

#### census-income

| Operation | roaring ns/op | gocroaring ns/op | roaring B/op | gocroaring B/op | roaring allocs | gocroaring allocs |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| creating | 16,589,097 | 16,199,160 | 9,031,890 | 3,392 | 9,330 | 201 |
| clone | 342,794 | 326,833 | 2,353,336 | 3,392 | 2,537 | 201 |
| iterate | 31,280,922 | 213,813,408 | 22,400 | 3,200 | 200 | 200 |
| toArray | 5,879,046 | 4,960,540 | 28,083,329 | 28,083,321 | 201 | 201 |
| pairwise AND | 1,584,371 | 1,211,460 | 3,065,401 | 1,592 | 3,824 | 199 |
| pairwise OR | 1,532,910 | 1,248,313 | 3,624,378 | 1,592 | 2,808 | 199 |
| fast OR | 389,339 | 324,923 | 42,464 | 1,800 | 24 | 2 |

#### census-income_srt

| Operation | roaring ns/op | gocroaring ns/op | roaring B/op | gocroaring B/op | roaring allocs | gocroaring allocs |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| creating | 15,796,842 | 14,681,891 | 9,082,342 | 3,392 | 10,076 | 201 |
| clone | 121,144 | 194,735 | 540,376 | 3,392 | 2,383 | 201 |
| iterate | 21,827,623 | 188,387,396 | 22,400 | 3,200 | 200 | 200 |
| toArray | 3,729,085 | 3,169,311 | 24,748,397 | 24,748,372 | 201 | 201 |
| pairwise AND | 451,692 | 339,413 | 647,613 | 1,592 | 3,567 | 199 |
| pairwise OR | 5,222,479 | 772,519 | 27,072,048 | 1,592 | 19,157 | 199 |
| fast OR | 178,874 | 126,711 | 53,056 | 1,800 | 29 | 2 |

#### census1881

| Operation | roaring ns/op | gocroaring ns/op | roaring B/op | gocroaring B/op | roaring allocs | gocroaring allocs |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| creating | 3,500,324 | 3,553,999 | 6,499,942 | 3,392 | 12,940 | 201 |
| clone | 338,981 | 354,890 | 2,103,482 | 3,392 | 3,929 | 201 |
| iterate | 2,979,836 | 31,061,332 | 22,400 | 3,200 | 200 | 200 |
| toArray | 598,546 | 495,695 | 4,068,245 | 4,068,240 | 200 | 200 |
| pairwise AND | 26,135 | 159,061 | 45,864 | 1,592 | 450 | 199 |
| pairwise OR | 3,882,382 | 1,132,140 | 23,262,300 | 1,592 | 18,938 | 199 |
| fast OR | 2,608,709 | 671,615 | 2,018,879 | 1,800 | 464 | 2 |

#### census1881_srt

| Operation | roaring ns/op | gocroaring ns/op | roaring B/op | gocroaring B/op | roaring allocs | gocroaring allocs |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| creating | 2,842,899 | 2,414,933 | 2,213,217 | 3,392 | 16,165 | 201 |
| clone | 131,131 | 235,569 | 316,173 | 3,392 | 6,077 | 201 |
| iterate | 2,745,233 | 21,508,570 | 22,400 | 3,200 | 200 | 200 |
| toArray | 545,365 | 470,329 | 2,776,355 | 2,776,353 | 200 | 200 |
| pairwise AND | 21,580 | 164,361 | 39,888 | 1,592 | 923 | 199 |
| pairwise OR | 330,432 | 351,363 | 707,542 | 1,592 | 12,068 | 199 |
| fast OR | 627,913 | 255,945 | 821,871 | 1,800 | 561 | 2 |

#### uscensus2000

| Operation | roaring ns/op | gocroaring ns/op | roaring B/op | gocroaring B/op | roaring allocs | gocroaring allocs |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| creating | 167,577 | 324,584 | 233,298 | 3,392 | 6,342 | 201 |
| clone | 82,810 | 233,873 | 145,289 | 3,392 | 5,443 | 201 |
| iterate | 36,179 | 293,381 | 22,400 | 3,200 | 200 | 200 |
| toArray | 14,464 | 22,502 | 25,968 | 25,968 | 200 | 200 |
| pairwise AND | 6,680 | 156,677 | 16,482 | 1,592 | 241 | 199 |
| pairwise OR | 196,919 | 334,682 | 411,597 | 1,592 | 10,724 | 199 |
| fast OR | 109,795 | 487,362 | 65,832 | 1,800 | 1,700 | 2 |

#### weather_sept_85

| Operation | roaring ns/op | gocroaring ns/op | roaring B/op | gocroaring B/op | roaring allocs | gocroaring allocs |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| creating | 34,691,750 | 34,815,911 | 32,349,597 | 3,392 | 30,529 | 201 |
| clone | 1,136,505 | 2,247,517 | 9,054,955 | 3,392 | 6,713 | 201 |
| iterate | 56,461,092 | 402,331,667 | 22,400 | 3,200 | 200 | 200 |
| toArray | 11,228,982 | 10,445,023 | 51,990,914 | 51,990,902 | 200 | 200 |
| pairwise AND | 5,280,902 | 4,825,924 | 3,543,417 | 1,592 | 6,194 | 199 |
| pairwise OR | 5,703,905 | 5,320,256 | 15,048,226 | 1,592 | 8,606 | 199 |
| fast OR | 2,003,237 | 1,827,189 | 132,720 | 1,800 | 73 | 2 |

#### weather_sept_85_srt

| Operation | roaring ns/op | gocroaring ns/op | roaring B/op | gocroaring B/op | roaring allocs | gocroaring allocs |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| creating | 42,762,407 | 40,305,885 | 26,052,229 | 3,392 | 27,368 | 201 |
| clone | 271,905 | 323,073 | 838,615 | 3,392 | 5,337 | 201 |
| iterate | 57,588,646 | 498,351,903 | 22,400 | 3,200 | 200 | 200 |
| toArray | 8,684,755 | 8,600,311 | 64,914,541 | 64,914,419 | 202 | 201 |
| pairwise AND | 811,007 | 534,931 | 588,663 | 1,592 | 6,350 | 199 |
| pairwise OR | 3,028,185 | 1,060,899 | 7,484,388 | 1,592 | 20,812 | 199 |
| fast OR | 124,608 | 56,112 | 146,346 | 1,800 | 177 | 2 |

#### wikileaks-noquotes

| Operation | roaring ns/op | gocroaring ns/op | roaring B/op | gocroaring B/op | roaring allocs | gocroaring allocs |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| creating | 2,440,103 | 1,882,933 | 2,238,945 | 3,392 | 21,541 | 201 |
| clone | 118,397 | 268,028 | 311,308 | 3,392 | 4,785 | 201 |
| iterate | 1,344,585 | 8,840,031 | 22,400 | 3,200 | 200 | 200 |
| toArray | 382,020 | 343,992 | 1,168,516 | 1,168,522 | 200 | 200 |
| pairwise AND | 165,292 | 299,451 | 86,280 | 1,592 | 2,220 | 199 |
| pairwise OR | 1,400,975 | 448,117 | 4,754,722 | 1,592 | 21,235 | 199 |
| fast OR | 110,944 | 72,997 | 186,713 | 1,800 | 101 | 2 |

#### wikileaks-noquotes_srt

| Operation | roaring ns/op | gocroaring ns/op | roaring B/op | gocroaring B/op | roaring allocs | gocroaring allocs |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| creating | 1,489,074 | 1,534,888 | 1,487,453 | 3,392 | 14,566 | 201 |
| clone | 73,449 | 247,063 | 149,913 | 3,392 | 4,151 | 201 |
| iterate | 1,122,314 | 8,933,274 | 22,400 | 3,200 | 200 | 200 |
| toArray | 234,169 | 219,990 | 1,206,259 | 1,206,261 | 200 | 200 |
| pairwise AND | 50,258 | 248,177 | 69,040 | 1,592 | 1,485 | 199 |
| pairwise OR | 3,046,920 | 363,320 | 18,377,363 | 1,592 | 19,775 | 199 |
| fast OR | 105,923 | 55,411 | 296,074 | 1,800 | 345 | 2 |

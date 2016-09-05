package roaring

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/RoaringBitmap/gocroaring"
	"github.com/RoaringBitmap/roaring"
	"github.com/fzandona/goroar"
	"github.com/willf/bitset"
)

// BENCHMARKS, to run them type "go test -bench Benchmark -run -"
const SZ1 = 15000000
const N1 = 650000

const SZ2 = 100000000
const N2 = 650000

// go test -bench BenchmarkIntersection -run -
func BenchmarkIntersectionBitset(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := bitset.New(0)
	for i := 0; i < N1; i++ {
		s1.Set(uint(r.Int31n(int32(SZ1))))
	}
	s2 := bitset.New(0)
	for i := 0; i < N2; i++ {
		s2.Set(uint(r.Int31n(int32(SZ2))))
	}
	b.StartTimer()
	card := uint(0)
	for j := 0; j < b.N; j++ {
		s3 := s1.Intersection(s2)
		card = card + s3.Count()
	}
}

// go test -bench BenchmarkIntersection -run -
func BenchmarkIntersectionCRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := gocroaring.New()
	for i := 0; i < N1; i++ {
		s1.Add(uint32(r.Int31n(int32(SZ1))))
	}
	s2 := gocroaring.New()
	for i := 0; i < N2; i++ {
		s2.Add(uint32(r.Int31n(int32(SZ2))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := gocroaring.And(s1, s2)
		card = card + int(s3.Cardinality())
	}
}

// go test -bench BenchmarkIntersection -run -
func BenchmarkIntersectionRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := roaring.New()
	for i := 0; i < N1; i++ {
		s1.Add(uint32(r.Int31n(int32(SZ1))))
	}
	s2 := roaring.New()
	for i := 0; i < N2; i++ {
		s2.Add(uint32(r.Int31n(int32(SZ2))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := roaring.And(s1, s2)
		card = card + int(s3.GetCardinality())
	}
}

// go test -bench BenchmarkIntersection -run -
func BenchmarkIntersectionGoroar(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := goroar.New()
	for i := 0; i < N1; i++ {
		s1.Add(uint32(r.Int31n(int32(SZ1))))
	}
	s2 := goroar.New()
	for i := 0; i < N2; i++ {
		s2.Add(uint32(r.Int31n(int32(SZ2))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := s1.Clone()
		s3.And(s2)
		card = card + s3.Cardinality()
	}
}

// go test -bench BenchmarkIntersectionDense -run -
func BenchmarkIntersectionDenseBitset(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := bitset.New(0)
	for i := 0; i < N1; i++ {
		s1.Set(uint(r.Int31n(int32(SZ1))))
	}
	s2 := bitset.New(0)
	for i := 0; i < N2; i++ {
		s2.Set(uint(r.Int31n(int32(SZ2))))
	}
	b.StartTimer()
	card := uint(0)
	for j := 0; j < b.N; j++ {
		s3 := s1.Intersection(s2)
		card = card + s3.Count()
	}
}

// go test -bench BenchmarkIntersectionDense -run -
func BenchmarkIntersectionDenseCRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := gocroaring.New()
	for i := 0; i < N1; i++ {
		s1.Add(uint32(r.Int31n(int32(SZ1))))
	}
	s2 := gocroaring.New()
	for i := 0; i < N2; i++ {
		s2.Add(uint32(r.Int31n(int32(SZ2))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := gocroaring.And(s1, s2)
		card = card + int(s3.Cardinality())
	}
}

// go test -bench BenchmarkIntersectionDense -run -
func BenchmarkIntersectionDenseRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := roaring.New()
	for i := 0; i < N1; i++ {
		s1.Add(uint32(r.Int31n(int32(SZ1))))
	}
	s2 := roaring.New()
	for i := 0; i < N2; i++ {
		s2.Add(uint32(r.Int31n(int32(SZ2))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := roaring.And(s1, s2)
		card = card + int(s3.GetCardinality())
	}
}

// go test -bench BenchmarkIntersectionDense -run -
func BenchmarkIntersectionDenseGoroar(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := goroar.New()
	for i := 0; i < N1; i++ {
		s1.Add(uint32(r.Int31n(int32(SZ1))))
	}
	s2 := goroar.New()
	for i := 0; i < N2; i++ {
		s2.Add(uint32(r.Int31n(int32(SZ2))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := s1.Clone()
		s3.And(s2)
		card = card + s3.Cardinality()
	}
}

// go test -bench BenchmarkUnion -run -
func BenchmarkUnionBitset(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := bitset.New(0)
	for i := 0; i < N1; i++ {
		s1.Set(uint(r.Int31n(int32(SZ1))))
	}
	s2 := bitset.New(0)
	for i := 0; i < N2; i++ {
		s2.Set(uint(r.Int31n(int32(SZ2))))
	}
	b.StartTimer()
	card := uint(0)
	for j := 0; j < b.N; j++ {
		s3 := s1.Union(s2)
		card = card + s3.Count()
	}
}

// go test -bench BenchmarkUnion -run -
func BenchmarkUnionCRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := gocroaring.New()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(uint32(r.Int31n(int32(sz))))
	}
	s2 := gocroaring.New()
	sz = 1000000
	initsize = 650
	for i := 0; i < initsize; i++ {
		s2.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := gocroaring.Or(s1, s2)
		card = card + int(s3.Cardinality())
	}
}

// go test -bench BenchmarkUnion -run -
func BenchmarkUnionRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := roaring.New()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(uint32(r.Int31n(int32(sz))))
	}
	s2 := roaring.New()
	sz = 1000000
	initsize = 650
	for i := 0; i < initsize; i++ {
		s2.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := roaring.Or(s1, s2)
		card = card + int(s3.GetCardinality())
	}
}

// go test -bench BenchmarkUnion -run -
func BenchmarkUnionGoRoar(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := goroar.New()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(uint32(r.Int31n(int32(sz))))
	}
	s2 := goroar.New()
	sz = 1000000
	initsize = 650
	for i := 0; i < initsize; i++ {
		s2.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := s1.Clone()
		s3.Or(s2) // goroar cheats here
		card = card + s3.Cardinality()
	}
}

// go test -bench BenchmarkSize -run -
func BenchmarkSizeBitset(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := bitset.New(0)
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Set(uint(r.Int31n(int32(sz))))
	}
	s2 := bitset.New(0)
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Set(uint(r.Int31n(int32(sz))))
	}
	fmt.Printf("%.1f MB ", float32(s1.BinaryStorageSize()+s2.BinaryStorageSize())/(1024.0*1024))

}

// go test -bench BenchmarkSize -run -
func BenchmarkSizeCRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := gocroaring.New()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(uint32(r.Int31n(int32(sz))))
	}
	s2 := gocroaring.New()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(uint32(r.Int31n(int32(sz))))
	}
	fmt.Printf("%.1f MB ", float32(s1.SerializedSizeInBytes()+s2.SerializedSizeInBytes())/(1024.0*1024))
}

// go test -bench BenchmarkSize -run -
func BenchmarkSizeRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := roaring.New()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(uint32(r.Int31n(int32(sz))))
	}
	s2 := roaring.New()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(uint32(r.Int31n(int32(sz))))
	}
	fmt.Printf("%.1f MB ", float32(s1.GetSerializedSizeInBytes()+s2.GetSerializedSizeInBytes())/(1024.0*1024))
}

// go test -bench BenchmarkSize -run -
func BenchmarkSizeGoroar(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := goroar.New()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(uint32(r.Int31n(int32(sz))))
	}
	s2 := goroar.New()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(uint32(r.Int31n(int32(sz))))
	}
	fmt.Printf("%.1f MB ", float32(s1.SizeInBytes()+s2.SizeInBytes())/(1024.0*1024))
}

func BenchmarkSetBitset(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	s := bitset.New(0)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Set(uint(r.Int31n(int32(sz))))
	}
}

// go test -bench BenchmarkSet -run -
func BenchmarkSetRCoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	s := gocroaring.New()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
}

// go test -bench BenchmarkSet -run -
func BenchmarkSetRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	s := roaring.New()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
}

// go test -bench BenchmarkSet -run -
func BenchmarkSetGoroar(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	s := goroar.New()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
}

func BenchmarkGetTestBitSet(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	initsize := 50000
	s := bitset.New(0)
	for i := 0; i < initsize; i++ {
		s.Set(uint(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Test(uint(r.Int31n(int32(sz))))
	}
}

// go test -bench BenchmarkGetTest -run -
func BenchmarkGetTestCRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	initsize := 50000
	s := gocroaring.New()
	for i := 0; i < initsize; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(uint32(r.Int31n(int32(sz))))
	}
}

// go test -bench BenchmarkGetTest -run -
func BenchmarkGetTestRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	initsize := 50000
	s := roaring.New()
	for i := 0; i < initsize; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(uint32(r.Int31n(int32(sz))))
	}
}

// go test -bench BenchmarkGetTest -run -
func BenchmarkGetTestGoroar(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	initsize := 50000
	s := goroar.New()
	for i := 0; i < initsize; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(uint32(r.Int31n(int32(sz))))
	}
}

func BenchmarkCountBitset(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := bitset.New(0)
	sz := 1000000
	initsize := 50000
	for i := 0; i < initsize; i++ {

		s.Set(uint(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Count()
	}
}

// go test -bench BenchmarkCount -run -
func BenchmarkCountCRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := gocroaring.New()
	sz := 1000000
	initsize := 50000
	for i := 0; i < initsize; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Cardinality()
	}
}

// go test -bench BenchmarkCount -run -
func BenchmarkCountRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := roaring.New()
	sz := 1000000
	initsize := 50000
	for i := 0; i < initsize; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.GetCardinality()
	}
}

// go test -bench BenchmarkCount -run -
func BenchmarkCountGoroar(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := goroar.New()
	sz := 1000000
	initsize := 50000
	for i := 0; i < initsize; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Cardinality()
	}
}

// go test -bench BenchmarkIterate -run -
func BenchmarkIterateBitset(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := bitset.New(0)
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s.Set(uint(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		c := uint(0)
		for i, e := s.NextSet(0); e; i, e = s.NextSet(i + 1) {
			c++
		}
	}
}

// go test -bench BenchmarkIterate -run -
func BenchmarkIterateRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := roaring.New()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		c := uint(0)
		i := s.Iterator()
		for i.HasNext() {
			i.Next()
			c++
		}
	}
}

// go test -bench BenchmarkIterate -run -
func BenchmarkIterateGoroar(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := goroar.New()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		c := uint32(0)
		for _ = range s.Iterator() {
			c++
		}
	}
}

// go test -bench BenchmarkSparseIterate -run -
func BenchmarkSparseIterateBitset(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := bitset.New(0)
	sz := 100000000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s.Set(uint(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		c := uint(0)
		for i, e := s.NextSet(0); e; i, e = s.NextSet(i + 1) {
			c++
		}
	}
}

// go test -bench BenchmarkSparseIterate -run -
func BenchmarkSparseIterateRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := roaring.New()
	sz := 100000000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		c := uint(0)
		i := s.Iterator()
		for i.HasNext() {
			i.Next()
			c++
		}
	}
}

// go test -bench BenchmarkSparseIterate -run -
func BenchmarkSparseIterateGoroar(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := goroar.New()
	sz := 100000000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		c := uint(0)
		for _ = range s.Iterator() {
			c++
		}
	}
}

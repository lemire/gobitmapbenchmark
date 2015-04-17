package roaring

import (
	"fmt"
	"github.com/fzandona/goroar"
	"github.com/tgruben/roaring"
	"github.com/willf/bitset"
	"math/rand"
	"testing"
)


// BENCHMARKS, to run them type "go test -bench Benchmark -run -"

// go test -bench BenchmarkIntersection -run -
func BenchmarkIntersectionBitset(b *testing.B) {
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
	b.StartTimer()
	card := uint(0)
	for j := 0; j < b.N; j++ {
		s3 := s1.Intersection(s2)
		card = card + s3.Count()
	}
}

// go test -bench BenchmarkIntersection -run -
func BenchmarkIntersectionRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := roaring.NewRoaringBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(int(r.Int31n(int32(sz))))
	}
	s2 := roaring.NewRoaringBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(int(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := roaring.And(s1, s2)
		card = card + s3.GetCardinality()
	}
}

// go test -bench BenchmarkIntersection -run -
func BenchmarkIntersectionGoroar(b *testing.B) {
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
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
	   	s3 := s1.Clone()
		s3.Or(s2)
		card = card + s3.Cardinality()
	}
}




// go test -bench BenchmarkIntersectionDense -run -
func BenchmarkIntersectionDenseBitset(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := bitset.New(0)
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Set(uint(r.Int31n(int32(sz))))
	}
	s2 := bitset.New(0)
	sz = 10000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Set(uint(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := uint(0)
	for j := 0; j < b.N; j++ {
		s3 := s1.Intersection(s2)
		card = card + s3.Count()
	}
}

// go test -bench BenchmarkIntersectionDense -run -
func BenchmarkIntersectionDenseRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := roaring.NewRoaringBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(int(r.Int31n(int32(sz))))
	}
	s2 := roaring.NewRoaringBitmap()
	sz = 10000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(int(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := roaring.And(s1, s2)
		card = card + s3.GetCardinality()
	}
}

// go test -bench BenchmarkIntersectionDense -run -
func BenchmarkIntersectionDenseGoroar(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := goroar.New()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(uint32(r.Int31n(int32(sz))))
	}
	s2 := goroar.New()
	sz = 10000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
	   	s3 := s1.Clone()
		s3.Or(s2)
		card = card + s3.Cardinality()
	}
}

// go test -bench BenchmarkUnion -run -
func BenchmarkUnionBitset(b *testing.B) {
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
	b.StartTimer()
	card := uint(0)
	for j := 0; j < b.N; j++ {
		s3 := s1.Union(s2)
		card = card + s3.Count()
	}
}

// go test -bench BenchmarkUnion -run -
func BenchmarkUnionRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := roaring.NewRoaringBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(int(r.Int31n(int32(sz))))
	}
	s2 := roaring.NewRoaringBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(int(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := roaring.Or(s1, s2)
		card = card + s3.GetCardinality()
	}
}

// go test -bench BenchmarkUnion -run -
func BenchmarkUnionRoaringAlt(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := roaring.NewRoaringBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(int(r.Int31n(int32(sz))))
	}
	s2 := roaring.NewRoaringBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(int(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := roaring.Or(s2, s1)
		card = card + s3.GetCardinality()
	}
}

// go test -bench BenchmarkUnion -run -
func BenchmarkUnionGoroar(b *testing.B) {
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
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := s1.Clone()
		s3.Or(s2)
		card = card + s3.Cardinality()
	}
}


// go test -bench BenchmarkUnion -run -
func BenchmarkUnionGoroarAlt(b *testing.B) {
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
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3 := s2.Clone()
		s3.Or(s1)
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
func BenchmarkSizeRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := roaring.NewRoaringBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(int(r.Int31n(int32(sz))))
	}
	s2 := roaring.NewRoaringBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(int(r.Int31n(int32(sz))))
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
func BenchmarkSetRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	s := roaring.NewRoaringBitmap()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Add(int(r.Int31n(int32(sz))))
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
func BenchmarkGetTestRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	initsize := 50000
	s := roaring.NewRoaringBitmap()
	for i := 0; i < initsize; i++ {
		s.Add(int(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(int(r.Int31n(int32(sz))))
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
func BenchmarkCountRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := roaring.NewRoaringBitmap()
	sz := 1000000
	initsize := 50000
	for i := 0; i < initsize; i++ {
		s.Add(int(r.Int31n(int32(sz))))
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
	s := roaring.NewRoaringBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s.Add(int(r.Int31n(int32(sz))))
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
	s := roaring.NewRoaringBitmap()
	sz := 100000000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s.Add(int(r.Int31n(int32(sz))))
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




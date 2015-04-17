package roaring

import (
	"math/rand"
	"testing"
    "github.com/willf/bitset"
    "github.com/tgruben/roaring"
    "github.com/fzandona/goroar"
	"fmt"
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
	s1 := NewRoaringBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(int(r.Int31n(int32(sz))))
	}
	s2 := NewRoaringBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(int(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3:= And(s1,s2)
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
		s1.Add(int(r.Int31n(int32(sz))))
	}
	s2 := NewRoaringBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(int(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3:= And(s1,s2)
		card = card + s3.GetCardinality()
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
	s1 := NewRoaringBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(int(r.Int31n(int32(sz))))
	}
	s2 := NewRoaringBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(int(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := 0
	for j := 0; j < b.N; j++ {
		s3:= Or(s1,s2)
		card = card + s3.GetCardinality()
	}
}




// go test -bench BenchmarkUnion -run -
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
	fmt.Printf("%.1f MB ",float32(s1.BinaryStorageSize()+s2.BinaryStorageSize())/(1024.0*1024))

}

// go test -bench BenchmarkUnion -run -
func BenchmarkSizeRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := NewRoaringBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(int(r.Int31n(int32(sz))))
	}
	s2 := NewRoaringBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(int(r.Int31n(int32(sz))))
	}
	fmt.Printf("%.1f MB ",float32(s1.GetSerializedSizeInBytes()+s2.GetSerializedSizeInBytes())/(1024.0*1024))

}




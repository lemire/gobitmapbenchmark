package bitmapbench

import (
	"math/rand/v2"
	"runtime"
	"testing"

	"github.com/RoaringBitmap/gocroaring"
	"github.com/RoaringBitmap/roaring/v2"
	"github.com/bits-and-blooms/bitset"
)

const (
	SZ1 = 15_000_000
	N1  = 650_000
	SZ2 = 100_000_000
	N2  = 650_000
)

func newRNG() *rand.Rand {
	return rand.New(rand.NewPCG(1, 2))
}

func fillBitset(r *rand.Rand, n, max int) *bitset.BitSet {
	s := bitset.New(uint(max))
	for range n {
		s.Set(uint(r.Int32N(int32(max))))
	}
	return s
}

func fillRoaring(r *rand.Rand, n, max int) *roaring.Bitmap {
	s := roaring.New()
	for range n {
		s.Add(uint32(r.Int32N(int32(max))))
	}
	return s
}

func fillCRoaring(r *rand.Rand, n, max int) *gocroaring.Bitmap {
	s := gocroaring.New()
	for range n {
		s.Add(uint32(r.Int32N(int32(max))))
	}
	return s
}

func BenchmarkIntersection(b *testing.B) {
	r := newRNG()
	bs1, bs2 := fillBitset(r, N1, SZ1), fillBitset(r, N2, SZ2)
	rb1, rb2 := fillRoaring(r, N1, SZ1), fillRoaring(r, N2, SZ2)
	cr1, cr2 := fillCRoaring(r, N1, SZ1), fillCRoaring(r, N2, SZ2)

	b.Run("bitset", func(b *testing.B) {
		var card uint
		for b.Loop() {
			card = bs1.Intersection(bs2).Count()
		}
		runtime.KeepAlive(card)
	})
	b.Run("roaring", func(b *testing.B) {
		var card uint64
		for b.Loop() {
			card = roaring.And(rb1, rb2).GetCardinality()
		}
		runtime.KeepAlive(card)
	})
	b.Run("gocroaring", func(b *testing.B) {
		var card uint64
		for b.Loop() {
			card = gocroaring.And(cr1, cr2).Cardinality()
		}
		runtime.KeepAlive(card)
	})
}

func BenchmarkIntersectionDense(b *testing.B) {
	r := newRNG()
	bs1, bs2 := fillBitset(r, N1, SZ1/10), fillBitset(r, N2, SZ2/10)
	rb1, rb2 := fillRoaring(r, N1, SZ1/10), fillRoaring(r, N2, SZ2/10)
	cr1, cr2 := fillCRoaring(r, N1, SZ1/10), fillCRoaring(r, N2, SZ2/10)

	b.Run("bitset", func(b *testing.B) {
		var card uint
		for b.Loop() {
			card = bs1.Intersection(bs2).Count()
		}
		runtime.KeepAlive(card)
	})
	b.Run("roaring", func(b *testing.B) {
		var card uint64
		for b.Loop() {
			card = roaring.And(rb1, rb2).GetCardinality()
		}
		runtime.KeepAlive(card)
	})
	b.Run("gocroaring", func(b *testing.B) {
		var card uint64
		for b.Loop() {
			card = gocroaring.And(cr1, cr2).Cardinality()
		}
		runtime.KeepAlive(card)
	})
}

func BenchmarkUnion(b *testing.B) {
	r := newRNG()
	bs1, bs2 := fillBitset(r, N1, SZ1), fillBitset(r, N2, SZ2)
	rb1, rb2 := fillRoaring(r, N1, SZ1), fillRoaring(r, N2, SZ2)
	cr1, cr2 := fillCRoaring(r, N1, SZ1), fillCRoaring(r, N2, SZ2)

	b.Run("bitset", func(b *testing.B) {
		var card uint
		for b.Loop() {
			card = bs1.Union(bs2).Count()
		}
		runtime.KeepAlive(card)
	})
	b.Run("roaring", func(b *testing.B) {
		var card uint64
		for b.Loop() {
			card = roaring.Or(rb1, rb2).GetCardinality()
		}
		runtime.KeepAlive(card)
	})
	b.Run("gocroaring", func(b *testing.B) {
		var card uint64
		for b.Loop() {
			card = gocroaring.Or(cr1, cr2).Cardinality()
		}
		runtime.KeepAlive(card)
	})
}

func TestSize(t *testing.T) {
	r := newRNG()
	bs1, bs2 := fillBitset(r, N1, SZ1), fillBitset(r, N2, SZ2)
	rb1, rb2 := fillRoaring(r, N1, SZ1), fillRoaring(r, N2, SZ2)
	cr1, cr2 := fillCRoaring(r, N1, SZ1), fillCRoaring(r, N2, SZ2)

	mb := func(nbytes uint64) float64 { return float64(nbytes) / (1024 * 1024) }
	t.Logf("bitset:     %.3f MB", mb(uint64(bs1.BinaryStorageSize()+bs2.BinaryStorageSize())))
	t.Logf("roaring:    %.3f MB", mb(rb1.GetSerializedSizeInBytes()+rb2.GetSerializedSizeInBytes()))
	t.Logf("gocroaring: %.3f MB", mb(uint64(cr1.SerializedSizeInBytes()+cr2.SerializedSizeInBytes())))
}

func BenchmarkSet(b *testing.B) {
	const sz = 1_000_000
	b.Run("bitset", func(b *testing.B) {
		r := newRNG()
		s := bitset.New(0)
		for b.Loop() {
			s.Set(uint(r.Int32N(int32(sz))))
		}
	})
	b.Run("roaring", func(b *testing.B) {
		r := newRNG()
		s := roaring.New()
		for b.Loop() {
			s.Add(uint32(r.Int32N(int32(sz))))
		}
	})
	b.Run("gocroaring", func(b *testing.B) {
		r := newRNG()
		s := gocroaring.New()
		for b.Loop() {
			s.Add(uint32(r.Int32N(int32(sz))))
		}
	})
}

func BenchmarkContains(b *testing.B) {
	const sz, n = 1_000_000, 50_000
	r := newRNG()
	bs := fillBitset(r, n, sz)
	rb := fillRoaring(r, n, sz)
	cr := fillCRoaring(r, n, sz)

	b.Run("bitset", func(b *testing.B) {
		r := newRNG()
		var found bool
		for b.Loop() {
			found = bs.Test(uint(r.Int32N(int32(sz))))
		}
		runtime.KeepAlive(found)
	})
	b.Run("roaring", func(b *testing.B) {
		r := newRNG()
		var found bool
		for b.Loop() {
			found = rb.Contains(uint32(r.Int32N(int32(sz))))
		}
		runtime.KeepAlive(found)
	})
	b.Run("gocroaring", func(b *testing.B) {
		r := newRNG()
		var found bool
		for b.Loop() {
			found = cr.Contains(uint32(r.Int32N(int32(sz))))
		}
		runtime.KeepAlive(found)
	})
}

func BenchmarkCount(b *testing.B) {
	const sz, n = 1_000_000, 50_000
	r := newRNG()
	bs := fillBitset(r, n, sz)
	rb := fillRoaring(r, n, sz)
	cr := fillCRoaring(r, n, sz)

	b.Run("bitset", func(b *testing.B) {
		var card uint
		for b.Loop() {
			card = bs.Count()
		}
		runtime.KeepAlive(card)
	})
	b.Run("roaring", func(b *testing.B) {
		var card uint64
		for b.Loop() {
			card = rb.GetCardinality()
		}
		runtime.KeepAlive(card)
	})
	b.Run("gocroaring", func(b *testing.B) {
		var card uint64
		for b.Loop() {
			card = cr.Cardinality()
		}
		runtime.KeepAlive(card)
	})
}

func BenchmarkIterate(b *testing.B) {
	const sz, n = 150_000, 65_000
	r := newRNG()
	bs := fillBitset(r, n, sz)
	rb := fillRoaring(r, n, sz)
	cr := fillCRoaring(r, n, sz)

	b.Run("bitset", func(b *testing.B) {
		var c uint
		for b.Loop() {
			c = 0
			for i, ok := bs.NextSet(0); ok; i, ok = bs.NextSet(i + 1) {
				c += i
			}
		}
		runtime.KeepAlive(c)
	})
	b.Run("roaring", func(b *testing.B) {
		var c uint64
		for b.Loop() {
			c = 0
			it := rb.Iterator()
			for it.HasNext() {
				c += uint64(it.Next())
			}
		}
		runtime.KeepAlive(c)
	})
	b.Run("gocroaring", func(b *testing.B) {
		var c uint64
		for b.Loop() {
			c = 0
			it := cr.Iterator()
			for it.HasNext() {
				c += uint64(it.Next())
			}
		}
		runtime.KeepAlive(c)
	})
}

func BenchmarkSparseIterate(b *testing.B) {
	const sz, n = 100_000_000, 65_000
	r := newRNG()
	bs := fillBitset(r, n, sz)
	rb := fillRoaring(r, n, sz)
	cr := fillCRoaring(r, n, sz)

	b.Run("bitset", func(b *testing.B) {
		var c uint
		for b.Loop() {
			c = 0
			for i, ok := bs.NextSet(0); ok; i, ok = bs.NextSet(i + 1) {
				c += i
			}
		}
		runtime.KeepAlive(c)
	})
	b.Run("roaring", func(b *testing.B) {
		var c uint64
		for b.Loop() {
			c = 0
			it := rb.Iterator()
			for it.HasNext() {
				c += uint64(it.Next())
			}
		}
		runtime.KeepAlive(c)
	})
	b.Run("gocroaring", func(b *testing.B) {
		var c uint64
		for b.Loop() {
			c = 0
			it := cr.Iterator()
			for it.HasNext() {
				c += uint64(it.Next())
			}
		}
		runtime.KeepAlive(c)
	})
}

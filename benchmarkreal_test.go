package bitmapbench

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/RoaringBitmap/gocroaring"
	"github.com/RoaringBitmap/roaring/v2"
)

const datadir = "realdata"

func readOneIntegerFile(path string) ([]uint32, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if i := bytes.IndexByte(dat, '\n'); i >= 0 {
		dat = dat[:i]
	}
	dat = bytes.TrimSpace(dat)
	if len(dat) == 0 {
		return []uint32{}, nil
	}
	parts := bytes.Split(dat, []byte{','})
	out := make([]uint32, 0, len(parts))
	for _, p := range parts {
		p = bytes.TrimSpace(p)
		if len(p) == 0 {
			continue
		}
		v, err := strconv.ParseUint(string(p), 10, 32)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		out = append(out, uint32(v))
	}
	return out, nil
}

func readAllIntegerFiles(dirname string) ([][]uint32, error) {
	var data [][]uint32
	err := filepath.WalkDir(dirname, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".txt" {
			return nil
		}
		one, err := readOneIntegerFile(path)
		if err != nil {
			return err
		}
		data = append(data, one)
		return nil
	})
	return data, err
}

func datasetDirs() ([]string, error) {
	entries, err := os.ReadDir(datadir)
	if err != nil {
		return nil, err
	}
	var dirs []string
	for _, e := range entries {
		if e.IsDir() {
			dirs = append(dirs, filepath.Join(datadir, e.Name()))
		}
	}
	return dirs, nil
}

func convertToRoaring(arrays [][]uint32) []*roaring.Bitmap {
	out := make([]*roaring.Bitmap, len(arrays))
	for i, a := range arrays {
		bm := roaring.BitmapOf(a...)
		bm.RunOptimize()
		out[i] = bm
	}
	return out
}

func convertToCRoaring(arrays [][]uint32) []*gocroaring.Bitmap {
	out := make([]*gocroaring.Bitmap, len(arrays))
	for i, a := range arrays {
		bm := gocroaring.New(a...)
		bm.RunOptimize()
		out[i] = bm
	}
	return out
}

func massClone(bitmaps []*roaring.Bitmap) []*roaring.Bitmap {
	out := make([]*roaring.Bitmap, len(bitmaps))
	for i, bm := range bitmaps {
		out[i] = bm.Clone()
	}
	return out
}

func massCClone(bitmaps []*gocroaring.Bitmap) []*gocroaring.Bitmap {
	out := make([]*gocroaring.Bitmap, len(bitmaps))
	for i, bm := range bitmaps {
		out[i] = bm.Clone()
	}
	return out
}

func TestSerializationRoundTrip(t *testing.T) {
	dirs, err := datasetDirs()
	if err != nil {
		t.Fatal(err)
	}
	for _, dir := range dirs {
		t.Run(filepath.Base(dir), func(t *testing.T) {
			arrays, err := readAllIntegerFiles(dir)
			if err != nil {
				t.Fatal(err)
			}
			if len(arrays) == 0 {
				t.Fatal("no arrays loaded")
			}
			bitmaps := convertToRoaring(arrays)
			cbitmaps := convertToCRoaring(arrays)
			for i := range arrays {
				buf := make([]byte, cbitmaps[i].SerializedSizeInBytes())
				if err := cbitmaps[i].Write(buf); err != nil {
					t.Fatalf("gocroaring write %d: %v", i, err)
				}
				got := roaring.New()
				if _, err := got.ReadFrom(bytes.NewReader(buf)); err != nil {
					t.Fatalf("roaring read %d: %v", i, err)
				}
				if !got.Equals(bitmaps[i]) {
					t.Fatalf("gocroaring -> roaring mismatch at %d", i)
				}

				var buffer bytes.Buffer
				if _, err := bitmaps[i].WriteTo(&buffer); err != nil {
					t.Fatalf("roaring write %d: %v", i, err)
				}
				cgot, err := gocroaring.Read(buffer.Bytes())
				if err != nil {
					t.Fatalf("gocroaring read %d: %v", i, err)
				}
				if !cgot.Equals(cbitmaps[i]) {
					t.Fatalf("roaring -> gocroaring mismatch at %d", i)
				}
			}
		})
	}
}

func BenchmarkReal(b *testing.B) {
	b.Logf("gocroaring / CRoaring %d.%d.%d", gocroaring.CRoaringMajor, gocroaring.CRoaringMinor, gocroaring.CRoaringRevision)
	dirs, err := datasetDirs()
	if err != nil {
		b.Fatal(err)
	}
	for _, dir := range dirs {
		name := filepath.Base(dir)
		arrays, err := readAllIntegerFiles(dir)
		if err != nil {
			b.Fatal(err)
		}
		howmany := len(arrays)
		if howmany == 0 {
			b.Fatalf("no arrays loaded from %s", dir)
		}
		b.Logf("read %d arrays from %s", howmany, dir)

		bitmaps := convertToRoaring(arrays)
		cbitmaps := convertToCRoaring(arrays)

		b.Run("roaring/creating/"+name, func(b *testing.B) {
			var out []*roaring.Bitmap
			for b.Loop() {
				out = convertToRoaring(arrays)
			}
			runtime.KeepAlive(out)
		})
		b.Run("gocroaring/creating/"+name, func(b *testing.B) {
			var out []*gocroaring.Bitmap
			for b.Loop() {
				out = convertToCRoaring(arrays)
			}
			runtime.KeepAlive(out)
		})

		b.Run("roaring/clone/"+name, func(b *testing.B) {
			var out []*roaring.Bitmap
			for b.Loop() {
				out = massClone(bitmaps)
			}
			runtime.KeepAlive(out)
		})
		b.Run("gocroaring/clone/"+name, func(b *testing.B) {
			var out []*gocroaring.Bitmap
			for b.Loop() {
				out = massCClone(cbitmaps)
			}
			runtime.KeepAlive(out)
		})

		b.Run("roaring/iterate/"+name, func(b *testing.B) {
			var tot uint64
			for b.Loop() {
				tot = 0
				for _, bm := range bitmaps {
					it := bm.Iterator()
					for it.HasNext() {
						tot += uint64(it.Next())
					}
				}
			}
			runtime.KeepAlive(tot)
		})
		b.Run("gocroaring/iterate/"+name, func(b *testing.B) {
			var tot uint64
			for b.Loop() {
				tot = 0
				for _, bm := range cbitmaps {
					it := bm.Iterator()
					for it.HasNext() {
						tot += uint64(it.Next())
					}
				}
			}
			runtime.KeepAlive(tot)
		})

		b.Run("roaring/toArray/"+name, func(b *testing.B) {
			var tot uint64
			for b.Loop() {
				tot = 0
				for _, bm := range bitmaps {
					tot += uint64(len(bm.ToArray()))
				}
			}
			runtime.KeepAlive(tot)
		})
		b.Run("gocroaring/toArray/"+name, func(b *testing.B) {
			var tot uint64
			for b.Loop() {
				tot = 0
				for _, bm := range cbitmaps {
					tot += uint64(len(bm.ToArray()))
				}
			}
			runtime.KeepAlive(tot)
		})

		b.Run("roaring/pairwiseAND/"+name, func(b *testing.B) {
			var tot uint64
			for b.Loop() {
				tot = 0
				for i := 0; i+1 < howmany; i++ {
					tot += roaring.And(bitmaps[i], bitmaps[i+1]).GetCardinality()
				}
			}
			runtime.KeepAlive(tot)
		})
		b.Run("gocroaring/pairwiseAND/"+name, func(b *testing.B) {
			var tot uint64
			for b.Loop() {
				tot = 0
				for i := 0; i+1 < howmany; i++ {
					tot += gocroaring.And(cbitmaps[i], cbitmaps[i+1]).Cardinality()
				}
			}
			runtime.KeepAlive(tot)
		})

		b.Run("roaring/pairwiseOR/"+name, func(b *testing.B) {
			var tot uint64
			for b.Loop() {
				tot = 0
				for i := 0; i+1 < howmany; i++ {
					tot += roaring.Or(bitmaps[i], bitmaps[i+1]).GetCardinality()
				}
			}
			runtime.KeepAlive(tot)
		})
		b.Run("gocroaring/pairwiseOR/"+name, func(b *testing.B) {
			var tot uint64
			for b.Loop() {
				tot = 0
				for i := 0; i+1 < howmany; i++ {
					tot += gocroaring.Or(cbitmaps[i], cbitmaps[i+1]).Cardinality()
				}
			}
			runtime.KeepAlive(tot)
		})

		b.Run("roaring/fastOR/"+name, func(b *testing.B) {
			var out *roaring.Bitmap
			for b.Loop() {
				out = roaring.FastOr(bitmaps...)
			}
			runtime.KeepAlive(out)
		})
		b.Run("gocroaring/fastOR/"+name, func(b *testing.B) {
			var out *gocroaring.Bitmap
			for b.Loop() {
				out = gocroaring.FastOr(cbitmaps...)
			}
			runtime.KeepAlive(out)
		})
	}
}

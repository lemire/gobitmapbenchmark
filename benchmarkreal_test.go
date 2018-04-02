package roaring

import (
	"fmt"
	//"math/rand"
	"github.com/RoaringBitmap/gocroaring"
	"github.com/RoaringBitmap/roaring"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
  "bytes"
	"testing"
	//"bufio"

)

func readOneIntegerFile(path string) ([]uint32, error) {
	var answer []uint32
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("couldn't read ", path)
		return answer, err
	}
	s := string(dat)
	idx := strings.Index(s, "\n")
	if idx != -1 {
		s = s[0:idx]
	}
	justnumbers := strings.Split(s, ",")
	howmany := len(justnumbers)
	answer = make([]uint32, howmany)
	for i := 0; i < howmany; i++ {
		var value int
		value, err = strconv.Atoi(justnumbers[i])
		if err != nil {
			fmt.Println("Not a number? ", justnumbers[i])
			break
		}
		answer[i] = uint32(value)
	}
	return answer, err
}

func readAllIntegerFiles(dirname string, extension string) ([][]uint32, error) {
	var data [][]uint32
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == extension {
			onearray, err := readOneIntegerFile(path)
			if err == nil {
				data = append(data, onearray)
			} else {
				return err
			}
		}
		return nil
	})
	return data, err
}

const datadir = "realdata"

func convertToRoaring(arrays [][]uint32) []*roaring.Bitmap {
	howmany := len(arrays)
	answer := make([]*roaring.Bitmap, howmany)
	for i := 0; i < howmany; i++ {
		answer[i] = roaring.BitmapOf(arrays[i]...)
    answer[i].RunOptimize()
	}
	return answer
}
func convertToCRoaring(arrays [][]uint32) []*gocroaring.Bitmap {
	howmany := len(arrays)
	answer := make([]*gocroaring.Bitmap, howmany)
	for i := 0; i < howmany; i++ {
		answer[i] = gocroaring.New(arrays[i]...)
    answer[i].RunOptimize()
	}
	return answer
}



func massClone(bitmaps []*roaring.Bitmap) []*roaring.Bitmap {
	answer := make([]*roaring.Bitmap, len(bitmaps))
	for i, b := range bitmaps {
		answer[i] = b.Clone()
	}
	return answer
}

func massCClone(bitmaps []*gocroaring.Bitmap) []*gocroaring.Bitmap {
	answer := make([]*gocroaring.Bitmap, len(bitmaps))
	for i, b := range bitmaps {
		answer[i] = b.Clone()
	}
	return answer
}

// go test -bench BenchmarkReal -run -
func BenchmarkReal(b *testing.B) {
	fmt.Println("gocroaring version ", gocroaring.CRoaringMajor, ".", gocroaring.CRoaringMinor, ".", gocroaring.CRoaringRevision)
	_ = filepath.Walk(datadir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && path != datadir {
			arrays, err := readAllIntegerFiles(path, ".txt")
			if err != nil {
				fmt.Println("can't benchmark ", err)
			}
			howmany := len(arrays)
			fmt.Println(" ====== ")

			fmt.Println("read ", howmany, " arrays from ", path)
			fmt.Println("")

			var bitmaps []*roaring.Bitmap
			b.Run("roaring creating "+path, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					bitmaps = convertToRoaring(arrays)
				}
			})
			var cbitmaps []*gocroaring.Bitmap
			b.Run("gocroaring creating "+path, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					cbitmaps = convertToCRoaring(arrays)
				}
			})



			for i := 0; i < howmany; i++ {
				buf := make([]byte, cbitmaps[i].SerializedSizeInBytes())
				cbitmaps[i].Write(buf) // we omit error handling
        newrb := roaring.New()
				newrb.ReadFrom(bytes.NewReader(buf))
				if !newrb.Equals(bitmaps[i]) {
					panic("bug")
				}
			}
			for i := 0; i < howmany; i++ {
				buffer := new(bytes.Buffer)
				_, err:= bitmaps[i].WriteTo(buffer)
        if err != nil {
          panic("bug")
        }
        buf := buffer.Bytes()
        newrb,_ := gocroaring.Read(buf)
				if !newrb.Equals(cbitmaps[i]) {
					panic("bug")
				}
			}
			fmt.Println("")
			for i := 0; i < howmany; i++ {
				if uint64(cbitmaps[i].SerializedSizeInBytes()) != bitmaps[i].GetSerializedSizeInBytes() {
					fmt.Println("bitmap ", i, " differ in size?")
					fmt.Println("roaring bitmap size: ", bitmaps[i].GetSerializedSizeInBytes())
					fmt.Println("croaring bitmap size: ", cbitmaps[i].SerializedSizeInBytes())
					bs := bitmaps[i].Stats()

					fmt.Println("roaring stats: cardinality = ", bs.Cardinality, " n_containers = ", bs.Containers, " n_array_containers = ", bs.ArrayContainers, " n_run_containers = ", bs.RunContainers, " n_bitset_containers = ", bs.BitmapContainers, " n_bytes_array_containers = ", bs.ArrayContainerBytes, " n_bytes_run_containers = ", bs.RunContainerBytes, " n_bytes_bitset_containers = ", bs.BitmapContainerBytes, " n_values_array_containers = ", bs.ArrayContainerValues, " n_values_run_containers = ", bs.RunContainerValues, " n_values_bitset_containers = ", bs.BitmapContainerValues)
					s := cbitmaps[i].Stats()

					fmt.Println("croaring stats: cardinality = ", s["cardinality"], " n_containers = ", s["n_containers"], " n_array_containers = ", s["n_array_containers"], " n_run_containers = ", s["n_run_containers"], " n_bitset_containers = ", s["n_bitset_containers"], " n_bytes_array_containers = ", s["n_bytes_array_containers"], " n_bytes_run_containers = ", s["n_bytes_run_containers"], " n_bytes_bitset_containers = ", s["n_bytes_bitset_containers"], " n_values_array_containers = ", s["n_values_array_containers"], " n_values_run_containers = ", s["n_values_run_containers"], " n_values_bitset_containers = ", s["n_values_bitset_containers"])

				}
			}
			b.Run("roaring clone "+path, func(b *testing.B) {

				for i := 0; i < b.N; i++ {
					_ = massClone(bitmaps)
				}
			})
			b.Run("gocroaring clone "+path, func(b *testing.B) {
				for i := 0; i < b.N; i++ {

					_ = massCClone(cbitmaps)
				}
			})
			fmt.Println("")


			fmt.Println("")

			b.Run("roaring iterate "+path, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tot := uint64(0)
					for _, b := range bitmaps {
						it := b.Iterator()
						for it.HasNext() {
							tot += uint64(it.Next())
						}
					}
				}
			})
			b.Run("gocroaring iterate "+path, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tot := uint64(0)
					for _, b := range cbitmaps {
						it := b.Iterator()
						for it.HasNext() {
							tot += uint64(it.Next())
						}
					}

				}
			})
			fmt.Println("")

			b.Run("roaring toArray "+path, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tot := uint64(0)
					for _, b := range bitmaps {
						tot += uint64(len(b.ToArray()))
					}
				}
			})
			b.Run("gocroaring toArray "+path, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tot := uint64(0)
					for _, b := range cbitmaps {
						tot += uint64(len(b.ToArray()))
					}
				}
			})
			fmt.Println("")

			b.Run("roaring pairwise AND "+path, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tot := uint64(0)
					for i := 0; i+1 < howmany; i++ {
						b := roaring.And(bitmaps[i], bitmaps[i+1])
						tot += uint64(b.GetCardinality())
					}
				}
			})
			b.Run("gocroaring pairwise AND "+path, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tot := uint64(0)
					for i := 0; i+1 < howmany; i++ {
						b := gocroaring.And(cbitmaps[i], cbitmaps[i+1])
						tot += uint64(b.Cardinality())
					}
				}
			})
			fmt.Println("")
			b.Run("roaring pairwise OR "+path, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tot := uint64(0)
					for i := 0; i+1 < howmany; i++ {
						b := roaring.Or(bitmaps[i], bitmaps[i+1])
						tot += uint64(b.GetCardinality())
					}
				}
			})
			b.Run("gocroaring pairwise OR "+path, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tot := uint64(0)
					for i := 0; i+1 < howmany; i++ {
						b := gocroaring.Or(cbitmaps[i], cbitmaps[i+1])
						tot += uint64(b.Cardinality())
					}
				}
			})
			fmt.Println("")
			b.Run("roaring fastOR "+path, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					roaring.FastOr(bitmaps...)
				}
			})
			b.Run("gocroaring fastOR "+path, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					gocroaring.FastOr(cbitmaps...)
				}
			})
			fmt.Println("")
			fmt.Println("")

		}
		return nil
	})
}

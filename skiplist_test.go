package skiplist

//TODO search && remove benchmark

import (
	"fmt"
	"math/rand"
	"testing"
)

var (
	sl            *SL
	data          []Int
	benchData1k   []Int
	benchData1kSL *SL
	benchData1m   []Int
	benchData1mSL *SL
	//benchData1g []Int too slow
)

func init() {
	for _, val := range rand.Perm(20) {
		data = append(data, Int(val))
	}
	for _, val := range rand.Perm(200) {
		data = append(data, Int(val))
	}
	for i := 0; i < 1024; i++ {
		benchData1k = append(benchData1k, Int(rand.Int31()))

	}
	for i := 0; i < 1024*1024; i++ {
		benchData1m = append(benchData1m, Int(rand.Int31()))

	}
}

type Int int

func (val Int) Less(data Lesser) bool {
	switch data.(type) {
	case Sentinel:
		return true
	}
	return val < data.(Int)
}

func (val Int) Equal(data Lesser) bool {
	switch data.(type) {
	case Sentinel:
		return false
	}
	return val == data.(Int)
}

func TestNew(t *testing.T) {
	sl = New(0.25, len(data), Sentinel{})
}

func TestInsert(t *testing.T) {
	for _, val := range data {
		sl.Insert(val)
	}
	sl.Insert(Int(999))
}

func TestFind(t *testing.T) {
	node := sl.Find(Int(999))
	fmt.Printf("node 999 %v\n", node)
	node = sl.Find(Int(0))
	fmt.Printf("node 0 %v\n", node)
}

func TestIter(t *testing.T) {
	sl.Iter(func(data Lesser) {
		fmt.Printf("%v,", data)
	})
}

func TestRemove(t *testing.T) {
	fmt.Println("===============")
	sl.Iter(func(data Lesser) {
		fmt.Printf("%v,", data)
	})
	fmt.Println("")
	sl.Remove(Int(11))
	sl.Remove(Int(199))
	sl.Remove(Int(800))
	sl.Remove(Int(999))
	sl.Iter(func(data Lesser) {
		fmt.Printf("%v,", data)
	})
	fmt.Println("")
	fmt.Println("===============")
}

func BenchmarkElem1K(b *testing.B) {
	benchData1mSL := New(0.25, len(benchData1m), Sentinel{})
	for n := 0; n < b.N; n++ {
		for _, data := range benchData1k {
			benchData1kSL.Insert(data)
		}
	}
	benchData1kSL = sl

}

func BenchmarkElem1M(b *testing.B) {
	benchData1mSL := New(0.25, len(benchData1m), Sentinel{})
	for n := 0; n < b.N; n++ {
		for _, data := range benchData1m {
			benchData1mSL.Insert(data)
		}
	}

}

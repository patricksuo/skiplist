package skiplist

//TODO indexable

import (
	"math"
	"math/rand"
)

type Lesser interface {
	Less(Lesser) bool
	Equal(Lesser) bool
}

//Sentinel is end of the skip list and always larger than other
type Sentinel struct{}

func (val Sentinel) Less(data Lesser) bool {
	return false
}

func (val Sentinel) Equal(data Lesser) bool {
	return false
}

type SL struct {
	head        *Node //keep trace of all "express lane" layers, does NOT keep any useful data
	growFactory float32
	maxHigh     int
	sentinel    Lesser
}

type Node struct {
	next []*Node
	Data Lesser
}

func New(growFactory float32, exceptedMax int, sentinel Lesser) *SL {
	maxHigh := int(math.Log2(float64(exceptedMax))) + 1
	head := &Node{Data: sentinel, next: make([]*Node, maxHigh, maxHigh)}
	for i := 0; i < maxHigh; i++ {
		head.next[i] = &Node{next: nil, Data: sentinel}
	}
	return &SL{
		maxHigh:     maxHigh,
		growFactory: growFactory,
		head:        head,
		sentinel:    sentinel,
	}
}

func (sl *SL) Insert(data Lesser) {
	//find node that node.next.data >= data for each level
	lNodes := make([]*Node, sl.maxHigh, sl.maxHigh)
	node := sl.head
	for i := sl.maxHigh - 1; i >= 0; i-- {
		for node.next[i].Data.Less(data) {
			node = node.next[i]
		}
		lNodes[i] = node
	}
	high := 1
	for high < sl.maxHigh && rand.Float32() <= sl.growFactory {
		high++
	}
	newNode := &Node{Data: data, next: make([]*Node, high, high)}
	for i := 0; i < high; i++ {
		newNode.next[i] = lNodes[i].next[i]
		lNodes[i].next[i] = newNode
	}
}

func (sl *SL) Remove(data Lesser) {
	//find first node that node.next.data >= data for each level
	lNodes := make([]*Node, sl.maxHigh, sl.maxHigh)
	node := sl.head
	for i := sl.maxHigh - 1; i >= 0; i-- {
		for node.next[i].Data.Less(data) {
			node = node.next[i]
		}
		lNodes[i] = node
	}
	for i, node := range lNodes {
		if node.next[i].Data.Equal(data) {
			node.next[i] = node.next[i].next[i]
		}
	}
}

func (sl *SL) Find(data Lesser) *Node {
	node := sl.head
	for i := sl.maxHigh - 1; i >= 0; i-- {
		for node.next[i].Data.Less(data) {
			node = node.next[i]
		}
	}
	if data.Equal(node.next[0].Data) {
		return node
	}
	return nil
}

func (sl *SL) Iter(cb func(Lesser)) {
	node := sl.head
	for len(node.next) > 0 {
		cb(node.Data)
		node = node.next[0]
	}
}

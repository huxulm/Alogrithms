// ref: https://medium.com/javarevisited/internal-working-of-hashmap-in-java-97aeac3c7beb
package hashmap

import (
	"fmt"
	"hash/fnv"
)

var defaultCapcity uint64 = 1 << 10

type node struct {
	key   interface{}
	value interface{}
	next  *node
}

// Hashmap is a golang implementation of hashmap
type Hashmap struct {
	capcity uint64
	size    uint64
	table   []*node
}

// New return new HashMap instance
func New() *Hashmap {
	return &Hashmap{
		capcity: defaultCapcity,
		table:   make([]*node, defaultCapcity),
	}
}

// Make creates a new Hashmap instance with input size and capcity
func Make(size, capcity uint64) *Hashmap {
	return &Hashmap{
		size:    size,
		capcity: capcity,
		table:   make([]*node, capcity),
	}
}

func (hm *Hashmap) Contains(key interface{}) bool {
	return hm.table[hm.hash(key)] != nil
}

// Get returns value associated with given key
func (hm *Hashmap) Get(key interface{}) interface{} {
	node := hm.getNodeByHash(hm.hash(key))

	if node != nil {
		return node.value
	}

	return nil
}

func (hm *Hashmap) Put(key, value interface{}) interface{} {
	return hm.putValue(hm.hash(key), key, value)
}

func (hm *Hashmap) putValue(hash uint64, key interface{}, value interface{}) interface{} {
	if hm.capcity == 0 {
		hm.capcity = defaultCapcity
		hm.table = make([]*node, defaultCapcity)
	}

	node := hm.getNodeByHash(hash)
	if node == nil {
		hm.table[hash] = newNode(key, value)
	} else if node.key == key {
		hm.table[hash] = newNodeWithNext(key, value, node)
		return value
	} else {
		hm.Resize()
		return hm.putValue(hash, key, value)
	}

	hm.size++

	return value
}

func (hm *Hashmap) Resize() {
	hm.capcity <<= 1
	tempTable := hm.table
	hm.table = make([]*node, hm.capcity)
	for i := 0; i < len(tempTable); i++ {
		node := tempTable[i]
		if node == nil {
			continue
		}
		hm.table[hm.hash(node.key)] = node
	}
}

func newNode(key, value interface{}) *node {
	return &node{key: key, value: value}
}

func newNodeWithNext(key, value interface{}, next *node) *node {
	return &node{key: key, value: value, next: next}
}

func (hm *Hashmap) hash(key interface{}) uint64 {
	h := fnv.New64()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))
	hashValue := h.Sum64()
	return (hm.capcity - 1) & (hashValue ^ (hashValue >> 16))
}

func (hm *Hashmap) getNodeByHash(hash uint64) *node {
	return hm.table[hash]
}

package go_bloom_filter

import (
	"github.com/spaolacci/murmur3"
	"github.com/willf/bitset"
	"sync"
)

type BloomFilter struct {
	bitMap    *bitset.BitSet
	mu        sync.RWMutex
	size      uint
	hashCount uint
}

func NewBloomFilter(size uint, hashCount uint) *BloomFilter {
	bf := BloomFilter{bitMap: bitset.New(size), size: size, hashCount: hashCount}
	return &bf
}

func (bf BloomFilter) Lookup(value string) bool {
	bf.mu.RLock()
	for i := uint(0); i < bf.hashCount; i++ {
		result := bf.hash(value) % uint64(bf.size)
		if !bf.bitMap.Test(uint(result)) {
			return false
		}
	}
	defer bf.mu.RUnlock()
	return true
}

func (bf *BloomFilter) Add(value string) {
	bf.mu.Lock()
	for i := uint(0); i < bf.hashCount; i++ {
		result := bf.hash(value) % uint64(bf.size)
		bf.bitMap.SetTo(uint(result), true)
	}
	bf.mu.Unlock()
}

func (bf BloomFilter) hash(value string) uint64 {
	m := murmur3.New64()
	_, _ = m.Write([]byte(value))
	return m.Sum64()
}

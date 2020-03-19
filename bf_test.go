package go_bloom_filter

import (
	"fmt"
	"testing"
)

func TestBloomFilter_Add(t *testing.T) {
	bf := NewBloomFilter(100000, 7)
	bf.Add("333")
	bf.Add("100000")
	bf.Add("500001")
	bf.Add("510001")
}

func TestBloomFilter_Lookup(t *testing.T) {
	bf := NewBloomFilter(100000, 7)
	bf.Add("333")
	if !bf.Lookup("333") {
		t.Fatal()
	}
	if bf.Lookup("111") {
		t.Fatal()
	}
}

func BenchmarkBloomFilter_Add(b *testing.B) {
	bf := NewBloomFilter(100000000, 7)
	for i := 0; i < b.N; i++ {
		bf.Add(fmt.Sprintf("%d", i))
	}
}

func BenchmarkBloomFilter_Lookup(b *testing.B) {
	bf := NewBloomFilter(100000000, 7)
	for i := 0; i < b.N; i++ {
		bf.Lookup(fmt.Sprintf("%d", i))
	}
}

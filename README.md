# A Simple Bloom Filter for Go

[![Build Status](https://travis-ci.org/sysatom/go-bloom-filter.svg?branch=master)](https://travis-ci.org/sysatom/go-bloom-filter)

### Install

```bash
go get github.com/sysatom/go_bloom_filter
```

### Usage

```go
package main

import (
	"fmt"
	"github.com/sysatom/go_bloom_filter"
)

func main() {
	bf := go_bloom_filter.NewBloomFilter(100000, 7)
	for i := 0; i < 5000; i++ {
		bf.Add(fmt.Sprintf("%d", i))
	}
	fmt.Println(bf.Lookup("333"))
	fmt.Println(bf.Lookup("100000"))
	fmt.Println(bf.Lookup("500001"))
	fmt.Println(bf.Lookup("510001"))
}
```

### Requirements

This project requires Go 1.14 or newer.

### License

You can find the license for this code in [the LICENSE file](LICENSE).

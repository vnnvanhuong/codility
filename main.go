package main

import (
	"fmt"
	"vnnvanhuong/codility/prefixsums"
)

func main() {
	r := prefixsums.GenomicRangeQuery("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6})
	fmt.Printf("Result: %d", r)
}

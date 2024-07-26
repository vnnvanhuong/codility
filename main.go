package main

import (
	"fmt"
	"vnnvanhuong/codility/binarysearch"
	"vnnvanhuong/codility/fibonacci"
)

func main() {
	fmt.Println("--------Fibonacci: Ladder--------")
	fibonacci.RunLadder()

	fmt.Println("--------Binary Search-----------")
	binarysearch.RunBinarySearch()
	binarysearch.RunBoards()
}

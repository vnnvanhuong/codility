package main

import (
	"fmt"
	"vnnvanhuong/codility/binarysearch"
	"vnnvanhuong/codility/fibonacci"
)

func main() {
	fmt.Println("--------Fibonacci: Ladder")
	fibonacci.RunLadder()

	fmt.Println("--------Binary Search")
	binarysearch.RunBinarySearch()
	
	fmt.Println("--------Binary Search: Boards")
	binarysearch.RunBoards()

	fmt.Println("--------Binary Search: MinMaxDivision")
	binarysearch.RunMinMaxDivision()
}

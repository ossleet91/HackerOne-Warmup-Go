package main

import (
	"fmt"
)

func printLine() {
	fmt.Printf("\n================================================================================\n")
}

func main() {
	fizzBuzz()
	printLine()

	var arrSumVals = []int32{1, 2, 3, 4, 10, 11}
	fmt.Printf("Array sum of %v = %d\n", arrSumVals, arraySum(arrSumVals))
	printLine()

	var arrBigSumVals = []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	fmt.Printf("Array big sum of %v = %d\n", arrBigSumVals, arrayBigSum(arrBigSumVals))
	printLine()

	var alice = []int32{5, 6, 7}
	var bob = []int32{3, 6, 10}
	fmt.Printf("Triplets compare of Alice %v & Bob %v is %v\n", alice, bob, tripletsCompare(alice, bob))
	printLine()
}

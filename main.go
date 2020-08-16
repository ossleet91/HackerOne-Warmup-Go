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
	fmt.Printf("Array sum of {%v} = %d\n", arrSumVals, arraySum([]int32{1, 2, 3, 4, 10, 11}))
	printLine()
}
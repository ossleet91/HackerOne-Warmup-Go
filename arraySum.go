package main

func arraySum(array []int32) int32 {
	var sum int32

	for _, num := range array {
		sum += num
	}

	return sum
}

func arrayBigSum(array []int64) int64 {
	var sum int64

	for _, num := range array {
		sum += num
	}

	return sum
}
package main

// Compare the Triplets problem
// https://www.hackerrank.com/challenges/compare-the-triplets/problem

func tripletsCompare(alice, bob []int32) []int32 {
	var a, b int32

	for i := 0; i < len(alice); i++ {
		if alice[i] > bob[i] {
			a += 1
		} else if alice[i] < bob[i] {
			b += 1
		}
	}

	return []int32{a, b}
}

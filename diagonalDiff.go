package main

// Diagonal difference problem
// https://www.hackerrank.com/challenges/diagonal-difference/problem

import (
	"math"
)

func diagonalDiff(matrix [][]int32) int32 {
	var leftD, rightD int32
	var i, j int

	outerLen := len(matrix)
	innerLen := len(matrix)
	if outerLen != innerLen {
		panic("given matrix is not a square matrix")
	}

	for i = 0; i < outerLen; i++ {
		for j = 0; j < innerLen; j++ {
			if i == j {
				leftD += matrix[i][j]
			}
		}
		rightD += matrix[i][j-1-i]
	}

	return int32(math.Abs(float64(leftD - rightD)))
}

package max_array_sum

import (
	"math"
)

// https://www.hackerrank.com/challenges/max-array-sum/problem?h_l=interview&playlist_slugs%5B%5D%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D%5B%5D=dynamic-programming
// TODO:Logic is correct, but not enough fast - need to reduce algo complexity

var sumsOfTail []int32

func Precalculate(arr []int32) {
	sumsOfTail = make([]int32, len(arr))
	for i := 0; i < len(sumsOfTail); i += 1 {
		sumsOfTail[i] = math.MinInt32
	}

	for i := len(arr) - 1; i >= 0; i-- {
		sumsOfTail[i] = arr[i]
		if i+2 < len(arr) {
			sumsOfTail[i] += sumsOfTail[i+2]
		}

	}
}

func MaxSubsetSum(arr []int32) int32 {
	Precalculate(arr)

	var max int32 = math.MinInt32
	for i := 0; i < len(arr); i += 1 {
		for j := i + 2; j < len(arr); j += 1 {
			out := arr[i] + sumsOfTail[j]
			if max < out {
				max = out
			}
		}
	}
	return max
}

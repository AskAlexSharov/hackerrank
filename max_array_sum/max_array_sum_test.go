package max_array_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNonAdjacentSubsets(t *testing.T) {
	in := []int32{3, 5, -7, 8, 10}
	//expect := [][]int32{
	//	{-2, 3},
	//	{-2, 3, 5},
	//	{-2, -4},
	//	{-2, 5},
	//	{1, -4},
	//	{1, 5},
	//	{3, 5},
	//}
	//res := GetNonAdjacentSubsets(in)
	//i := 0
	//for _, out := range res {
	//	if !reflect.DeepEqual(out, expect[i]) {
	//		t.Errorf("Expected: %+v, got: %+v\n", expect[i], out)
	//	}
	//	i += 1
	//}

	result := MaxSubsetSum(in)
	assert.Equal(t, 8, result)
}

func BenchmarkMaxSubsetSum(b *testing.B) {
	in := make([]int32, 1001)
	for i := 0; i < b.N; i++ {
		_ = MaxSubsetSum(in)
	}
}

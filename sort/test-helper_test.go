package sort

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func testHelper(t testing.TB, sort func(arr []int)) {
	for i := 0; i < 100; i++ {
		arrSize := rand.Intn(100)
		arr := make([]int, arrSize)
		cpy := make([]int, arrSize)
		for j := 0; j < arrSize; j++ {
			arr[j] = rand.Intn(100)
		}
		copy(cpy, arr)
		sort(arr)
		//check
		for j := 1; j < len(arr); j++ {
			if arr[j] < arr[j-1] {
				require.Failf(t, "arr isnt sorted", "arr%v, src%v", arr, cpy)
			}
		}
	}
}

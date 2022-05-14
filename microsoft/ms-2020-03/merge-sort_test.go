package ms_2020_03

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			in:  []int{1, 2, 3},
			out: []int{1, 2, 3},
		},
		{
			in:  []int{3, 2, 1},
			out: []int{1, 2, 3},
		},
		{
			in:  []int{1},
			out: []int{1},
		},
		{
			in:  []int{6, 3, 5, 8, 45, 2, 4, 6},
			out: []int{2, 3, 4, 5, 6, 6, 8, 45},
		},
	}

	for i, test := range tests {
		lst := SliceToList(test.in)
		lst = MergeSort(lst)
		res := ListToSlice(lst)
		assert.Equalf(t, test.out, res, "case(%d) in(%+v) sorted(%+v) must be (%+v)", i, test.in, res, test.out)
	}

	const ssize = 5e6
	arr := make([]int, ssize)
	for i := 0; i < ssize; i++ {
		arr[i] = rand.Int()
	}
	lst := SliceToList(arr)

	//startSort := time.Now()
	sort.Ints(arr)
	//elapsedSort := time.Since(startSort)

	//startMerge := time.Now()
	lst = MergeSort(lst)
	//elapsedMerge := time.Since(startMerge)

	ret := ListToSlice(lst)

	assert.Equalf(t, arr, ret, "on really big data MergeSort fails :(")
	//assert.Greaterf(t, elapsedSort, elapsedMerge, "sort.Ints(%+v) MergeSort(%+v)", elapsedSort, elapsedMerge)
}

//func cmpSlice(a, b []int) bool {
//	if len(a) != len(b) {
//		return false
//	}
//	for i := 0; i < len(a); i++ {
//		if a[i] != b[i] {
//			return false
//		}
//	}
//	return true
//}

func ListToSlice(lst *List) []int {
	res := []int{}
	for ; lst != nil; lst = lst.nxt {
		res = append(res, lst.key)
	}
	return res
}

func SliceToList(arr []int) *List {
	if len(arr) == 0 {
		return nil
	}
	res := &List{
		key: arr[0],
	}
	ptr := res
	for i := 1; i < len(arr); i++ {
		ptr.nxt = &List{
			key: arr[i],
		}
		ptr = ptr.nxt
	}
	return res
}

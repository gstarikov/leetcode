package recursion_i

func getRow(rowIndex int) []int {
	switch rowIndex {
	case 0:
		return []int{1}
	case 1:
		return []int{1, 1}
	}
	curr := make([]int, 2, rowIndex+1)
	curr[0] = 1
	curr[1] = 1
	for i := 2; i <= rowIndex; i++ {
		prev := curr[0]
		for j := 1; j < len(curr); j++ {
			tmp := curr[j]
			curr[j] = prev + curr[j]
			prev = tmp
		}
		curr = append(curr, 1)
	}
	return curr
}

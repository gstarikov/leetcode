package sort

func InsertSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	for j := 0; j < len(arr); j++ {
		minIdx := j
		minVal := arr[minIdx] //only to a bit speed up (it will be placed in register)
		for i := j + 1; i < len(arr); i++ {
			val := arr[i]
			if val < minVal {
				minIdx = i
				minVal = val
			}
		}
		if j != minIdx {
			swap(j, minIdx, arr)
		}
	}
}

func swap(i, j int, arr []int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

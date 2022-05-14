package sort

func QuickSort(arr []int) {
	quickSortImp(0, len(arr)-1, arr)
}

func quickSortImp(i, j int, arr []int) {
	if i >= j {
		return
	}
	p := xoara(arr, i, j)
	quickSortImp(i, p, arr)
	quickSortImp(p+1, j, arr)
}

func xoara(arr []int, i, j int) int {
	p := (i + j) / 2
	pVal := arr[p]
	for { //xoara method
		for ; arr[i] < pVal; i++ {
		}
		for ; arr[j] > pVal; j-- {
		}
		if i >= j {
			return j
		}
		swap(i, j, arr)
		i++
		j--
	}
}

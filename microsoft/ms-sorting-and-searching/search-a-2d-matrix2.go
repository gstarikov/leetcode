package ms_sorting_and_searching

/*
Write an efficient algorithm that searches
for a value target in an m x n integer matrix matrix.
This matrix has the following properties:
    Integers in each row are sorted in ascending from left to right.
    Integers in each column are sorted in ascending from top to bottom.

Constraints:
    m == matrix.length
    n == matrix[i].length
    1 <= n, m <= 300
    -10^9 <= matrix[i][j] <= 10^9
    All the integers in each row are sorted in ascending order.
    All the integers in each column are sorted in ascending order.
    -10^9 <= target <= 10^9
*/
func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	// идея простая - у нас есть диагональ, левее и выше которой элементы меньше
	// сначала делаем бинарный поиск по диагонали, а потом ищем в строке и в столбце, тоже бинарно
	// и ещё 2 поиска вышу и правее (?)

	n, m := len(matrix), len(matrix[0])
	//ln := n*m - 1
	//border cases
	if matrix[0][0] > target || matrix[n-1][m-1] < target {
		return false
	}
	//scan diag
	var x, y int
	for idx := 1; idx < n; idx++ {
		y = idx
		x = y * m / n
		v := matrix[y][x]
		if v > target {
			break
		}
	}

	//target is located on
	// matrix[y][0:x]
	// matrix[y-1][x-1:]
	// matrix[0:y][x]
	// matrix[y-1:][x-1]
	var res = false

	res = res || helperBinSearch(0, m-1, target, func(idx int) int {
		return matrix[y][idx]
	})
	if y > 0 {
		res = res || helperBinSearch(0, m-1, target, func(idx int) int {
			return matrix[y-1][idx]
		})
	}

	res = res || helperBinSearch(0, n-1, target, func(idx int) int {
		return matrix[idx][x]
	})
	if x > 0 {
		res = res || helperBinSearch(0, n-1, target, func(idx int) int {
			return matrix[idx][x-1]
		})
	}
	return res
}
func searchMatrix2v2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	//ln := n*m - 1
	//border cases
	if matrix[0][0] > target || matrix[n-1][m-1] < target {
		return false
	}
	//scan lines or rows
	var maxIdx, otherDim, maxOtherDim int
	var getElem getElem
	switch {
	case n > m:
		maxIdx = n - 1
		maxOtherDim = m - 1
		getElem = func(idx int) int {
			return matrix[idx][otherDim]
		}
	default:
		maxIdx = m - 1
		maxOtherDim = n - 1
		getElem = func(idx int) int {
			return matrix[otherDim][idx]
		}
	}
	for otherDim = 0; otherDim <= maxOtherDim; otherDim++ {
		if helperBinSearch(0, maxIdx, target, getElem) {
			return true
		}
	}
	return false
}

type getElem func(idx int) int //in will be fast enough because inlining

func helperBinSearch(idxStart, idxEnd, target int, getElem getElem) bool {
	if idxStart < 0 {
		idxStart = 0
	}
	if idxEnd < 0 {
		idxEnd = 0
	}
	for idxStart <= idxEnd {
		mid := (idxStart + idxEnd) / 2
		v := getElem(mid)
		switch {
		case target == v:
			return true
		case target > v:
			idxStart = mid + 1
		case target < v:
			idxEnd = mid - 1
		default:
			panic("impossible case")
		}
	}
	return false
}

func searchMatrix2hint(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	//border cases
	if matrix[0][0] > target || matrix[n-1][m-1] < target {
		return false
	}

	x, y := 0, n-1
	for x < m && y >= 0 {
		v := matrix[y][x]
		switch {
		case v > target:
			y--
		case v < target:
			x++
		case v == target:
			return true
		default:
			panic("impossible case")
		}
	}
	return false
}

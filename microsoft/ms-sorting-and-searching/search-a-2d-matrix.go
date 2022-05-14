package ms_sorting_and_searching

/*
Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:
    Integers in each row are sorted from left to right.
    The first integer of each row is greater than the last integer of the previous row.

Constraints:
    m == matrix.length
    n == matrix[i].length
    1 <= m, n <= 100
    -10^4 <= matrix[i][j], target <= 10^4
*/

func searchMatrix(matrix [][]int, target int) bool {
	//at first find line
	//then search pos in line
	t, b := 0, len(matrix)-1
	var line int
	switch { //border cases
	case matrix[t][0] > target,
		matrix[b][len(matrix[b])-1] < target:
		return false
	case matrix[b][0] == target:
		return true
	case matrix[b][0] < target:
		line = b
	default:
		b-- //skip last line because we already check it
	L:
		for {
			mid := (t + b) / 2
			me0, me1 := matrix[mid][0], matrix[mid+1][0]
			switch {
			case target == me0,
				target == me1:
				return true
			case target < me0:
				b = mid - 1
			case target > me0 && target < me1:
				line = mid
				break L
			case target > me1:
				t = mid + 1
			case t >= b:
				return false
			}
		}
	}

	//work on line
	lval := matrix[line]
	l, r := 0, len(lval)-1
	for l <= r {
		mid := (l + r) / 2
		v := lval[mid]
		switch {
		case target == v:
			return true
		case target < v:
			r = mid - 1
		case target > v:
			l = mid + 1
		}
	}
	return false
}

func searchMatrixHint(matrix [][]int, target int) bool {
	switch {
	case len(matrix) == 0, len(matrix[0]) == 0:
		return false
	}
	m, n := len(matrix), len(matrix[0])

	l, r := 0, n*m-1

	for l <= r {
		mid := (l + r) / 2
		midVal := matrix[mid/n][mid%n]
		switch {
		case target == midVal:
			return true
		case target < midVal:
			r = mid - 1
		case target > midVal:
			l = mid + 1
		}
	}
	return false
}

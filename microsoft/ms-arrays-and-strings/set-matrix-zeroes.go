package ms_arrays_and_strings

/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
You must do it in place.

Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

Constraints:

    m == matrix.length
    n == matrix[0].length
    1 <= m, n <= 200
    -2^31 <= matrix[i][j] <= 2^31 - 1



Follow up:

    A straightforward solution using O(mn) space is probably a bad idea.
    A simple improvement uses O(m + n) space, but still not the best solution.
    Could you devise a constant space solution?

*/

func SetZeroes(matrix [][]int) {
	// scan matrix setting flags(0) and skipping rows/ cols with flag
	cols := len(matrix)
	rows := len(matrix[0])

	firstColZero := false
	for j := 0; j < cols; j++ {
		if matrix[j][0] == 0 {
			firstColZero = true
			break
		}
	}
	firstRowZero := false
	for j := 0; j < rows; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}

	//scan and tag
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	//set zeroes on rows
	for i := 1; i < rows; i++ { //skip first row because there is flags for cols
		if matrix[0][i] == 0 {
			for j := 0; j < cols; j++ {
				matrix[j][i] = 0
			}
		}
	}
	//set zeroes on cols
	for i := 1; i < cols; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < rows; j++ {
				matrix[i][j] = 0
			}
		}

	}
	//set zeroes on first row
	if firstColZero {
		for j := 0; j < cols; j++ {
			matrix[j][0] = 0
		}
	}
	if firstRowZero {
		for j := 0; firstRowZero && j < rows; j++ {
			matrix[0][j] = 0
		}
	}
}

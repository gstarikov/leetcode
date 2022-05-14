package ms_arrays_and_strings

/*
Given an m x n matrix, return all elements of the matrix in spiral order.

Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

Constraints:

    m == matrix.length
    n == matrix[i].length
    1 <= m, n <= 10
    -100 <= matrix[i][j] <= 100

*/

func SpiralOrder(matrix [][]int) []int {
	// это ж блин змейка!!! которая всегда поворачивает направо
	path := len(matrix) * len(matrix[0])
	ret := make([]int, path)
	var i, j int
	iStep, jStep := 0, 1
	const visited = 666
	for pos := 0; pos < path; pos++ { //its simple that check on every turn zeroes is around or not
		ret[pos] = matrix[i][j]
		matrix[i][j] = visited
		nextI, nextJ := i+iStep, j+jStep
		//check for turn
		if nextI == len(matrix) || // эти проверки срабатывают только в первый раз, а потом по виртуальным стенам идём
			nextJ == len(matrix[0]) || // наверное стоит сделать "двигающиеся стены
			nextI < 0 || // и исходя из направления движения проверять только нужный указатель
			nextJ < 0 ||
			matrix[nextI][nextJ] == visited {
			//its time to turn right
			switch {
			case iStep == 0 && jStep == 1: //right
				iStep, jStep = 1, 0
			case iStep == 1 && jStep == 0: //down
				iStep, jStep = 0, -1
			case iStep == 0 && jStep == -1: //left
				iStep, jStep = -1, 0
			case iStep == -1 && jStep == 0: //up
				iStep, jStep = 0, 1
			}
		}
		//do real step
		i, j = i+iStep, j+jStep
	}
	return ret
}

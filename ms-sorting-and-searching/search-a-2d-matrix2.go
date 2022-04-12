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
	// идея простая - у нас есть диагональ, левее и выше которой элементы меньше
	// сначала делаем бинарный поиск по диагонали, а потом ищем в строке и в столбце, тоже бинарно

}

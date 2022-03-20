package ms_trees_and_graphs

/*
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.

Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

Constraints:
    m == grid.length
    n == grid[i].length
    1 <= m, n <= 300
    grid[i][j] is '0' or '1'.
*/

func NumIslands(grid [][]byte) int {
	//scan whole map
	// and on every 1 do
	//  add 1 to island counter
	//  recursive go throuth all 1 points and set it to 0

	var ret int
	for i, row := range grid {
		for j, v := range row {
			if v == islandPoint {
				ret++
				doClear(i, j, grid)
			}
		}
	}
	return ret
}

const waterPoint = '0'
const islandPoint = '1'

func doClear(i, j int, grid [][]byte) {
	//do recursive besause we need to have ability to return
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == waterPoint {
		return
	}
	grid[i][j] = waterPoint
	doClear(i-1, j, grid) //left
	doClear(i+1, j, grid) //right
	doClear(i, j+1, grid) //down
	doClear(i, j-1, grid) //up
}

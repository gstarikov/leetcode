package array_and_string

/*
Given an m x n matrix mat, return an array of all the elements of the array in a diagonal order.

Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,4,7,5,3,6,8,9]

123
456
789

Example 2:

Input: mat = [[1,2],[3,4]]
Output: [1,2,3,4]

*/

func findDiagonalOrder(mat [][]int) []int {
	my, mx := len(mat), len(mat[0])
	ret := make([]int, mx*my)
	retPos := 0
	dir := false
	// three cases (like loop enrolling to speedup)
	// 1) 123 - move over x<=>y diag (1,2,3 lines) x,y:=0...min(mx,my)
	//    230
	//    300
	//    000
	//    000
	// 2) 000 - (mx < my, reverse in case of mx > my) move over x,y:=mx...0, my-mx...my
	//    004
	//    045
	//    450
	//    500
	// 3) 000 - (mx <= my) move over x,y:=my-mx..mx, y:=mx...my
	//    000
	//    000
	//    006
	//    067

	addElem := func(x, y int) { //it must be inlined by compiler
		ret[retPos] = mat[y][x]
		retPos++
	}

	//border cases - to speed up while time in some cases
	switch {
	case mx == 0 || my == 0:
		return ret
	case mx == 1:
		for i := 0; i < my; i++ {
			addElem(0, i)
		}
		return ret
	case my == 1:
		for i := 0; i < mx; i++ {
			addElem(i, 0)
		}
		return ret
	}

	//ToDo: replace traverse to -- && ++ operations instead of calculations

	//case 1
	m := helperMin(mx, my)
	for i := 0; i < m; i++ {
		dir = !dir
		if dir { //like loop enrolling
			for j := 0; j <= i; j++ {
				x, y := j, i-j
				addElem(x, y)
			}
		} else {
			for j := i; j >= 0; j-- {
				x, y := j, i-j
				addElem(x, y)
			}
		}
	}

	//case 2
	switch {
	case mx < my:
		for i := mx; i < my; i++ {
			dir = !dir
			if dir {
				for j := 0; j < mx; j++ {
					x, y := j, i-j
					addElem(x, y)
				}
			} else {
				for j := mx; j >= 0; j-- {
					x, y := j, i-j
					addElem(x, y)
				}
			}
		}
	case mx > my:
		for i := my; i < mx; i++ {
			dir = !dir
			if dir {
				for j := my - 1; j >= 0; j-- {
					x, y := i-j, j
					addElem(x, y)
				}
			} else {
				for j := 0; j < my; j++ {
					x, y := i-j, j
					addElem(x, y)
				}
			}
		}
	}

	//case 3
	switch {
	case mx < my:
		for i := my - mx + 1; i < my; i++ {
			dir = !dir
			if dir {
				for j := i; j < my; j++ {
					x, y := mx-(j-i)-1, j
					addElem(x, y)
				}
			} else {
				for j := i; j < my; j++ {
					x, y := mx-(my-j), my-(j-i)-1
					addElem(x, y)
				}
			}
		}
	case mx == my:
		for i := 1; i < mx; i++ {
			dir = !dir
			if dir {
				for j := i; j < mx; j++ {
					x, y := j, mx-j //something wrong
					addElem(x, y)
				}
			} else {
				for j := i; j < mx; j++ {
					x, y := mx-j, j
					addElem(x, y)
				}
			}
		}
	case mx > my:
		for i := mx - my + 1; i < mx; i++ {
			dir = !dir
			if dir {
				for j := i; j < mx; j++ {
					x, y := j, my-(j-i)-1
					addElem(x, y)
				}
			} else {
				for j := i; j < mx; j++ {
					x, y := mx-(j-i)-1, my-(mx-j)
					addElem(x, y)
				}
			}
		}

	}

	return ret
}

func helperMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func helperMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package ms_other

import "sort"

/*
A city's skyline is the outer contour of the silhouette formed by all the buildings in that city when viewed from a distance.
Given the locations and heights of all the buildings, return the skyline formed by these buildings collectively.

The geometric information of each building is given in the array buildings where buildings[i] = [lefti, righti, heighti]:

    lefti is the x coordinate of the left edge of the ith building.
    righti is the x coordinate of the right edge of the ith building.
    heighti is the height of the ith building.

You may assume all buildings are perfect rectangles grounded on an absolutely flat surface at height 0.

The skyline should be represented as a list of "key points" sorted by their x-coordinate in the form [[x1,y1],[x2,y2],...].
Each key point is the left endpoint of some horizontal segment in the skyline except the last point in the list,
which always has a y-coordinate 0 and is used to mark the skyline's termination where the rightmost building ends.
Any ground between the leftmost and rightmost buildings should be part of the skyline's contour.

Note: There must be no consecutive horizontal lines of equal height in the output skyline.
For instance, [...,[2 3],[4 5],[7 5],[11 5],[12 7],...] is not acceptable;
the three lines of height 5 should be merged into one in the final output as such: [...,[2 3],[4 5],[12 7],...]

Constraints:

    1 <= buildings.length <= 10^4
    0 <= lefti < righti <= 2^31 - 1
    1 <= heighti <= 2^31 - 1
    buildings is sorted by lefti in non-decreasing order.

*/

func getSkylineVector(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return nil
	}

	//main idea - work with points and insert new points to return vector
	//cons - its slow because slice copying. dual linked list will be better but ill lose binary search in that case...
	var ret [][]int
	for _, building := range buildings {
		l := building[0]
		r := building[1]
		h := building[2]

		//find left point
		//iterate till right point
		//insert turn points as necessary

		lPos := sort.Search(len(ret), func(i int) bool {
			return ret[i][0] >= l
		})

		//do merge insert
		//take prev point, take cur point
		var prevH int
		switch {
		case lPos >= len(ret): // add points to end
			ret = append(ret, []int{l, h})
			ret = append(ret, []int{r, 0})
			continue
		case ret[lPos][0] != l: //insert before
			switch {
			case lPos == 0: //no prev value
				ret = append(ret, []int{0, 0})
				copy(ret[lPos+1:], ret[lPos:])
				ret[lPos] = []int{l, h}
				prevH = h
			case ret[lPos-1][1] < h: //insert new point
				ret = append(ret, []int{0, 0})
				copy(ret[lPos+1:], ret[lPos:])
				prevH = ret[lPos-1][1]
			default: //do nothing
				prevH = h
			}
		default: //the same point
			prevH = ret[lPos][1]
		}

		helperMerge(lPos, r, h, prevH, ret[lPos:])
		panic("not implemented")

	}
	return ret
}

func helperMerge(lPos, r, h, prevH int, data [][]int) [][]int {

	//iterate till r and insert/remove points. there is no prev point....
	var i int
	for i = lPos; i < len(data) && data[i][1] <= r; i++ {
		//rx := ret[i][0]
		rh := data[i][1]
		switch {
		case rh <= h && prevH < h: //remove point
			copy(data[i+1:], data[i:])
		case rh <= h && prevH >= h: //replace point
			data[i][1] = h
		case rh > h: //up point
			prevH = rh
		}
	}
	//last point
	switch {
	case i >= len(data):
		data = append(data, []int{r, 0})
	case data[i][0] > r && data[i][1] < h: //insert new point
		data = append(data, []int{0, 0})
		copy(data[i+1:], data[i:])
		data[i][0] = r
		data[i][1] = h
	}
	return data
}

func getSkylineBitMap(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return nil
	}
	//two stages:
	// 1 - scan input data and fill height list
	// 2 - optimize height list and build xy representation
	// it isnt optimal because stores empty values. its more effective to store as vector

	var height []int // index = x pos, val = height

	//stage 1
	for _, building := range buildings {
		l := building[0]
		r := building[1]
		h := building[2]

		//grow up list
		if len(height) <= r {
			height = append(height, make([]int, r-len(height))...)
		}

		//fill height list
		for j := l; j < r; j++ {
			if height[j] < h {
				height[j] = h
			}
		}
	}

	//stage 2
	var ret [][]int
	var lastH int
	for x, h := range height {
		if lastH != h { //new point detected
			ret = append(ret, []int{x, h})
			lastH = h
		}
	}
	ret = append(ret, []int{len(height), 0})

	return ret
}

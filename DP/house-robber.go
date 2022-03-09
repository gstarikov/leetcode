package DP

func HouseRobber(houses []int) RobVec {
	rv0 := houseRobber(houses, 0)
	rv1 := houseRobber(houses, 1)
	if rv1.sum > rv0.sum && len(rv1.sequence) > 0 {
		return rv1
	}
	return rv0
}

type RobVec struct {
	sequence []int
	sum      int
}

func houseRobber(houses []int, start int) RobVec {
	// check for border conditions
	switch {
	case start >= len(houses): // out of array
		return RobVec{}
	case start+1 == len(houses): // last elem, so return it
		return RobVec{
			sequence: []int{start},
			sum:      houses[start],
		}
	}

	res := [2]RobVec{}
	retIxd := 0
	for i := range res {
		hr := houseRobber(houses, start+2+i)
		res[i].sequence = append([]int{start}, hr.sequence...) // too many memcpy. its better to append to the end and sort of finish
		res[i].sum = houses[start] + hr.sum

		if res[i].sum > res[retIxd].sum {
			retIxd = i
		}
	}

	return res[retIxd]
}

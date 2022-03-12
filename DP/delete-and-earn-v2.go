package DP

func DeleteAndEarnV2(nums []int) int {
	de := prepare(nums)
	return de.delAndEarn(de.data)
}

type delAndEarn struct {
	data map[int]int
}

func prepare(nums []int) delAndEarn {
	ret := delAndEarn{
		data: make(map[int]int),
	}
	for _, v := range nums {
		ret.put(v)
	}
	return ret
}

func (e *delAndEarn) put(v int) {
	cnt := e.data[v]
	e.data[v] = cnt + 1
}

func (e *delAndEarn) delAndEarn(d map[int]int) int {
	//iterate
	mx := 0
	for k := range d {
		var earn int
		//make a copy
		cp := copyMap(d)
		// add key * value to val
		earn = cp[k] * k
		// remove elems
		delete(cp, k)
		delete(cp, k-1)
		delete(cp, k+1)
		// and go recursion
		earn += e.delAndEarn(cp)

		if earn > mx {
			mx = earn
		}
	}
	return mx
}

func copyMap(d map[int]int) map[int]int { // that is optimized by compiler, to leave it as-is
	ret := make(map[int]int)
	for k, v := range d {
		ret[k] = v
	}
	return ret
}

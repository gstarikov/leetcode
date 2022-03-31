package ms_2020_03

import "fmt"

func CountOfTriplets(peoples []int, t int) int {
	var ret int
	targetAge := t

	m := map[int]int{}
	for _, v := range peoples {
		cnt := m[v]
		m[v] = cnt + 1
	}

	for age1 := 1; age1 < targetAge; age1++ {
		//take first person
		cnt1 := m[age1]
		if cnt1 == 0 {
			continue
		}
		if cnt1 >= 3 && age1*3 == targetAge {
			ret += 1 * (cnt1 - 1) * (cnt1 - 2) //todo: number of combination
		}
		for age2 := age1 + 1; age2 < targetAge; age2++ {
			cnt2 := m[age2]
			if cnt2 == 0 {
				continue
			}
			age3 := targetAge - age1 - age2
			if age3 < age2 {
				continue
			}
			cnt3 := m[age3]
			if cnt3 != 0 {
				var inc int
				if age3 == age2 {
					inc = cnt1 * (cnt2 - 1)
				} else {
					inc = cnt1 * cnt2 * cnt3
				}
				fmt.Printf("peoples[%+v], age1(%d->%d) age2(%d->%d) age3(%d->%d) decision (%d += %d)\n",
					peoples,
					age1, cnt1, age2, cnt2, age3, cnt3,
					ret, inc)

				ret += inc
			}
		}
	}
	return ret
}

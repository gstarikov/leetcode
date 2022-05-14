package ms_2020_03

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
			ret += calcTripletsCombinations(3, cnt1)
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
					if cnt2 >= 2 {
						inc = cnt1 * calcTripletsCombinations(2, cnt2) //number of combinations
					}
				} else {
					inc = cnt1 * cnt2 * cnt3
				}
				//fmt.Printf("peoples[%+v], age1(%d->%d) age2(%d->%d) age3(%d->%d) decision (%d += %d)\n",
				//	peoples,
				//	age1, cnt1, age2, cnt2, age3, cnt3,
				//	ret, inc)

				ret += inc
			}
		}
	}
	return ret
}

/*
	1 -> 0
	11 -> 0
	111 -> 1
	1111 -> 1110 | 1101 | 1011 | 0111 -> 4
	11111 ->	11100 | 11010 | 11001 |
				10110 | 10101 | 10011 |
				01110 | 01101 | 01011 |
				00111 -> 10
	n - total
	k - subset
	count = (n+k-1)! / (k!*(n-1)!
*/
func calcTripletsCombinations(k, n int) int {
	//cnt1 := n-k
	//cnt2 :=
	if n == k {
		return 1
	}
	p, s := factorial(n), (factorial(k) * factorial(n-k))
	//fmt.Printf("k(%d) n(%d) %d/%d", k, n, p, s)
	return p / s
}

//var preCalcOnce = sync.Once{}
//var PreCalc = map[int]int
func factorial(k int) int { //todo caching
	if k == 0 {
		return 0
	}
	//preCalc.Do()
	ret := 1
	for i := 2; i <= k; i++ {
		ret *= i
	}
	return ret
}

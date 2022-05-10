package ms_other

import (
	"sort"
)

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

//переделать на потоковую версию, работающую с командами "дом появился, дом закончился"

func getSkylineTry3(buildings [][]int) [][]int {
	var pq pQueue
	var ret [][]int
	var i int //building index
	for {
		//get building & first end of building
		var l int
		okBuilding := i < len(buildings)
		if okBuilding {
			building := buildings[i]
			l = building[0]
		}
		pnt, okShadow := pq.Get()
		if !okShadow && !okBuilding { //done - no more buildings or shadows
			break
		}
		//check who is nearest - new building or end of building
		switch {
		case !okShadow && okBuilding: //there is no one building end, so add end of building and start point
			i = proceedSameStart(buildings, i, 0, &pq, &ret, false)
		case okShadow && !okBuilding: //no new buildings
			pq.Remove()
			pntNxt, ok := pq.Get()
			vl := []int{pnt.x, 0}
			if ok {
				vl[1] = pntNxt.h
			} //last point H will be zero
			ret = append(ret, vl)
			//there is both start and shadow of building
		case l == pnt.x: //building starts on place of end of other building
			pq.Remove()
			i = proceedSameStart(buildings, i, pnt.h, &pq, &ret, true)

		case l < pnt.x: //there is some building
			i = proceedSameStart(buildings, i, pnt.h, &pq, &ret, false)
		case l > pnt.x: //shadow ends before new building begins
			//dont process new building
			pq.Remove()
			pntNxt, ok := pq.Get()
			if ok {
				ret = append(ret, []int{pnt.x, pntNxt.h})
			} else {
				ret = append(ret, []int{pnt.x, 0})
			}
		}
	}
	return ret
}

func proceedSameStart(buildings [][]int, pos int, curH int, pq *pQueue, r *[][]int, diffH bool) int {
	l := buildings[pos][0]
	pq.Add(point{
		x: buildings[pos][1],
		h: buildings[pos][2],
	})
	maxH := buildings[pos][2]

	for pos++; pos < len(buildings) && buildings[pos][0] == l; pos++ {
		x := buildings[pos][1]
		h := buildings[pos][2]
		pq.Add(point{
			x: x,
			h: h,
		})
		if maxH < h {
			maxH = h
		}
	}

	switch {
	case !diffH && curH < maxH:
		*r = append(*r, []int{l, maxH})
	case diffH:
		pntNxt, _ := pq.Get()
		if curH != pntNxt.h { //могут так сложиться звёзды, что новая точка на той же высоте как закончившаяся
			*r = append(*r, []int{l, pntNxt.h})
		}
	}

	return pos
}

type point struct {
	x, h int
}

//задачи очереди:
// хранить пары Х и высота
// обеспечить учёт множественных пар (добавить и удалить)
// получение следующей Х точки
// получение текущей высоты
type pQueue struct {
	//optimal structure will be RB\AVL tree
	//but ill will using ordered slice
	queu []point
}

//type pQueueVal struct {
//	point
//}

/*
варианты:
- Х совпадает. устанавшиваем Н как у самого высокого.
- Х не совпадает. добавляем.
пост обработка:
- проверяем всех предыдущих и устанавливаем Н как у самого высокого
- убираем дубли по высоте
*/
func (p *pQueue) Add(in point) {
	pos := sort.Search(len(p.queu), func(i int) bool {
		return p.queu[i].x >= in.x
	})

	//1 step
	switch {
	case pos < len(p.queu) && p.queu[pos].x == in.x:
		if p.queu[pos].h < in.h {
			p.queu[pos].h = in.h
		}
	case pos < len(p.queu) && p.queu[pos].h >= in.h: //do nothing
		return
	default:
		p.queu = append(p.queu, point{})
		copy(p.queu[pos+1:], p.queu[pos:])
		p.queu[pos] = in
	}
	//check prev & update prev
	//binary search !!! because p.queu[:pos-1] must be in decreasing H order
	//lPos := pos
	//я ищу первый элемент который меньше либо равен моему.
	//слева все больше, с точки и далее равны либо меньше
	lPos := sort.Search(pos, func(i int) bool {
		return p.queu[i].h <= in.h
	})
	//remove shaded buildings
	switch {
	case lPos != pos: //found something
		copy(p.queu[lPos:], p.queu[pos:])
		dec := pos - lPos
		p.queu = p.queu[:len(p.queu)-dec]
	}
}

func (p pQueue) Get() (point, bool) {
	if len(p.queu) == 0 {
		return point{}, false
	}
	return p.queu[0], true
}

func (p *pQueue) Remove() {
	if len(p.queu) == 0 {
		return
	}
	//p.queu = p.queu[1:]
	copy(p.queu, p.queu[1:])
	p.queu = p.queu[:len(p.queu)-1]
}

func getSkylineVector(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return nil
	}

	//main idea - work with points and insert new points to return vector
	//cons - its slow because slice copying. dual linked list will be better but ill lose binary search in that case...
	var ret []point
	for _, building := range buildings {
		l := building[0]
		r := building[1]
		h := building[2]

		//find left point
		//iterate till right point
		//insert turn points as necessary

		lPos := sort.Search(len(ret), func(i int) bool {
			return ret[i].x >= l
		})

		//do merge insert
		//take prev point, take cur point
		var prevH int
		switch {
		case lPos >= len(ret): // add points to end
			ret = append(ret, point{x: l, h: h})
			ret = append(ret, point{x: r, h: 0})
			continue
		case ret[lPos].x != l: //insert before
			switch {
			case lPos == 0: //no prev value
				ret = append(ret, point{x: 0, h: 0})
				copy(ret[lPos+1:], ret[lPos:])
				ret[lPos] = point{x: l, h: h}
				prevH = h
			case ret[lPos-1].h < h: //insert new point
				ret = append(ret, point{0, 0})
				copy(ret[lPos+1:], ret[lPos:])
				prevH = ret[lPos-1].h
			default: //do nothing
				prevH = h
			}
		default: //the same point
			prevH = ret[lPos].h
		}

		helperMerge(lPos, l, r, h, prevH, ret[lPos:])
		panic("not implemented")

	}
	panic("not implemented")
	return [][]int{}
}

/*
3 группы ситуаций:
1 : первая точка
1.1 первая точка не совпадает ни с одной из существующих. решаем что делать:
	предыдущая ниже нашей=>вставляем новую точку
	предыдущая выше или равна нашей=>ничего не делаем
1.2 первая совпадает с существующей
	текущая выше нашей=>ничего не делаем
	текущая ниже либо равна нашей, а наша равна предыдущей=>удаляем текущую
	в остальных случаях (текущая ниже нашей, а наша ниже\выше предыдущей)=>поднимаем до текущей

2: точки между границами интервала (такой же алгоритм как в 1.2)

3: последняя точка
3.1 не попадает ни на одну из существующих
	отличается от предыдущей и отличается от следующей=>вставляем нашу точку
	меньше предыдущей=>ничего не делаем
	больше предыдущей и совпадает со следующей=>заменяем следующую
3.2 попадает на существующую- мы должны упасть до нуля. ничего не делаем
*/

func helperMerge(lPos, l, r, h, prevH int, data []point) []point {
	if len(data) < 2 {
		panic("invalid data")
	}

	switch {
	case data[lPos].x != l && prevH < h: //1.1 кейс
		data = append(data, point{})
		copy(data[lPos+1:], data[lPos:])
		data[lPos].x = l
		data[lPos].h = h
		prevH = h
		lPos++
	case data[lPos].x == l:
		prevH = data[lPos].h
		if prevH < h {
			data[lPos].h = h
		}
		lPos++
	}

	//iterate till r and insert/remove points. there is no prev point....
	var lastH int
	var i int
	for i = lPos; i < len(data)-1 && data[i].h < r; i++ {
		lastH = data[i].h

		rh := data[i].h
		switch {
		case rh > h: //do nothing
			prevH = rh
		case rh <= h && h == prevH: //remove dub
			copy(data[i:], data[i+1:])
			data = data[:len(data)-1]
		default: //up point
			data[i].h = h
			prevH = h
		}
	}
	//last point
	switch {
	case i >= len(data)-1: //assume it is zero
		last := &data[len(data)-1]
		switch {
		case last.x == r: //do nothing
		case last.x > r:
			//изучаем какая предыдущая
			//либо предыдущая выше нашей-ничего не делаем
			//либе предыдущая такая же как наша... о как... а мы же потерями информацию - чья это была предыдущая. то ли мы подняли, то ли...
			//нам нужна информация по владельцу предыдущей точки - мы или не мы. если не мы, то надо знать какая была высота..
			//зчит нельзя в цикле идти до последней
			data = append(data, point{})
			copy(data[i+1:], data[lPos:])
			last.x = r
			last.h = lastH
		case last.x < r:
			last.h = h
			data = append(data, point{x: r, h: 0})
		}

	case data[i].x == r: //do nothing
	case data[i].x > r:
		data = append(data, point{})
		copy(data[i+1:], data[i:])
		data[i].x = r
		data[i].h = lastH
	case data[i].x < r:
		panic("impossible case")
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

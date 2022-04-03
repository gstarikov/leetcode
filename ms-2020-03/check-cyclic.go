package ms_2020_03

type pack struct {
	id  string
	dep []string
}

// print true in case of cyclic
func CheckCyc(packs []pack) bool {
	m := map[string]*pack{}
	for i := range packs {
		_, ok := m[packs[i].id]
		if ok {
			panic("package list has dublicate entry") // todo: something strange...
		}
		m[packs[i].id] = &packs[i]
	}

	checkedMap := map[string]struct{}{}
	for _, v := range packs {
		if cyclic := helperCheckDepList(
			map[string]struct{}{},
			checkedMap,
			m,
			[]string{v.id},
		); cyclic {
			return true
		}
	}
	return false
}

func helperCheckDepList(
	rootMap map[string]struct{},
	checkedMap map[string]struct{},
	packAccessMap map[string]*pack,
	depList []string,
) bool {
	for _, v := range depList {
		//it is good?
		if _, ok := checkedMap[v]; ok {
			continue
		}
		//it is cyclic
		if _, ok := rootMap[v]; ok {
			return true
		}
		//do check
		elem, ok := packAccessMap[v]
		if !ok {
			continue // skip unknown dependencies. alternatove - do panic
		}
		//clone root map (its very memory expensive, but alternative - work with sorted array who is cpu expensive)
		newRoot := map[string]struct{}{}
		for k := range rootMap {
			newRoot[k] = struct{}{}
		}
		newRoot[elem.id] = struct{}{} //mark it as one of the roots

		deps := elem.dep
		check := helperCheckDepList(newRoot, checkedMap, packAccessMap, deps)
		if check == false {
			checkedMap[v] = struct{}{} //mark as non cyclic
		} else {
			return true
		}
	}
	return false // non cyclic
}

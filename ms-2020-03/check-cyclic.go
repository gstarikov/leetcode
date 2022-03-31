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
			continue // todo: something strange...
		}
		m[packs[i].id] = &packs[i]
	}

	for _, v := range packs {
		// internal
		//visitMap := map[string]struct{}{}
		//deepCheck(visitMap, m, )
		for _, d := range v.dep {
			if d == v.id {
				continue
			}
			//find v.dep[i] -> paks
			deps, ok := m[d]
			if !ok {
				//its strange - internal dep
				continue
			}

			for _, dd := range deps.dep {
				if dd == v.id {
					//its cyclic
					return true
				}
			}
		}
	}
	return false
}

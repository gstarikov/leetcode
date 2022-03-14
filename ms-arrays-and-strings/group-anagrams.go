package ms_arrays_and_strings

import "sort"

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
*/

func GroupAnagrams(strs []string) [][]string {
	hash := map[string][]string{}
	for _, v := range strs {
		t := v //there is copy
		tmp := []byte(t)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})

		vec := hash[string(tmp)]
		vec = append(vec, v)
		hash[string(tmp)] = vec
	}

	var ret [][]string
	for _, v := range hash { //golang's hash are to slow in case of iterating :(
		ret = append(ret, v)
	}
	return ret
}

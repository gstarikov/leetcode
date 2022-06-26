package array_and_string

func strStr(haystack string, needle string) int {
	switch {
	case len(needle) == 0:
		return 0
	case len(needle) > len(haystack):
		return -1
	}
	for i, c := range haystack {
		if needle[0] == byte(c) {
			var j int
			for j = 1; j < len(needle) && i+j < len(haystack) && needle[j] == haystack[i+j]; j++ {
			}
			if j == len(needle) {
				return i
			}
		}
	}
	return -1
}

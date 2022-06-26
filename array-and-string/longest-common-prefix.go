package array_and_string

func longestCommonPrefix(strs []string) string {
	switch {
	case len(strs) == 0:
		return ""
	case len(strs[0]) == 0:
		return ""
	}
	var pos int
L:
	for pos = 0; pos < len(strs[0]); pos++ {
		c := strs[0][pos]
		for i := 1; i < len(strs); i++ {
			if pos >= len(strs[i]) || c != strs[i][pos] {
				break L
			}
		}
	}
	return strs[0][:pos]
}

package ms_arrays_and_strings

func ReverseWords(s string) string {
	buf := make([]byte, 0, len(s))

	i, l := len(s)-1, 0
	for i >= 0 {
		//skip spaces
		for ; i >= 0 && s[i] == ' '; i-- {
		}
		if i < 0 {
			break
		}
		l = i
		//search word begin
		for ; i >= 0 && s[i] != ' '; i-- {
		}
		// between i & l are word
		buf = append(buf, s[i+1:l+1]...)
		buf = append(buf, ' ')
	}

	return string(buf[:len(buf)-1])
}

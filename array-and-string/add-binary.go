package array_and_string

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		t := a
		a = b
		b = t
	}
	res := make([]byte, len(a)+1)
	var co int
	lb := len(b)
	la := len(a)
	lr := len(res)
	for i := lb; i != 0; i-- {
		lb--
		la--
		lr--
		bt := int(b[lb]) - '0'
		at := int(a[la]) - '0'
		sm := at + bt + co
		sr := sm % 2
		co = sm >> 1
		res[lr] = byte(sr + '0')
	}
	for i := la - lb; i != 0; i-- {
		la--
		lr--
		at := int(a[la]) - '0'
		sm := at + co
		sr := sm % 2
		co = sm >> 1
		res[lr] = byte(sr + '0')
	}
	if co != 0 {
		res[0] = '1'
		return string(res)
	} else {
		return string(res[1:])
	}
}

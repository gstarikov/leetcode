package generic

type genStruct[T any] struct {
	ptr *T
}

func testGetStruct() {
	gs := genStruct[string]{}
	_ = len(*gs.ptr)
}

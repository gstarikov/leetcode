package ms_design

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

type LRUCacheOp int

const (
	opGet LRUCacheOp = iota
	opPut LRUCacheOp = iota
)

type testOp struct {
	op       LRUCacheOp
	key, val int
	ln       int
}

func TestLRUCache(t *testing.T) {
	const capacity = 2
	tests := []struct {
		capacity int
		ops      []testOp
	}{
		{ //0
			capacity: 2,
			ops: []testOp{
				{op: opGet, key: 1, val: -1, ln: 0}, // 0
				{op: opPut, key: 1, val: 11, ln: 1}, // 1
				{op: opGet, key: 1, val: 11, ln: 1}, // 2
				{op: opPut, key: 2, val: 22, ln: 2}, // 3
				{op: opGet, key: 1, val: 11, ln: 2}, // 4
				{op: opGet, key: 2, val: 22, ln: 2}, // 5
				{op: opGet, key: 3, val: -1, ln: 2}, // 6
				{op: opPut, key: 3, val: 33, ln: 2}, // 7
				{op: opGet, key: 1, val: -1, ln: 2}, // 8
				{op: opGet, key: 2, val: 22, ln: 2}, // 9
				{op: opGet, key: 3, val: 33, ln: 2}, //10
				{op: opPut, key: 3, val: 44, ln: 2}, //11
				{op: opGet, key: 3, val: 44, ln: 2}, //12
			},
		},
		{ //1
			capacity: 2,
			ops: []testOp{
				{op: opPut, key: 2, val: 1, ln: 1},  // 0
				{op: opPut, key: 3, val: 2, ln: 2},  // 1
				{op: opGet, key: 3, val: 2, ln: 2},  // 2
				{op: opGet, key: 2, val: 1, ln: 2},  // 3
				{op: opPut, key: 4, val: 3, ln: 2},  // 4
				{op: opGet, key: 2, val: 1, ln: 2},  // 5
				{op: opGet, key: 3, val: -1, ln: 2}, // 6
				{op: opGet, key: 4, val: 3, ln: 2},  // 7
			},
		},
	}

	for i, test := range tests {
		lru := Constructor(test.capacity)
		for j, op := range test.ops {
			switch op.op {
			case opGet:
				val := lru.Get(op.key)
				assert.Equalf(t, op.val, val, "case(%d->%d)", i, j)
			case opPut:
				lru.Put(op.key, op.val)
			default:
				panic("impossible case")
			}
			assert.Equalf(t, op.ln, lru.lruLen, "case(%d->%d)", i, j)
			assert.Equalf(t, op.ln, len(lru.cache), "case(%d->%d)", i, j)
		}

	}
}

func BenchmarkLRUCache(b *testing.B) {
	const capacity = 2
	const maxKey = 100
	const maxVal = 100
	lru := Constructor(capacity)
	rnd := rand.New(rand.NewSource(666))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		switch rnd.Int63n(2) {
		case 0:
			key := rnd.Int63n(maxKey)
			_ = lru.Get(int(key))
		case 1:
			key := rnd.Int63n(maxKey)
			val := rnd.Int63n(maxVal)
			lru.Put(int(key), int(val))
		}
	}
}

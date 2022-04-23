package ms_design

import "math"

/*
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

    LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
    int get(int key) Return the value of the key if the key exists, otherwise return -1.
    void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.

The functions get and put must each run in O(1) average time complexity.

Constraints:

    1 <= capacity <= 3000
    0 <= key <= 10^4
    0 <= value <= 10^5
    At most 2 * 10^5 calls will be made to get and put.

*/

type keyType uint16
type valType uint16

//type idxType uint16

type LRUCache struct {
	//data
	cache map[keyType]cacheData //key - int, data - int
	//lru check
	lruHeadPtr     *lruList
	lruLen, lruCap int
}

type cacheData struct {
	data   valType
	lruPtr *lruList //idx to lruList
}

type lruList struct {
	key        keyType  //key here is duplicated to allow remove from map in case of eviction
	prev, next *lruList //dual linked to speedup remove
}

func Constructor(capacity int) LRUCache {
	if capacity >= math.MaxUint16 || capacity <= 0 {
		panic("cache wrong size. too big or too small")
	}
	return LRUCache{
		cache:      make(map[keyType]cacheData),
		lruHeadPtr: nil,
		lruLen:     0,
		lruCap:     capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	v, found := this.cache[keyType(key)]
	if !found {
		return -1
	}
	if this.lruHeadPtr != v.lruPtr { //move to top only if it isnt at top
		//found, so move val to top and return
		elemPos := v.lruPtr
		//remove from list
		elemPos.prev.next = elemPos.next
		elemPos.next.prev = elemPos.prev
		//insert at top
		oldHead := this.lruHeadPtr
		elemPos.next = oldHead
		elemPos.prev = oldHead.prev
		oldHead.prev.next = elemPos
		oldHead.prev = elemPos
		//set up new head
		this.lruHeadPtr = elemPos
	}

	// return data :)
	return int(v.data)
}

func (this *LRUCache) Put(key int, value int) {
	v := this.Get(key) //check and move to top
	if v != -1 {       //found, so replace value
		cacheElem := this.cache[keyType(key)]
		cacheElem.data = valType(value)
		this.cache[keyType(key)] = cacheElem
		return
	}
	//not found, must insert. so reuse existing or "allocate" new
	var newElem *lruList
	switch {
	case this.lruLen == this.lruCap: //fully filled
		//do eviction
		// we reuse last elem to avoid unnecessary allocations
		newElem = this.lruHeadPtr.prev
		//remove from map
		delete(this.cache, newElem.key)
	case this.lruHeadPtr == nil: //empty
		//allocate new entry and place as head
		newElem = &lruList{ //replace with pool ?
			key: keyType(key),
		}
		newElem.next = newElem
		newElem.prev = newElem
		this.lruLen++
	default:
		//allocate new entry and place as head
		newElem = &lruList{ //replace with pool ?
			key:  keyType(key),
			prev: this.lruHeadPtr.prev,
			next: this.lruHeadPtr,
		}
		this.lruLen++
		//relink prev
		newElem.prev.next = newElem
		//relink head
		this.lruHeadPtr.prev = newElem
	}

	// set head
	this.lruHeadPtr = newElem
	// fill elem
	newElem.key = keyType(key)
	//add to map
	this.cache[newElem.key] = cacheData{
		data:   valType(value),
		lruPtr: newElem,
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

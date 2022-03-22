package backtraking

import (
	"sort"
)

/*
Given an m x n board of characters and a list of strings words, return all words on the board.
Each word must be constructed from letters of sequentially adjacent cells,
where adjacent cells are horizontally or vertically neighboring.
The same letter cell may not be used more than once in a word.

Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]],
words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]

Input: board = [["a","b"],["c","d"]],
words = ["abcb"]
Output: []

Constraints:
    m == board.length
    n == board[i].length
    1 <= m, n <= 12
    board[i][j] is a lowercase English letter.
    1 <= words.length <= 3 * 104
    1 <= words[i].length <= 10
    words[i] consists of lowercase English letters.
    All the strings of words are unique.

*/

func FindWords(board [][]byte, words []string) []string {
	prefixTree := Constructor()

	// build prefix tree
	for _, word := range words { //take a word
		prefixTree.Insert(word)
	}

	ret := checkBoard(board, prefixTree)
	return ret
}

func cloneBoard(board [][]byte) [][]byte {
	ret := make([][]byte, len(board))
	for i := range ret {
		ret[i] = append(ret[i], board[i]...) //it is highly optimized in compiler
	}
	return ret
}

func checkBoard(board [][]byte, words Trie) []string {
	//take first symbol and try to find it in board
	mi := len(board)
	mj := len(board[0])
	ret := []string{}
	for i := 0; i < mi; i++ {
		for j := 0; j < mj; j++ {
			//do recursive search
			arr := words.checkRecursive(board, i, j, "", words.node)
			ret = append(ret, arr...)
		}
	}
	return ret
}

func (this *Trie) checkRecursive(board [][]byte, i, j int, prefix string, node *PrefixTreeNode) (ret []string) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return nil
	}

	symbol := board[i][j]
	prefix += string(symbol)

	finish, next := getNextNode(symbol, node)
	if finish {
		ret = append(ret, prefix)
		// remove word from tree
		this.Remove(prefix)
	}
	if next == nil {
		return
	}

	//clone board
	//board = cloneBoard(board) //very expensive. but i dont know how to reduce cost
	const visitMark = 0
	board[i][j] = visitMark

	ret = append(ret, this.checkRecursive(board, i-1, j, prefix, next)...)
	ret = append(ret, this.checkRecursive(board, i+1, j, prefix, next)...)
	ret = append(ret, this.checkRecursive(board, i, j-1, prefix, next)...)
	ret = append(ret, this.checkRecursive(board, i, j+1, prefix, next)...)

	board[i][j] = symbol //restore original symbol

	return ret
}

// todo: do refactor - DRY

// compact prefix tree is bad because we going symbol by symbol
type PrefixTreeNode struct {
	elems []PrefixTreeElem // для ускорения надо бы сортировать и искать уже в отсортированном и т.п. вставка поэлементно идёт, то просто вставлять на нужное место
}

type PrefixTreeElem struct {
	prefix byte //high bit is used as word end tag
	node   *PrefixTreeNode
}

const completeWordTag byte = 1 << 7

type Trie struct {
	node *PrefixTreeNode
}

func Constructor() Trie {
	return Trie{node: &PrefixTreeNode{}}
}

func (this *Trie) Insert(word string) {
	//do iterative search & insert
	node := this.node
	for len(word) > 0 {
		char := word[0]
		word = word[1:]
		pos := sort.Search(len(node.elems), func(i int) bool {
			return (node.elems[i].prefix & ^completeWordTag) >= char
		})

		switch {
		case pos == len(node.elems), //not found at all
			(node.elems[pos].prefix & ^completeWordTag) != char: //not found, but there is greater elem
			elem := PrefixTreeElem{
				prefix: char,
				node:   nil,
			}
			if len(word) > 0 {
				elem.node = &PrefixTreeNode{}
			} else { //tag word end
				elem.prefix |= completeWordTag
			}
			if pos == len(node.elems) {
				node.elems = append(node.elems, elem)
			} else {
				node.elems = append(node.elems, PrefixTreeElem{}) //append space
				copy(node.elems[pos+1:], node.elems[pos:])
				node.elems[pos] = elem
			}
		case len(word) == 0: //we found elem, and it is the last one
			//tag word end (without additional checks because check cost equal to simple set
			node.elems[pos].prefix |= completeWordTag
		case node.elems[pos].node == nil: //we found elem, but he isnt last, so add node
			node.elems[pos].node = &PrefixTreeNode{}
		}
		node = node.elems[pos].node
	}
}

func (this *Trie) Remove(word string) bool {
	node := this.node
	var pos int
	var last *PrefixTreeElem
	for i := range word {
		char := word[i]
		pos = sort.Search(len(node.elems), func(i int) bool {
			return (node.elems[i].prefix & ^completeWordTag) >= char
		})
		switch {
		case pos == len(node.elems), //not found at all
			(node.elems[pos].prefix & ^completeWordTag) != char:
			return false
		default:
			last = &node.elems[pos]
			node = node.elems[pos].node
		}
	}
	last.prefix &= ^completeWordTag

	// real remove isnt necessary because our code is one shot. its not necessary to free memory
	//two cases
	// word is prefix for other words => so we remove only tag
	// word is not prefix, so we remove
	/*	type idxPos struct {
			idx int
			ptr *PrefixTreeNode
		}
		stack := make([]idxPos, 0, len(word))
		node := this.node
		for i := range word {
			char := word[i]
			pos := sort.Search(len(node.elems), func(i int) bool {
				return (node.elems[i].prefix & ^completeWordTag) >= char
			})
			switch {
			case pos == len(node.elems), //not found at all
				(node.elems[pos].prefix & ^completeWordTag) != char:
				return false
			default:
				stack = append(stack, idxPos{
					idx: pos,
					ptr: node,
				})
				node = node.elems[pos].node
			}
		}
		mayDelete := node == nil
		if !mayDelete { //only remove tag
			last := stack[len(stack)-1]
			last.ptr.elems[last.idx].prefix &= ^completeWordTag
		} else { //remove elems
			//remove tag from last
			{
				elem := stack[len(stack)-1]
				elem.ptr.elems[elem.idx].prefix &= ^completeWordTag
			}

			for i := len(stack) - 1; i >= 0; i-- {
				elem := stack[i]
				//clean empty sub-arrays
				if elem.ptr.elems[elem.idx].node != nil &&
					elem.ptr.elems[elem.idx].node.elems != nil &&
					len(elem.ptr.elems[elem.idx].node.elems) == 0 {
					elem.ptr.elems[elem.idx].node.elems = nil
				}
				switch {
				case len(elem.ptr.elems) == 0:
				case elem.ptr.elems[elem.idx].node != nil && len(elem.ptr.elems[elem.idx].node.elems) > 0: //its is path to other words
					return true //nothing to remove its time to stop
				case elem.ptr.elems[elem.idx].prefix&completeWordTag != 0: //time to stop because other word ends here
					return true
				case elem.idx == 0: //first one
					elem.ptr.elems = elem.ptr.elems[1:]
				case elem.idx == len(elem.ptr.elems)-1: //last one
					elem.ptr.elems = elem.ptr.elems[0:elem.idx]
				default: //somevere in the middle
					copy(elem.ptr.elems[elem.idx:], elem.ptr.elems[elem.idx-1:len(elem.ptr.elems)-1])
					elem.ptr.elems = elem.ptr.elems[0 : len(elem.ptr.elems)-1]
				}
			}
		}*/
	return true
}

func getNextNode(char byte, node *PrefixTreeNode) (wordComplete bool, next *PrefixTreeNode) { //DRY
	pos := sort.Search(len(node.elems), func(i int) bool {
		return (node.elems[i].prefix & ^completeWordTag) >= char
	})
	switch {
	case pos == len(node.elems), //not found at all
		(node.elems[pos].prefix & ^completeWordTag) != char:
		return false, nil
	default:
		wordComplete = (node.elems[pos].prefix & completeWordTag) == completeWordTag
		next = node.elems[pos].node
		return
	}
}

func (this *Trie) getNode(word string) (ok bool, lastNode PrefixTreeElem) {
	node := this.node
	for len(word) > 0 {
		char := word[0]
		word = word[1:]
		pos := sort.Search(len(node.elems), func(i int) bool {
			return (node.elems[i].prefix & ^completeWordTag) >= char
		})
		switch {
		case pos == len(node.elems), //not found at all
			(node.elems[pos].prefix & ^completeWordTag) != char:
			return false, PrefixTreeElem{}
		case len(word) == 0: //we found elem, and it is the last one
			return true, node.elems[pos]
		case node.elems[pos].node != nil: //it isnt end
			node = node.elems[pos].node
		default:
			//next node is nill, but word isnt ended
			return false, node.elems[pos]

		}
	}
	panic("impossible case")
}

func (this *Trie) Search(word string) bool {
	ok, elem := this.getNode(word)
	return ok && (elem.prefix&completeWordTag) != 0
}

func (this *Trie) StartsWith(prefix string) bool {
	ok, _ := this.getNode(prefix)
	return ok
}

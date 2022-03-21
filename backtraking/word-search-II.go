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
	ret := []string{}
	for _, word := range words { //take a word
		//check word and add to ret in case of there is on board
		if checkBoard(board, word) {
			ret = append(ret, word)
		}
	}
	return ret
}

func cloneBoard(board [][]byte) [][]byte {
	ret := make([][]byte, len(board))
	for i := range ret {
		ret[i] = append(ret[i], board[i]...) //it is highly optimized in compiler
	}
	return ret
}

func checkBoard(board [][]byte, words string) bool {
	//take first symbol and try to find it in board
	mi := len(board)
	mj := len(board[0])
	symbol := words[0]
	for i := 0; i < mi; i++ {
		for j := 0; j < mj; j++ {
			if board[i][j] == symbol {
				//do recursive search
				ok := checkRecursive(board, i, j, words)
				if ok {
					return true //there is no requirement to report how many instances of one word are on board
				}
			}
		}
	}
	return false
}

func checkRecursive(board [][]byte, i, j int, chunk string) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != chunk[0] {
		return false
	}
	if len(chunk) == 1 { //do in this manner to avoid excessive recursive call (params pass & etc)
		return true
	}
	//clone board
	board = cloneBoard(board) //very expensive. but i dont know how to reduce cost

	chunk = chunk[1:]
	const visitMark = 0
	board[i][j] = visitMark
	ret := checkRecursive(board, i-1, j, chunk) ||
		checkRecursive(board, i+1, j, chunk) ||
		checkRecursive(board, i, j-1, chunk) ||
		checkRecursive(board, i, j+1, chunk)
	return ret
}

// compact prefix tree is bad because we going symbol by symbol
type PrefixTreeNode struct {
	elems []PrefixTreeElem // для ускорения надо бы сортировать и искать уже в отсортированном и т.п. вставка поэлементно идёт, то просто вставлять на нужное место
}

type PrefixTreeElem struct {
	prefix byte
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

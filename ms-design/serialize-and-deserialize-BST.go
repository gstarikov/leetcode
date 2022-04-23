package ms_design

import (
	"bytes"
	"strconv"
)

/*
Serialization is converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer,
or transmitted across a network connection link to be reconstructed later in the same or another computer environment.
Design an algorithm to serialize and deserialize a binary search tree.
There is no restriction on how your serialization/deserialization algorithm should work.
You need to ensure that a binary search tree can be serialized to a string, and this string can be deserialized to the original tree structure.
The encoded string should be as compact as possible.

Example 1:

Input: root = [2,1,3]
Output: [2,1,3]

Example 2:

Input: root = []
Output: []

Constraints:
    The number of nodes in the tree is in the range [0, 10^4].
    0 <= Node.val <= 10^4
    The input tree is guaranteed to be a binary search tree.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func BSTConstructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	//return string(buildString(root))
	ret := []byte{}
	buildBinString(root, &ret)
	return string(ret)
}

const sep = ';'
const null = "null"

const lBit = uint16(1 << 15)
const rBit = uint16(1 << 14)
const dataMask = ^(lBit | rBit)

func decodeUint16(v uint16) (l, r bool, val int) {
	l = (v & lBit) != 0
	r = (v & rBit) != 0
	val = int(v & dataMask)
	return
}

func codeUint16(root *TreeNode) uint16 {
	var ret uint16
	if root.Left != nil {
		ret |= lBit
	}
	if root.Right != nil {
		ret |= rBit
	}
	ret |= uint16(root.Val) & dataMask //forcefly drop bits
	return ret
}

func convUint16ToBytes(v uint16) []byte {
	return []byte{
		byte(v & 0xFF),
		byte((v & 0xFF00) >> 8),
	}
}

func decodeBytesToUint16(str []byte) (v uint16) {
	//if len(str) < 2 {
	//	panic("bytes is too short")
	//}
	v = uint16(str[0]) | (uint16(str[1]) << 8)
	return v
}

func buildString(root *TreeNode) []byte {
	if root == nil {
		return []byte(null + string(sep))
	}
	str := []byte(strconv.Itoa(root.Val) + string(sep))
	str = append(str, buildString(root.Left)...)
	str = append(str, buildString(root.Right)...)
	return str
}

func buildBinString(root *TreeNode, str *[]byte) {
	if root == nil {
		return
	}
	*str = append(*str, convUint16ToBytes(codeUint16(root))...)
	buildBinString(root.Left, str)
	buildBinString(root.Right, str)
	return
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	dp := []byte(data)
	root := decodeBinString(&dp)
	return root
}

func decodeString(data []byte) (*TreeNode, []byte) {
	idx := bytes.IndexByte(data, sep)
	if idx == -1 {
		return nil, nil
	}
	token := data[:idx]
	data = data[idx+1:]
	if bytes.Equal(token, []byte(null)) {
		return nil, data
	}
	n, err := strconv.Atoi(string(token))
	if err != nil {
		panic("bad format")
	}
	root := &TreeNode{
		Val:   n,
		Left:  nil,
		Right: nil,
	}
	root.Left, data = decodeString(data)
	root.Right, data = decodeString(data)
	return root, data
}

func decodeBinString(data *[]byte) *TreeNode {
	if len(*data) == 0 {
		return nil
	}

	root := &TreeNode{}
	var v uint16
	v = decodeBytesToUint16(*data)
	*data = (*data)[2:]
	l, r, val := decodeUint16(v)
	root.Val = val
	if l {
		root.Left = decodeBinString(data)
	}
	if r {
		root.Right = decodeBinString(data)
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */

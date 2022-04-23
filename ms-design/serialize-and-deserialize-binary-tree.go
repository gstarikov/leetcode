package ms_design

/*
Constraints:

    The number of nodes in the tree is in the range [0, 10^4].
    -1000 <= Node.val <= 1000

the only one difference from prev (BST) is negative values

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

const lBit = uint16(1 << 15)
const rBit = uint16(1 << 14)
const sBit = uint16(1 << 13)
const dataMask = ^(lBit | rBit | sBit)

func decodeUint16(v uint16) (l, r bool, val int) {
	l = (v & lBit) != 0
	r = (v & rBit) != 0
	s := (v & sBit) != 0
	val = int(v & dataMask)
	if s {
		val *= -1
	}
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
	v := root.Val
	if v < 0 {
		v *= -1
		ret |= sBit
	}
	ret |= uint16(v) & dataMask //forcefly drop bits
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

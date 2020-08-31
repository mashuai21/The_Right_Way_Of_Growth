package main

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
	Size int
}

func NewBinaryTree() *BinaryTree {
	bst := &BinaryTree{
		Root: nil,
		Size: 0,
	}
	return bst
}

func (bst *BinaryTree) GetSize() int {
	return bst.Size
}

func (bst *BinaryTree) IsEmpty() bool {
	return bst.Size == 0
}

func (bst *BinaryTree) Add(data int) {
	bst.Root = bst.add(bst.Root, data)
}

func (bst *BinaryTree) add(n *Node, data int) *Node {
	if n == nil { // 根节点
		bst.Size++
		return &Node{
			Data:  data,
			Left:  nil,
			Right: nil,
		}
	} else {
		if data < n.Data {
			n.Left = bst.add(n.Left, data)
		} else {
			n.Right = bst.add(n.Right, data)
		}
	}
	return n
}

func (bst *BinaryTree) IsIn(data int) bool {
	return bst.isin(bst.Root, data)
}

func (bst *BinaryTree) isin(n *Node, data int) bool {
	if n == nil {
		return false // 空树
	}

	if data == n.Data {
		return true
	} else if data < n.Data {
		return bst.isin(n.Left, data)

	} else {
		return bst.isin(n.Right, data)
	}
}


func (bst *BinaryTree) FindMax() int {
	if bst.Size == 0 {
		panic("空树")
	}
	return bst.findmax(bst.Root).Data
}

func (bst *BinaryTree) findmax(n *Node) *Node {
	if n.Right == nil {
		return n
	} else {
		return bst.findmax(n.Right)
	}
}
func (bst *BinaryTree) FindMix() int {
	if bst.Size == 0 {
		panic("空树")
	}
	return bst.findmix(bst.Root).Data
}

func (bst *BinaryTree) findmix(n *Node) *Node {
	if n.Left == nil {
		return n
	} else {
		return bst.findmax(n.Left)
	}
}


package BinaryTree

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil, nil}
}

func (tree *BinaryTree) InsertItem(value int) {
	if tree.root == nil {
		tree.root = NewNode(value)
		return
	}
	current := tree.root
	for {
		if value > current.value {
			if current.right == nil {
				current.right = NewNode(value)
				return
			}
			current = current.right
		} else {
			if current.left == nil {
				current.left = NewNode(value)
				return
			}
			current = current.left
		}
	}
}

func (tree *BinaryTree) SearchItem(value int) (*Node, bool) {
	if tree.root == nil {
		return nil, false
	}
	current := tree.root
	for current != nil {
		if value == current.value {
			return current, true
		} else if value > current.value {
			current = current.right
		} else {
			current = current.left
		}
	}
	return nil, false
}

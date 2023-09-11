package bst

type BSTNode struct {
	Key   int
	Val   interface{}
	left  *BSTNode
	right *BSTNode
}

func NewBSTNode(key int, val interface{}) *BSTNode {
	return &BSTNode{
		Key:   key,
		Val:   val,
		left:  nil,
		right: nil,
	}
}

func (n *BSTNode) do(do func(*BSTNode)) {
	if n != nil {
		n.left.do(do)
		do(n)
		n.right.do(do)
	}
}

type BST struct {
	root *BSTNode
}

func (b *BST) Search(key int) (n *BSTNode, found bool) {
	for current := b.root; current != nil; {
		if current.Key == key {
			return current, true
		} else if key < current.Key {
			current = current.left
		} else {
			current = current.right
		}
	}
	return nil, false
}

func (b *BST) Insert(n *BSTNode) {
	// super is the pointer of the pointer
	// that points to the new node
	super := &b.root
	for *super != nil {
		if n.Key < (*super).Key {
			super = &(*super).left
		} else {
			super = &(*super).right
		}
	}
	*super = n
}

func (b *BST) Insert2(n *BSTNode) {
	var parent *BSTNode
	current := b.root
	for current != nil {
		parent = current
		if n.Key < current.Key {
			current = current.left
		} else {
			current = current.right
		}
	}
	if parent != nil {
		if parent.Key > n.Key {
			parent.left = n
		} else {
			parent.right = n
		}
	} else {
		b.root = n
	}
}

func (b *BST) InsertRec(n *BSTNode) {
	b.root = b.root.insertRec(n)
}

func (b *BSTNode) insertRec(n *BSTNode) *BSTNode {
	if b == nil {
		return n
	}
	if n.Key < b.Key {
		b.left = b.left.insertRec(n)
	} else if n.Key > b.Key {
		b.right = b.right.insertRec(n)
	}
	return b
}

func (b *BST) Delete(key int) {
	// find the node should be removed
	// super is the pointer of the pointer that
	// points to the node should be removed
	super := &b.root
	for key != (*super).Key && *super != nil {
		if key < (*super).Key {
			super = &(*super).left
		} else {
			super = &(*super).right
		}
	}

	if *super == nil {
		panic("key not found")
	}

	// at most one child
	if (*super).left == nil || (*super).right == nil {
		if (*super).left != nil {
			*super = (*super).left
		} else {
			*super = (*super).right
		}
		return
	}

	// if the deleted node have not only one child
	// find the successor
	// super_successor is the pointer of the pointer
	// that points to the successor
	super_successor := &(*super).right
	for (*super_successor).left != nil {
		super_successor = &(*super_successor).left
	}

	// replace the deleted node with the successor
	(*super).Key = (*super_successor).Key
	(*super).Val = (*super_successor).Val
	// connect the parent of successor to the child of successor.
	*super_successor = (*super_successor).right
}

// inorder traverse a tree
func (b *BST) List() []*BSTNode {
	// ret is a list records ordered nodes
	// that will be returned
	ret := []*BSTNode{}
	stack := []*BSTNode{}
	current := b.root
	for len(stack) != 0 || current != nil {
		if current != nil {
			// push current into stack
			stack = append(stack, current)
			current = current.left
		} else {
			// pop stack
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, current)
			current = current.right
		}
	}
	return ret
}

func (b *BST) Do(do func(*BSTNode)) {
	b.root.do(do)
}

package main

type NodeValue interface {
	// return positive number if n1 > n2
	// return 0 if n1 == n2
	// return negative number if n1 < n2
	Compare(NodeValue) int
}

type Node struct {
	Val   NodeValue
	left  *Node
	right *Node
}

type Bst struct {
	root *Node
}

func (b *Bst) Insert(n *Node) {
	// super is the pointer of the pointer
	// that points to the new node
	super := &b.root
	for *super != nil {
		if n.Val.Compare((*super).Val) < 0 {
			super = &(*super).left
		} else {
			super = &(*super).right
		}
	}
	*super = n
}

func (b *Bst) Delete(val NodeValue) {
	// find the node should be removed
	// super is the pointer of the pointer that
	// points to the node should be removed
	super := &b.root
	for val.Compare((*super).Val) != 0 && *super != nil {
		if val.Compare((*super).Val) < 0 {
			super = &(*super).left
		} else {
			super = &(*super).right
		}
	}

	if *super == nil {
		panic("can not found node")
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
	(*super).Val = (*super_successor).Val
	// connect the parent of successor to the child of successor.
	*super_successor = (*super_successor).right
}

// inorder traverse a tree
func (b *Bst) List() []*Node {
	// ret is a list records ordered nodes
	// that will be returned
	ret := []*Node{}
	stack := []*Node{}
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

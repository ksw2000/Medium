package avl

type AVLNode struct {
	Key    int
	Val    interface{}
	left   *AVLNode
	right  *AVLNode
	height int // the height of nodes
}

func NewAVLNode(key int, val interface{}) *AVLNode {
	return &AVLNode{
		Key:    key,
		Val:    val,
		left:   nil,
		right:  nil,
		height: 0,
	}
}

func (n *AVLNode) updateHeight() {
	left, right := -1, -1
	if n.left != nil {
		left = n.left.height
	}
	if n.right != nil {
		right = n.right.height
	}
	if left > right {
		n.height = left + 1
	} else {
		n.height = right + 1
	}
}

/*
	    n
	  /  \
	n1    n2
   /  \
n3     n4

       n1
	  /  \
    n3    n
		/  \
	   n4  n2

n3: no change children -> no change height
n4: no change children -> no change height
n2: no change children -> no change height

n1: right child is changed -> should change height
n: left child is changed -> should change height
*/

func (n *AVLNode) rightRotation() *AVLNode {
	n1 := n.left
	n.left = n1.right
	n1.right = n
	n.updateHeight()
	n1.updateHeight()
	return n1
}

/*
	 n
	/ \
 n2   n1
	 /  \
   n3    n4

	  n1
	 /  \
	n    n4
  /  \
n2   n3

n3: no change children -> no change height
n4: no change children -> no change height
n2: no change children -> no change height

n1: left child is changed -> should change height
n: right child is changed -> should change height
*/

func (n *AVLNode) leftRotation() *AVLNode {
	n1 := n.right
	n.right = n1.left
	n1.left = n
	n.updateHeight()
	n1.updateHeight()
	return n1
}

func (n *AVLNode) balanceFactor() int {
	// calculate balance factor
	left, right := -1, -1
	if n.left != nil {
		left = n.left.height
	}
	if n.right != nil {
		right = n.right.height
	}
	return left - right
}

// return the new root
func (n *AVLNode) rebalance() *AVLNode {
	bf := n.balanceFactor()
	if bf > 1 {
		// LR
		if n.left != nil && n.left.balanceFactor() < 0 {
			n.left = n.left.leftRotation()
		}
		// LL
		n = n.rightRotation()
	} else if bf < -1 {
		// RL
		if n.right != nil && n.right.balanceFactor() > 0 {
			n.right = n.right.rightRotation()
		}
		// RR
		n = n.leftRotation()
	}
	return n
}

func (n *AVLNode) insert(m *AVLNode) *AVLNode {
	if n == nil {
		return m
	}

	if m.Key < n.Key {
		n.left = n.left.insert(m)
	} else if m.Key > n.Key {
		n.right = n.right.insert(m)
	} else {
		n.Val = m.Val
	}
	n.updateHeight()
	return n.rebalance()
}

func (n *AVLNode) search(key int) *AVLNode {
	for current := n; current != nil; {
		if key < current.Key {
			current = current.left
		} else if key > current.Key {
			current = current.right
		} else {
			return current
		}
	}
	// if not found
	return nil
}

func (n *AVLNode) do(do func(*AVLNode)) {
	if n != nil {
		n.left.do(do)
		do(n)
		n.right.do(do)
	}
}

type AVLTree struct {
	root *AVLNode
}

func (t *AVLTree) Insert(n *AVLNode) {
	t.root = t.root.insert(n)
}

func (t *AVLTree) Search(key int) *AVLNode {
	return t.root.search(key)
}

func (t *AVLTree) Do(do func(*AVLNode)) {
	t.root.do(do)
}

package bst

import (
	"fmt"
	"testing"
)

type person struct {
	name string // as value
	tall int    // as key
}

// go test -v -run TestBST
func TestBST(t *testing.T) {
	nodes := []*person{
		{"Hanayo", 156}, {"Hanamaru", 152}, {"Maki", 161},
		{"Ruby", 154}, {"Riko", 160}, {"Eli", 162}, {"Mari", 163},
		{"Chisato", 155}, {"You", 157}, {"Ayumu", 159},
	}

	// Create a new binary search tree
	tree := new(BST)

	// Insert node to the tree
	for _, n := range nodes {
		tree.Insert(NewBSTNode(n.tall, n.name))
	}

	// print the tree inorder (way I, recursive)
	tree.Do(func(node *BSTNode) {
		fmt.Printf("(%d, %v)->", node.Key, node.Val)
	})
	fmt.Println("nil")

	// Print the tree inorder (way II, iterative)
	for _, node := range tree.List() {
		fmt.Printf("(%d, %v)->", node.Key, node.Val)
	}
	fmt.Println("nil")

	// Delete some nodes in tree
	tree.Delete(nodes[6].tall)
	tree.Delete(nodes[3].tall)
	tree.Delete(nodes[0].tall)

	for _, node := range tree.List() {
		fmt.Printf("(%d, %v)->", node.Key, node.Val)
	}
	fmt.Println("nil")

	// Search by key: 152
	if n, exist := tree.Search(152); exist {
		fmt.Println(152, n.Val)
	}
}

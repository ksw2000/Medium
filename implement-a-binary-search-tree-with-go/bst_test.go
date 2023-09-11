package bst

import (
	"fmt"
	"math/rand"
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

// go test -run TestBSTCorrectness
func TestBSTCorrectness(t *testing.T) {
	// generate an int list
	data := generateData(1024)
	// create an empty binary search tree
	tree := new(BST)
	// insert data
	for i := range data {
		tree.Insert(&BSTNode{Key: data[i]})
	}

	// check the binary search tree is in increasing order
	// and the total number of nodes is equal to the length of `data`
	i := 0
	start := 0
	tree.Do(func(b *BSTNode) {
		if b.Key >= start {
			start = b.Key
		} else {
			t.Fail()
		}
		i++
	})
	if i != len(data) {
		t.Fail()
	}

	// check that we can access all element by tree.Search()
	for i := range data {
		_, found := tree.Search(data[i])
		if !found {
			t.Fail()
		}
	}

	// check that we can remove each nodes
	k := len(data)
	for i := range data {
		tree.Delete(data[i])
		k--
		// check that the tree is in increasing order
		// and the total number of nodes is equal properly
		if i%32 == 0 {
			n := 0
			start := 0
			tree.Do(func(b *BSTNode) {
				if b.Key >= start {
					start = b.Key
				} else {
					t.Fail()
				}
				n++
			})
			if n != k {
				t.Fail()
			}
		}
	}
}

func generateData(size int) []int {
	rand.NewSource(0)
	list := make([]int, size)
	for i := range list {
		list[i] = i
	}
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	return list
}

// go test -bench=BenchmarkInsert -run=none
func BenchmarkInsertIterative(b *testing.B) {
	for n := 0; n < b.N; n++ {
		list := generateData(4096)
		tree := new(BST)
		for i := range list {
			tree.Insert(&BSTNode{Key: list[i]})
		}
	}
}

func BenchmarkInsertIterative2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		list := generateData(4096)
		tree := new(BST)
		for i := range list {
			tree.Insert2(&BSTNode{Key: list[i]})
		}
	}
}

func BenchmarkInsertRecursive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		list := generateData(4096)
		tree := new(BST)
		for i := range list {
			tree.InsertRec(&BSTNode{Key: list[i]})
		}
	}
}

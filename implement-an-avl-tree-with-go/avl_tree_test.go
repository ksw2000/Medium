package avl

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"testing"

	"github.com/karask/go-avltree"
)

func hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32() % 255)
}

// go test -v -run TestAVLTree
func TestAVLTree(t *testing.T) {
	tree := new(AVLTree)
	list := []string{
		"Chika", "You", "Ruby",
		"Riko", "Yoshiko", "Mari",
		"Kanan", "Dia", "Hanamaru",
	}

	for _, v := range list {
		fmt.Printf("Insert (%d %s)\n", hash(v), v)
		tree.Insert(NewAVLNode(int(hash(v)), v))
		tree.Do(func(a *AVLNode) {
			l := "nil"
			if left := a.left; left != nil {
				l = fmt.Sprintf("%3d", left.Key)
			}
			r := "nil"
			if right := a.right; right != nil {
				r = fmt.Sprintf("%3d", right.Key)
			}
			fmt.Printf("%3d left: %s right: %s balance: %2d height: %2d\n", a.Key, l, r, a.balanceFactor(), a.height)
		})
		fmt.Println("------------------")
	}

	tree.Do(func(a *AVLNode) {
		fmt.Printf("(%d, %s) -> ", a.Key, a.Val)
	})
	fmt.Println()
}

func TestAVLTreeCorrectness(t *testing.T) {
	list := make([]int, 1024)
	for i := range list {
		list[i] = i
	}
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	fmt.Println(list)
	tree := new(AVLTree)
	for _, e := range list {
		tree.Insert(NewAVLNode(e, e))
	}
	start := 0
	tree.Do(func(n *AVLNode) {
		// decreasing in order
		if n.Key < start {
			t.Fail()
		}
		start = n.Key

		// each node's balance factor is in [-1, 1]
		if n.balanceFactor() > 1 || n.balanceFactor() < -1 {
			t.Fail()
		}

		// search function is work
		if found := tree.Search(n.Key); found == nil || found.Key != n.Val.(int) {
			t.Fail()
		}
	})

	// try to delete
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	fmt.Println(list)

	for i := range list {
		tree.Delete(list[i])
		start := 0
		tree.Do(func(n *AVLNode) {
			// decreasing in order
			if n.Key < start {
				t.Fail()
			}
			start = n.Key
			// 		// each node's balance factor is in [-1, 1]
			if n.balanceFactor() > 1 || n.balanceFactor() < -1 {
				t.Fail()
			}
		})
	}
}

func generateData() ([]int, []int) {
	rand.New(rand.NewSource(0))
	list := make([]int, 1024)
	for i := range list {
		list[i] = i
	}
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	list2 := make([]int, len(list))
	copy(list2, list)
	rand.Shuffle(len(list2), func(i, j int) {
		list2[i], list2[j] = list2[j], list2[i]
	})
	return list, list2
}

func BenchmarkAVLKarask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list, list2 := generateData()
		tree := new(avltree.AVLTree)
		for _, key := range list {
			tree.Add(key, key)
		}
		for _, key := range list2 {
			tree.Remove(key)
		}
	}
}

func BenchmarkAVLOurs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list, list2 := generateData()
		tree := new(AVLTree)
		for _, key := range list {
			tree.Insert(NewAVLNode(key, key))
		}
		for _, key := range list2 {
			tree.Delete(key)
		}
	}
}

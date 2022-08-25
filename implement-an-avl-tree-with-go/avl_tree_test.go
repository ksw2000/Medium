package avl

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"testing"
)

func hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32() % 255)
}

// go test -run TestAVLTree -v
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
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, rand.Int())
	}
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
}

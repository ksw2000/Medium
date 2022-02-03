package main

import "fmt"

type person struct {
	name string
	tall int
}

func (p1 *person) Compare(p2 NodeValue) int {
	if p1.name == p2.(*person).name {
		return 0
	} else if p1.tall == p2.(*person).tall {
		return 1
	}
	return p1.tall - p2.(*person).tall
}

func main() {
	nodes := []*person{
		{"Honoka", 157}, {"Hanamaru", 152}, {"Maki", 161},
		{"Ruby", 154}, {"Riko", 160}, {"Eli", 162}, {"Mari", 163},
		{"Chisato", 155}, {"You", 157}, {"Ayumu", 159},
	}

	tree := new(Bst)
	for _, n := range nodes {
		tree.Insert(&Node{
			Val: n,
		})
	}

	for _, node := range tree.List() {
		fmt.Printf("%v->", node.Val)
	}
	fmt.Println("nil")

	tree.Delete(nodes[6])
	tree.Delete(nodes[3])
	tree.Delete(nodes[0])

	for _, node := range tree.List() {
		fmt.Printf("%v->", node.Val)
	}
	fmt.Println("nil")
}

package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var n1, n2 int
	countTreeNodes(t1, &n1)
	countTreeNodes(t2, &n2)
	if n1 != n2 {
		return false
	}
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i< n1; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true

}

func countTreeNodes(t *tree.Tree, num *int) {
	if t == nil {
		return
	}
	countTreeNodes(t.Left, num)
	*num += 1
	countTreeNodes(t.Right, num)
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println(Same(t1, t2))
}
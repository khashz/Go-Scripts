package main

import (
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
	fmt.Printf("Adding elem %d from tree\n", t.Value)
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for x := 0; x < 10; x++ {
		elem_t1 := <-ch1
		elem_t2 := <-ch2
		fmt.Println(elem_t1, elem_t2)
		if elem_t1 != elem_t2 {
			return false
		}
	}
	close(ch1)
	close(ch2)
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walk func(t *tree.Tree, ch chan int)

	walk = func(t *tree.Tree, ch chan int) {
		if t == nil {
			return
		}
		walk(t.Left, ch)
		ch <- t.Value
		walk(t.Right, ch)
	}

	walk(t, ch)

	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	// for v1 := range c1 {
	// 	v2 := <-c2
	// 	if v1 != v2 {
	// 		return false
	// 	}
	// }

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(2), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))

	fmt.Println(tree.New(1).String())
}

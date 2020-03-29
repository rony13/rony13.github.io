package main

import (
	"fmt"
)

func main() {
	a := []int{-5, -5}

	fmt.Print(reversePairs(a))
}

type Node struct {
	Value int
	LSum  int
	RSum  int
	L     *Node
	R     *Node
}

func (n *Node) Insert(son *Node) {
	if son.Value <= n.Value {
		if n.L == nil {
			n.L = son
		} else {
			n.L.Insert(son)
		}
		n.LSum++
	} else {
		if n.R == nil {
			n.R = son
		} else {
			n.R.Insert(son)
		}
		n.RSum++
	}
}

func (n *Node) Search(v int) int {
	if n == nil {
		return 0
	}

	if n.Value > v {
		return n.L.Search(v) + 1 + n.RSum
	}

	if n.Value <= v {
		return n.R.Search(v)
	}
	return 0
}

func reversePairs(nums []int) int {
	var root *Node
	count := 0
	for i, n := range nums {
		if i == 0 {
			root = &Node{
				Value: n,
			}
			continue
		}
		count += root.Search(n * 2)
		root.Insert(&Node{
			Value: n,
		})
	}
	return count
}

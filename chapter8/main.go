package main

import (
	"chapter8/list"
	"errors"
	"fmt"
)

func main() {
	l1, l2 := getList()

	node, err := findIntersectNode(l1, l2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(node)
}

func findIntersectNode(l1 *list.List, l2 *list.List) (*list.Node, error) {
	len1 := l1.Len()
	len2 := l2.Len()
	p1 := l1.Head
	p2 := l2.Head
	minLen := 0
	// 链表长先走|len1 - len2|步
	if len1 > len2 {
		p1, _ = l1.GetNode(len1 - len2)
		minLen = len2
	} else if len1 < len2 {
		p2, _ = l2.GetNode(len2 - len1)
		minLen = len1
	}

	// 找第一个相交点
	for i := 0; i < minLen; i++ {
		if p1.Next == p2.Next {
			return p1.Next, nil
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return nil, errors.New("no intersect")
}

func getList() (*list.List, *list.List) {
	nodeX := &list.Node{Data: "x", Next: nil}
	nodeY := &list.Node{Data: "y", Next: nil}
	nodeZ := &list.Node{Data: "z", Next: nil}
	l3 := list.NewList()
	l3.Add(nodeX)
	l3.Add(nodeY)
	l3.Add(nodeZ)

	l1 := list.NewList()
	nodeA := &list.Node{Data: "a", Next: nil}
	l1.Add(nodeA)

	nodeB := &list.Node{Data: "b", Next: nil}
	l1.Add(nodeB)
	l1.Add(nodeX)

	l2 := list.NewList()
	nodeD := &list.Node{Data: "d", Next: nil}
	l2.Add(nodeD)

	nodeE := &list.Node{Data: "e", Next: nil}
	l2.Add(nodeE)

	nodeF := &list.Node{Data: "f", Next: nil}
	l2.Add(nodeF)
	l2.Add(nodeX)
	return l1, l2
}

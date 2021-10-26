package list

import (
	"errors"
	"fmt"
)

type Node struct {
	Data string
	Next *Node
}

type List struct {
	Head *Node
}

func NewList() *List  {
	return &List{}
}

func (l *List) Add(node *Node) {
	// 第一个节点
	if l.Head == nil {
		l.Head = node
		return
	}

	// 找到尾节点
	tail := l.Head
	for ; tail.Next != nil; tail = tail.Next {
	}
	tail.Next = node
}

func (l *List) Len() int {
	count := 0
	for tail := l.Head; tail.Next != nil; tail = tail.Next {
		count += 1
	}

	return count
}

func (l *List) GetNode(n int) (*Node, error) {
	if n <= 0 || l.Len() < n {
		return nil, errors.New("n is invalid")
	}

	node := l.Head
	for i := 0; i < n; i++ {
		node = node.Next
	}

	return node, nil
}

func (l *List) Print() {
	tail := l.Head
	for ; tail.Next != nil; tail = tail.Next {
		fmt.Println(tail)
	}
	fmt.Println(tail)
}
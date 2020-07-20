package list


type Node struct {
	data string
	next *Node
}

type List struct {
	head *Node
}

func NewList() List  {
	return List{}
}

func (l *List) add(data string) {
	node := &Node{data: data, next: nil}
	// 第一个节点
	if l.head == nil {
		l.head = node
		return
	}

	// 找到尾节点
	for tail := l.head; tail.next != nil; tail = tail.next {
		tail.next = node
	}
}

func (l *List) len() int {
	count := 0
	for tail := l.head; tail.next != nil; tail = tail.next {
		count += 1
	}

	return count
}

fun
package composite

import "fmt"

type ComponenDesc struct {
	id   int64
	name string
}

func (c *ComponenDesc) Id() int64 {
	return c.id
}

func (c *ComponenDesc) Name() string {
	return c.name
}

type IUIComponent interface {
	Id() int64
	Print()
}

type UINode struct {
	ComponenDesc
}

func NewUINode(id int64, name string) *UINode {
	return &UINode{
		ComponenDesc: ComponenDesc{
			id:   id,
			name: name,
		},
	}
}

func (u *UINode) Id() int64 {
	return u.id
}

func (u *UINode) Print() {
	fmt.Printf("UINode:%d\t%s\n", u.id, u.name)
}

type UIGroup struct {
	children []IUIComponent
}

func (u *UIGroup) Add(item IUIComponent) {
	u.children = append(u.children, item)
}

func (u *UIGroup) Remove(item IUIComponent) {
	tmp := make([]IUIComponent, 0)
	for _, child := range u.children {
		if child.Id() != item.Id() {
			tmp = append(tmp, child)
		}
	}
	u.children = tmp
}

func (u *UIGroup) GetChildren() []IUIComponent {
	return u.children
}

type UIFrame struct {
	ComponenDesc
	UIGroup
}

func NewUIFrame(id int64, name string) *UIFrame {
	return &UIFrame{
		ComponenDesc: ComponenDesc{
			id:   id,
			name: name,
		},
	}
}

func (u *UIFrame) Id() int64 {
	return u.id
}

func (u *UIFrame) Print() {
	fmt.Printf("UIFrame:%d\t%s\n", u.id, u.name)
	for _, child := range u.children {
		child.Print()
	}
	fmt.Println()
}

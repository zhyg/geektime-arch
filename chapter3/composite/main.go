package main

import (
	"arch/composite"
	"fmt"
)

func main() {
	win := composite.NewUIFrame(1, "WINDOW 窗口")
	p1 := composite.NewUINode(11, "LOGO图片")
	p2 := composite.NewUINode(12, "登录")
	p3 := composite.NewUINode(13, "注册")

	frame1 := composite.NewUIFrame(2, "FRAME1")
	p4 := composite.NewUINode(21, "用户名")
	p5 := composite.NewUINode(22, "文本框")
	p6 := composite.NewUINode(23, "密码")
	p7 := composite.NewUINode(24, "密码框")
	p8 := composite.NewUINode(25, "复选框")
	p9 := composite.NewUINode(26, "记住用户名")
	p10 := composite.NewUINode(27, "忘记密码")
	frame1.Add(p4)
	frame1.Add(p5)
	frame1.Add(p6)
	frame1.Add(p7)
	frame1.Add(p8)
	frame1.Add(p9)
	frame1.Add(p10)

	win.Add(p1)
	win.Add(p2)
	win.Add(p3)
	win.Add(frame1)

	win.Print()
}

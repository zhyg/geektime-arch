#电话拨号

## 分析

1. 按键（数字、Send）以事件发出，解耦界面操作和核心处理程序
2. 核心处理程序根据按键事件类型，来控制屏幕、扬声器、无线网络
3. 抽象各硬件接口（屏幕、扬声器），方便对接不同类型的硬件

## 组件图

![uml2](uml2.png)


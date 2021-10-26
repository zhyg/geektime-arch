#Q1: 请描述什么是依赖倒置原则，为什么有时候依赖倒置原则又被称为好菜坞原则？

## 依赖倒置原则

1. 高层模块（调用者）和低层模块（被调用者）应该通过抽象来依赖
2. 具体依赖抽象

DIP 主要用来指导框架层面的设计

## 好莱坞原则

1. don’t call us, we’ll call you.

底层模块将自己挂钩到系统上，高层模块决定什么时候和怎么使用底层模块

## 共同点

1. 目标都是解耦

# Q2:请描述一个你熟悉的框架，是如何实现依赖倒置原则的。

## gorm 插件系统

1. 通过[gorm](https://github.com/go-gorm/gorm)插件系统，开放框架的扩展点，支持六种SQL类型的回调。比如通过注册 callback 统计SQL执行时间、实现读写分离等

```go
func initializeCallbacks(db *DB) *callbacks {
	return &callbacks{
		processors: map[string]*processor{
			"create": {db: db},
			"query":  {db: db},
			"update": {db: db},
			"delete": {db: db},
			"row":    {db: db},
			"raw":    {db: db},
		},
	}
}
```

2. 读写分离插件

```go
db.Query().Before("gorm:query").Register("rwsplit:setdb", rwsplit.ReadDB)
db.Query().Before("gorm:before_create").Register("rwsplit:setdb", rwsplit.WriteDB)
db.Query().Before("gorm:before_update").Register("rwsplit:setdb", rwsplit.WriteDB)
db.Query().Before("gorm:before_delete").Register("rwsplit:setdb", rwsplit.WriteDB)
```

3. gorm 在执行过程中按顺序调用自定义的 callbacks
4. 用户实现自定义插件（统计SQL执行时间、实现读写分离等）只依赖六种回调定义，不关心回调的具体执行
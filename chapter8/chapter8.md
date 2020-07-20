# 1. 有两个单向链表（链表长度分别为m，n），这两个单向链表有可能在某个元素合并，如下图所示的这样，也可能不合并。现在给定两个链表的头指针，在不修改链表的情况下，如何快速地判断这两个链表是否合并？如果合并，找到合并的元素，也就是图中的 x 元素。请用（伪）代码描述算法，并给出时间复杂度和空间复杂度。
## 算法思想
1. 长的链表先走|len1 - len2|步
2. 同时遍历两个链表，找到第一个相同的节点

## 代码

见 main.go、list.go

## 复杂度

时间复杂度：O(n)
空间复杂度：O(n)

# 2.请画出DataNode服务器节点宕机的时候，HDFS的处理过程时序图。



```mermaid
sequenceDiagram
    participant Client
    participant DataNode1
    participant DataNode2（宕机）
    participant DataNode3
    participant NameNode
    
    DataNode2（宕机） ->> NameNode : 上报心跳
    
    NameNode ->> DataNode2（宕机） : 配置时间内没有收到心跳，主动检测
    alt 检测超时
    NameNode ->> NameNode : 标记DataNode2宕机
    end
    
    NameNode ->> Client : 通知读写DataNode1
    Client ->> DataNode1 : 读写
    NameNode ->> DataNode3 : 从DataNode1上同步DataNode2存储的副本
    DataNode3 ->> + DataNode1 : 同步副本
    DataNode1 -->> - DataNode3 : 副本数据
    DataNode3 -->> NameNode : 同步完成
   
```
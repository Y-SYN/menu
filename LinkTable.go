package main

import "fmt"

/* 定义数据结构 */
type Node struct {
	data interface{}
	next *Node
}
type LinkTable struct {
	headNode *Node //头节点，不存放数据
	length   int   //当前链表的长度
}

/* 初始化链表 */
func initLinkTable() *LinkTable {
	node := new(Node)
	linkTable := new(LinkTable)
	linkTable.headNode = node
	linkTable.length = 0
	return linkTable
}

/* 获取尾节点 */
func getTailNode(linkTable *LinkTable) *Node {
	curNode := linkTable.headNode
	//当前链表只有一个头节点
	if curNode.next == nil {
		return linkTable.headNode
	}
	for curNode.next != nil {
		curNode = curNode.next
	}
	return curNode
}

/* 插入元素（头插法） */
func insertHeadNode(data interface{}, linkTable *LinkTable) {
	node := &Node{data, nil}
	node.next = linkTable.headNode.next
	linkTable.headNode.next = node
	linkTable.length++
}

/* 尾插法 */
func insertTailNode(data interface{}, linkTable *LinkTable) {
	node := &Node{data, nil}
	curNode := getTailNode(linkTable)
	curNode.next = node
	linkTable.length++
}

/* 指定位置插入节点 */
func append(index int, data interface{}, linktable *LinkTable) {
	if index <= 0 || index > linktable.length+1 {
		fmt.Println("当前位置错误")
		return
	}
	curNode := linktable.headNode
	//找到需要插入节点的前一个节点
	for j := 0; j < index-1; j++ {
		curNode = curNode.next
	}
	node := new(Node)
	node.data = data
	node.next = curNode.next
	curNode.next = node
	linktable.length++
}

/* 展示当前链表的数据，从头节点开始 */
func showLinkTableData(linkTable *LinkTable) {
	curNode := linkTable.headNode.next
	for curNode != nil {
		fmt.Print(curNode.data)
		fmt.Print(" ")
		curNode = curNode.next
	}
	fmt.Println()
}

/* 按序号查找某节点 */
func findNodeByIndex(index int, linktable *LinkTable) *Node {
	if index <= 0 || index > linktable.length {
		fmt.Println("当前序号错误")
		return nil
	}
	curNode := linktable.headNode
	for j := 0; j < index; j++ {
		curNode = curNode.next
	}
	return curNode
}

/* 按值查找节点 */
func findNodeByValue(data interface{}, linktable *LinkTable) *Node {
	curNode := linktable.headNode.next
	for curNode != nil {
		if curNode.data == data {
			return curNode
		}
		curNode = curNode.next
	}
	return nil
}

/* 删除第i个节点 */
func deleteNode(index int, linktable *LinkTable) {
	if index <= 0 || index > linktable.length {
		fmt.Println("当前位置错误")
		return
	}
	curNode := linktable.headNode
	//找到需要删除节点的前一个节点
	for j := 0; j < index-1; j++ {
		curNode = curNode.next
	}
	curNode.next = curNode.next.next
	linktable.length--
}

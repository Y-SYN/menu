package main

import "fmt"

func main() {
	testLinkTable()
}

func testLinkTable() {
	linktable := initLinkTable()
	insertTailNode(1, linktable)
	insertTailNode(2, linktable)
	insertTailNode(3, linktable)
	insertTailNode(4, linktable)
	insertTailNode(5, linktable)
	insertTailNode(6, linktable)
	fmt.Println("初始链表,含6个节点:")
	showLinkTableData(linktable)
	fmt.Println("在index=3处,插入值为7的节点:")
	append(2, 7, linktable)
	showLinkTableData(linktable)
	fmt.Println("查找index=1的节点:")
	fmt.Println(findNodeByIndex(1, linktable))
	fmt.Println("查找值为7的节点:")
	fmt.Println(findNodeByValue(7, linktable))
	fmt.Println("删除第6个节点:")
	deleteNode(6, linktable)
	showLinkTableData(linktable)
}

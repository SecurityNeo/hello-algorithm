package main

import (
	"fmt"
	"math/rand"
	"time"
)

type linkNode struct {
	val  int
	next *linkNode
}

type linkNoder interface {
	Insert(data int)
	Print()
	getHead() *linkNode
}

func NewLinkNode() linkNoder {
	return &linkNode{}
}

// Insert 在链表尾部插入给定数据
func (head *linkNode) Insert(data int) {
	if head == nil {
		fmt.Printf("null link node")
		return
	}
	cur := head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &linkNode{val: data}
}

// Print 打印链表所有节点数据
func (head *linkNode) Print() {
	cur := head
	if cur == nil {
		fmt.Printf("null link node")
		return
	}
	for i := 0; cur != nil; i++ {
		fmt.Printf("[%d]==>", cur.val)
		cur = cur.next
	}
	fmt.Println("[end]")
}

// getHead 返回链表头
func (head *linkNode) getHead() *linkNode {
	return head
}

// genLinkNode 生成指定长度链表，返回链表头
func genLinkNode(nodeCount int) *linkNode {
	head := NewLinkNode()
	for i := 0; i < nodeCount; i++ {
		rand.Seed(time.Now().UnixNano())
		val := rand.Intn(100)
		head.Insert(val)
		time.Sleep(1 * time.Nanosecond)
	}
	// linkHead 带头节点
	linkHead := head.getHead()
	// link 不带头节点
	link := linkHead.next
	return link
}

// reverseLink 反转链表。时间复杂度 O(N) ；空间复杂度 O(1)
func reverseLink(link *linkNode) *linkNode {
	var pre *linkNode
	for link != nil {
		temp := link.next
		link.next = pre
		pre = link
		link = temp
	}
	return pre
}

// reverseLink2 单链表反转（不带头），原地反转
func reverseLink2(link *linkNode) *linkNode {
	// 先给链表加个头
	headNode := NewLinkNode()
	head := headNode.getHead()
	head.next = link

	beg := head.next
	end := head.next.next
	for end != nil {
		// 把end所在节点跳过，beg的next直接指向end的下一个节点
		beg.next = end.next
		// end的next指向第一个节点
		end.next = head.next
		// 头节点指向新放上来的end节点，此时end就到了第一个位置
		head.next = end
		// 把end移到beg的下一个节点，进行下一轮循环
		end = beg.next
	}
	// 把头去掉返回
	return head.next
}

// reverseLink3 单链表（不带头）反转。头插法：把链表中的每个节点依次插入到head的next里
func reverseLink3(link *linkNode) *linkNode {
	// 先给链表加个头
	headNode := NewLinkNode()
	head := headNode.getHead()
	head.next = link

	p1 := head.next
	// 先把头断掉
	head.next = nil
	for p1 != nil {
		p2 := p1.next
		p1.next = head.next
		head.next = p1
		p1 = p2
	}
	// 把头去掉返回
	return head.next
}

// link 单链表相关函数入口
func link() {
	// 1.链表反转
	link := genLinkNode(8)
	fmt.Println("before reverse:")
	link.Print()
	res := reverseLink(link)
	fmt.Println("reverse result:")
	res.Print()
	res2 := reverseLink2(res)
	res2.Print()
	res3 := reverseLink3(res2)
	res3.Print()

}

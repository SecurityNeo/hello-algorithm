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

func (head *linkNode) getEnd() *linkNode {
	cur := head
	for cur.next != nil {
		cur = cur.next
	}
	return cur
}

// getMiddle 返回链表中间节点指针
func (head *linkNode) getMiddle() *linkNode {
	slow := head.next
	fast := head.next

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// getLen 获取链表节点数量
func (head *linkNode) getLen() (length int) {
	cur := head
	for cur != nil {
		cur = cur.next
		length++
	}
	return
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

// isPalindrome 判断链表是否为回文链表
func isPalindrome(link *linkNode) {
	if link == nil || link.next == nil {
		fmt.Println("not enough node")
		return
	}

	mid := link.getMiddle()
	rMid := reverseLink3(mid)
	for rMid != nil {
		if rMid.val != link.val {
			fmt.Println("it is not a palindrome link")

			return
		}
		rMid = rMid.next
		link = link.next
	}
	fmt.Println("it is a palindrome link")
}

// linkPartition 链表荷兰国旗问题。pivot是基准值。
func linkPartition(link *linkNode, pivot int) {
	if link == nil || link.next == nil {
		fmt.Println("not enough node")
		return
	}
	// sH 小于区的第一个节点		sT 小于区的最后一个节点
	// eH 等于区的第一个节点		eT 等于区的最后一个节点
	// bH 大于区的第一个节点		bT 大于区的最后一个节点
	var sH *linkNode
	var sT *linkNode
	var eH *linkNode
	var eT *linkNode
	var bH *linkNode
	var bT *linkNode
	var nLink *linkNode

	for link != nil {
		// 先将第二个节点地址保存下来
		backupSecondNode := link.next
		link.next = nil
		if link.val < pivot {
			if sT == nil {
				sH = link
				sT = link
			} else {
				sT.next = link
				sT = sT.next
			}
		} else if link.val == pivot {
			if eT == nil {
				eH = link
				eT = link
			} else {
				eT.next = link
				eT = eT.next
			}
		} else {
			if bT == nil {
				bH = link
				bT = link
			} else {
				bT.next = link
				bT = bT.next
			}
		}
		link = backupSecondNode
	}

	if sT != nil && eT != nil {
		sT.next = eH
		eT.next = bH
		nLink = sH
	}
	if sT == nil && eT != nil {
		eT.next = bH
		nLink = eH
	}
	if sT != nil && eT == nil {
		sT.next = bH
		nLink = sH
	}
	nLink.Print()
}

// isCircleLink 判断单链表是否有环，返回入环节点
func isCircleLink(link *linkNode) *linkNode {
	if link == nil || link.next == nil {
		fmt.Println("not enough node")
		return nil
	}
	fast := link.next.next
	slow := link.next
	for slow != fast {
		if fast.next == nil || fast.next.next == nil {
			fmt.Println("link has no circle")
			return nil
		}
		slow = slow.next
		fast = fast.next.next
	}
	fmt.Print("link has circle,circle start at: ")
	fast = link
	for slow != fast {
		fast = fast.next
		slow = slow.next
	}
	fmt.Println(fast.val)
	return fast
}

// isIntersect 判断两个单链表是否相交
func isIntersect(linkA, linkB *linkNode) {
	// 1 两个无环单链表相交
	// 2 两个有环单链表相交
	// 2.1 两个链表在入环之前相交
	// 2.2 两个链表在环上相交
	loopA := isCircleLink(linkA)
	loopB := isCircleLink(linkB)
	// 1 两个无环单链表相交问题
	if loopA == nil && loopB == nil {
		if intersect := noLoopLinkIsIntersect(linkA, linkB); intersect != nil {
			fmt.Print("intersect at: ")
			fmt.Println(intersect.val)
			return
		}
		fmt.Println("disjoint")
		return
		// 2 两个有环单链表相交问题
	} else if loopA != nil && loopB != nil {
		if intersect := bothLoopLinkIsIntersect(linkA, loopA, linkB, loopB); intersect != nil {
			fmt.Print("intersect at: ")
			fmt.Println(intersect.val)
			return
		}
		fmt.Println("disjoint")
		return
	}
	fmt.Println("disjoint")
}

// noLoopLinkIsIntersect 判断两个无环单链表是否相交，返回相交节点
func noLoopLinkIsIntersect(linkA, linkB *linkNode) *linkNode {
	if linkA == nil || linkB == nil {
		return nil
	}
	var linkALen int
	var linkBLen int
	curA := linkA
	curB := linkB
	for curA.next != nil {
		curA = curA.next
		linkALen++
	}

	for curB.next != nil {
		curB = curB.next
		linkBLen++
	}

	if curA != curB {
		return nil
	}

	if linkALen > linkBLen {
		gap := linkALen - linkBLen
		pA := linkA
		pB := linkB
		for i := 0; i < gap; i++ {
			pA = pA.next
		}
		for pA != pB {
			pA = pA.next
			pB = pB.next
		}
		return pA
	} else if linkALen < linkBLen {
		gap := linkBLen - linkALen
		pA := linkA
		pB := linkB
		for i := 0; i < gap; i++ {
			pB = pB.next
		}
		for pA != pB {
			pA = pA.next
			pB = pB.next
		}
		return pA
	} else {
		pA := linkA
		pB := linkB
		for pA != pB {
			pA = pA.next
			pB = pB.next
		}
		return pA
	}
}

// bothLoopLinkIsIntersect 判断两个有环单链表linkA与linkB是否相交（loopA和loopB分别为两个链表的入环节点），返回相交节点
func bothLoopLinkIsIntersect(linkA, loopA, linkB, loopB *linkNode) *linkNode {
	// 1 两个链表不相交
	// 2 两个链表在入环之前相交
	// 3 两个链表在环上相交

	// 两个链表在入环之前相交的场景
	// 这种场景下可以简化为无环单链表的相交问题，将入环节点当做链表结尾
	if loopA == loopB {
		tempNode := loopA.next
		loopA.next = nil
		intersect := noLoopLinkIsIntersect(linkA, linkB)
		loopA.next = tempNode
		return intersect
	}

	// 两个链表不相交
	// 两个链表在环上相交
	cur := loopA.next
	for cur != loopA {
		if cur == loopB {
			return loopA
		}
		cur = cur.next
	}
	return nil
}

// link 单链表相关函数入口
func link() {
	// 1.链表反转01
	link := genLinkNode(8)
	fmt.Print("before reverse:")
	link.Print()
	res := reverseLink(link)
	fmt.Print("reverse result:")
	res.Print()
	// 1.链表反转02
	res2 := reverseLink2(res)
	res2.Print()
	// 1.链表反转03
	res3 := reverseLink3(res2)
	res3.Print()

	// 2.回文链表
	head := NewLinkNode()
	linkHead := head.getHead()
	linkHead.Insert(1)
	linkHead.Insert(2)
	linkHead.Insert(3)
	linkHead.Insert(2)
	linkHead.Insert(1)
	fmt.Print("target link:")
	linkHead.next.Print()
	isPalindrome(linkHead.next)
	fmt.Print("target link:")
	res3.Print()
	isPalindrome(res3)

	// 3.链表的荷兰国旗问题
	link2 := genLinkNode(8)
	fmt.Print("original node:       ")
	link2.Print()
	fmt.Print("linkPartition result:")
	linkPartition(link2, 50)

	// 4.判断单链表是否有环
	head2 := NewLinkNode()
	linkHead2 := head2.getHead()
	linkHead2.Insert(1)
	linkHead2.Insert(2)
	linkHead2.Insert(3)
	linkHead2.Insert(4)
	linkHead2.Insert(5)
	end := linkHead2.getEnd()
	end.next = linkHead2.next.next.next
	isCircleLink(linkHead2.next)

	// 5.判断两个单链表是否相交
	// 5.1 两个无环单链表
	link3 := genLinkNode(6)
	link4 := genLinkNode(7)
	isIntersect(link3, link4)

	// 5.2 两个有环单链表，环上相交
	head5 := NewLinkNode()
	linkHead5 := head5.getHead()
	linkHead5.Insert(10)
	linkHead5.Insert(20)
	linkHead5.Insert(30)
	linkHead5.Insert(40)
	linkHead5.Insert(50)
	linkHead5.Insert(60)
	link5 := linkHead5.next
	link5.getEnd().next = link5.next.next

	head6 := NewLinkNode()
	linkHead6 := head6.getHead()
	linkHead6.Insert(100)
	linkHead6.Insert(200)
	link6 := linkHead6.next
	link6.getEnd().next = link5.next.next.next

	isIntersect(link5, link6)
}

package main

import (
	"container/list"
	"errors"
	"fmt"
	"math"
)

type binaryTreeNode struct {
	val   int
	left  *binaryTreeNode
	right *binaryTreeNode
}

type stack struct {
	cap      int               // 栈存放的最大元素个数
	topIndex int               // 栈顶元素索引
	arr      []*binaryTreeNode //	数组模拟栈
}

// push 压栈操作
func (s *stack) push(node *binaryTreeNode) error {
	if s.topIndex == s.cap-1 {
		return errors.New("stack is full")
	}
	s.topIndex++
	s.arr[s.topIndex] = node
	return nil
}

// pop 从栈顶弹出一个节点
func (s *stack) pop() (node *binaryTreeNode, err error) {
	if s.topIndex == -1 {
		return nil, errors.New("stack is null")
	}
	node = s.arr[s.topIndex]
	s.arr[s.topIndex] = nil
	s.topIndex--
	return node, err
}

func (s *stack) list() {
	if s.topIndex == -1 {
		return
	}
	for i := 0; i < s.cap; i++ {
		fmt.Printf("[%d]->", s.arr[i].val)
	}
}

func (s *stack) posList() {
	if s.topIndex == -1 {
		return
	}
	for i := s.topIndex; i > -1; i-- {
		fmt.Printf("[%d]->", s.arr[i].val)
	}
}

func newStack(cap int) (s *stack, err error) {
	if cap > 1024 {
		return nil, errors.New("stack cap: 1024")
	}
	arr := make([]*binaryTreeNode, cap)
	return &stack{cap: cap, topIndex: -1, arr: arr}, nil
}

func (node *binaryTreeNode) getNodeCount() (nodeCount int) {
	if node == nil {
		return 0
	}
	nodeCount = node.left.getNodeCount() + node.right.getNodeCount() + 1
	return
}

// recursionAllPrint 二叉树递归序打印
func (node *binaryTreeNode) recursionAllPrint() {
	if node == nil {
		return
	}
	fmt.Printf("[%d]->", node.val)
	node.left.recursionAllPrint()
	fmt.Printf("[%d]->", node.val)
	node.right.recursionAllPrint()
	fmt.Printf("[%d]->", node.val)
}

// recursionPrePrint 二叉树先序遍历（递归）（头左右）
func (node *binaryTreeNode) recursionPrePrint() {
	if node == nil {
		return
	}
	fmt.Printf("[%d]->", node.val)
	node.left.recursionPrePrint()
	node.right.recursionPrePrint()
}

// recursionMidPrint 二叉树中序遍历（递归）（左头右）
func (node *binaryTreeNode) recursionMidPrint() {
	if node == nil {
		return
	}
	node.left.recursionMidPrint()
	fmt.Printf("[%d]->", node.val)
	node.right.recursionMidPrint()
}

// recursionPosPrint 二叉树后序遍历（递归）（左右头）
func (node *binaryTreeNode) recursionPosPrint() {
	if node == nil {
		return
	}
	node.left.recursionPosPrint()
	node.right.recursionPosPrint()
	fmt.Printf("[%d]->", node.val)
}

// unRecursionPrePrint 二叉树先序遍历（非递归）
func (node *binaryTreeNode) unRecursionPrePrint() {
	if node == nil {
		return
	}
	nodeStack, _ := newStack(7)
	// 先把头压入栈中
	if err := nodeStack.push(node); err != nil {
		fmt.Println(err)
		return
	}
	// 栈中有数据就弹出，然后先右后左该节点的子节点压入栈中
	for nodeStack.topIndex > -1 {
		tmpNode, _ := nodeStack.pop()
		if tmpNode != nil {
			fmt.Printf("[%d]->", tmpNode.val)
		}
		if tmpNode.right != nil {
			nodeStack.push(tmpNode.right)
		}
		if tmpNode.left != nil {
			nodeStack.push(tmpNode.left)
		}
	}
	fmt.Println()
}

// unRecursionMidPrint 二叉树中序遍历（非递归）
func (node *binaryTreeNode) unRecursionMidPrint() {
	if node == nil {
		return
	}
	nodeCount := node.getNodeCount()
	nodeStack, _ := newStack(nodeCount)
	head := node
	for nodeStack.topIndex > -1 || head != nil {
		if head != nil {
			nodeStack.push(head)
			head = head.left
		} else {
			tmpNode, _ := nodeStack.pop()
			fmt.Printf("[%d]->", tmpNode.val)
			head = tmpNode.right
		}
	}
	fmt.Println()
}

// unRecursionPosPrint 二叉树后续遍历（非递归）
// 先序遍历是头左右，更改入栈顺序后就是头右左，再反转打印就是左右头，即为后续遍历的正确顺序了
func (node *binaryTreeNode) unRecursionPosPrint() {
	if node == nil {
		return
	}
	nodeCount := node.getNodeCount()
	nodeStack, _ := newStack(nodeCount)
	showSTack, _ := newStack(nodeCount)

	nodeStack.push(node)
	for nodeStack.topIndex > -1 {
		tmpNode, _ := nodeStack.pop()
		showSTack.push(tmpNode)
		if tmpNode.left != nil {
			nodeStack.push(tmpNode.left)
		}
		if tmpNode.right != nil {
			nodeStack.push(tmpNode.right)
		}
	}
	//showSTack.posList()
	for showSTack.topIndex > -1 {
		nodeVal, _ := showSTack.pop()
		fmt.Printf("[%d]->", nodeVal.val)
	}
	fmt.Println()
}

// BFS 二叉树广度优先遍历（BFS）
func (node *binaryTreeNode) BFS() {
	if node == nil {
		return
	}
	// 队列先进先出
	queue := list.New()
	queue.PushBack(node)
	for queue.Len() > 0 {
		tmpNode := queue.Remove(queue.Front()).(*binaryTreeNode)
		fmt.Printf("[%d]->", tmpNode.val)
		if tmpNode.left != nil {
			queue.PushBack(tmpNode.left)
		}
		if tmpNode.right != nil {
			queue.PushBack(tmpNode.right)
		}
	}
	fmt.Println()
}

// isBST 判断一个二叉树是否为搜索二叉树。搜索二叉树中序是升序排列的。
// 1、左子树比根节点小，右子树比根节点大
// 2、中序遍历时，其节点的值一定是升序排列的
func (node *binaryTreeNode) isBST() {
	if node == nil {
		fmt.Println("BST: true")
	}
	max := math.MinInt
	head := node
	nodeCount := node.getNodeCount()
	nodeStack, _ := newStack(nodeCount)
	for nodeStack.topIndex > -1 || head != nil {
		if head != nil {
			nodeStack.push(head)
			head = head.left
		} else {
			tmpNode, _ := nodeStack.pop()
			if tmpNode.val > max {
				max = tmpNode.val
			} else {
				fmt.Println("BST: false")
				return
			}
			head = tmpNode.right
		}
	}
	fmt.Println("BST: true")
}

// isCBT 判断一个二叉树是否完全二叉树
// 1、存在一个节点只有右子树，此树不是完全二叉树
// 2、宽度优先遍历时，存在一个节点左右子树不双全后，后续节点必须是叶子节点，否则此树不是完全二叉树
func (node *binaryTreeNode) isCBT() {
	// 遍历过程中，遇到第一个节点左右两个孩子不双全时置为true
	flag := false
	if node == nil {
		fmt.Println("CBT: true")
		return
	}

	queue := list.New()
	queue.PushBack(node)
	for queue.Len() > 0 {
		tmpNode := queue.Remove(queue.Front()).(*binaryTreeNode)
		if (flag && (tmpNode.left != nil || tmpNode.right != nil)) || (tmpNode.left == nil && tmpNode.right != nil) {
			fmt.Println("CBT: false")
			return
		}

		if tmpNode.left != nil {
			queue.PushBack(tmpNode.left)
		}
		if tmpNode.right != nil {
			queue.PushBack(tmpNode.right)
		}
		if tmpNode.left == nil || tmpNode.right == nil {
			flag = true
		}
	}
	fmt.Println("CBT: true")
}

func binaryTree() {
	binaryTree7 := &binaryTreeNode{7, nil, nil}
	binaryTree6 := &binaryTreeNode{6, nil, nil}
	binaryTree5 := &binaryTreeNode{5, nil, nil}
	binaryTree4 := &binaryTreeNode{4, nil, nil}
	binaryTree3 := &binaryTreeNode{3, binaryTree6, binaryTree7}
	binaryTree2 := &binaryTreeNode{2, binaryTree4, binaryTree5}
	binaryTree := &binaryTreeNode{1, binaryTree2, binaryTree3}
	fmt.Print("recursionAllPrint: ")
	binaryTree.recursionAllPrint()
	fmt.Println()
	fmt.Print("recursionPrePrint: ")
	binaryTree.recursionPrePrint()
	fmt.Println()
	fmt.Print("recursionMidPrint: ")
	binaryTree.recursionMidPrint()
	fmt.Println()
	fmt.Print("recursionPosPrint: ")
	binaryTree.recursionPosPrint()
	fmt.Println()
	fmt.Print("unRecursionPrePrint: ")
	binaryTree.unRecursionPrePrint()

	fmt.Print("unRecursionMidPrint: ")
	binaryTree.unRecursionMidPrint()

	fmt.Print("unRecursionPosPrint: ")
	binaryTree.unRecursionPosPrint()

	fmt.Print("bfs: ")
	binaryTree.BFS()

	binaryTree.isBST()

	bt7 := &binaryTreeNode{4, nil, nil}
	bt8 := &binaryTreeNode{7, nil, nil}
	bt9 := &binaryTreeNode{13, nil, nil}
	bt4 := &binaryTreeNode{1, nil, nil}
	bt6 := &binaryTreeNode{14, bt9, nil}
	bt5 := &binaryTreeNode{6, bt7, bt8}
	bt3 := &binaryTreeNode{10, nil, bt6}
	bt2 := &binaryTreeNode{3, bt4, bt5}
	bt := &binaryTreeNode{8, bt2, bt3}

	bt.isBST()

	bt.isCBT()
	binaryTree.isCBT()
}

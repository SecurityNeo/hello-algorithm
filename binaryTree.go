package main

import (
	"errors"
	"fmt"
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

func newStack(cap int) (s *stack, err error) {
	if cap > 1024 {
		return nil, errors.New("stack cap: 1024")
	}
	arr := make([]*binaryTreeNode, cap)
	return &stack{cap: cap, topIndex: -1, arr: arr}, nil
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

// recursionPrePrint 二叉树前序遍历（递归）（头左右）
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

// unRecursionPrePrint 二叉树前序遍历（非递归）（深度优先）
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
	for nodeStack.topIndex < nodeStack.cap-1 && nodeStack.topIndex > -1 {
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
}

// unRecursionMidPrint 二叉树中序遍历（非递归）
func (node *binaryTreeNode) unRecursionMidPrint() {
	if node == nil {
		return
	}
	nodeStack, _ := newStack(7)
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
	fmt.Println()
	fmt.Print("unRecursionMidPrint: ")
	binaryTree.unRecursionMidPrint()
}

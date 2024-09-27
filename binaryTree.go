package main

import (
	"container/list"
	"errors"
	"fmt"
	"math"
	"strconv"
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

// getNodeCount 二叉树的节点个数
func (node *binaryTreeNode) getNodeCount() (nodeCount int) {
	if node == nil {
		return 0
	}
	nodeCount = node.left.getNodeCount() + node.right.getNodeCount() + 1
	return
}

// getMaxDepth 二叉树最大深度
func (node *binaryTreeNode) getMaxDepth() (maxDepth int) {
	if node == nil {
		return 0
	}
	leftMaxDepth := node.left.getMaxDepth()
	rightMaxDepth := node.right.getMaxDepth()
	if leftMaxDepth >= rightMaxDepth {
		return leftMaxDepth + 1
	}
	return rightMaxDepth + 1
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

func (node *binaryTreeNode) btPreMarshal() {
	var res string
	var process func(node *binaryTreeNode)
	process = func(node *binaryTreeNode) {
		if node == nil {
			res = res + "#NULL"
		} else {
			res = res + "#" + strconv.Itoa(node.val)
			process(node.left)
			process(node.right)
		}
	}
	process(node)
	fmt.Println(res)
}

func btPreUnMarshal(content, split string) {

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
		fmt.Println("BST(unRecursion): true")
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
				fmt.Println("BST(unRecursion): false")
				return
			}
			head = tmpNode.right
		}
	}
	fmt.Println("BST(unRecursion): true")
}

// isBSTRecursion 判断一个二叉树是否为搜索二叉树的递归实现
func (node *binaryTreeNode) isBSTRecursion() {
	if node == nil {
		fmt.Println("BST(recursion): true")
		return
	}
	if isBst, _, _ := isBSTRecursionProcess(node); isBst {
		fmt.Println("BST(recursion): true")
		return
	}
	fmt.Println("BST(recursion): false")
}

func isBSTRecursionProcess(node *binaryTreeNode) (isBst bool, min int, max int) {
	if node.left == nil && node.right == nil {
		return true, node.val, node.val
	}

	min = node.val
	max = node.val

	var leftIsBst bool
	var leftMin int
	var leftMax int
	var rightIsBst bool
	var rightMin int
	var rightMax int

	if node.left != nil {
		leftIsBst, leftMin, leftMax = isBSTRecursionProcess(node.left)
		min = int(math.Min(float64(min), float64(leftMin)))
		max = int(math.Max(float64(max), float64(leftMax)))
	} else {
		// 使左子树的判断条件成立
		leftIsBst = true
		// leftMax 只做判断使用，值不会返回
		leftMax = node.val - 1
	}

	if node.right != nil {
		rightIsBst, rightMin, rightMax = isBSTRecursionProcess(node.right)
		min = int(math.Min(float64(min), float64(rightMin)))
		max = int(math.Max(float64(max), float64(rightMax)))
	} else {
		// 使右子树的判断条件成立
		rightIsBst = true
		// leftMax 只做判断使用，值不会返回
		rightMin = node.val + 1
	}
	// 左子树的最大值比当前节点值小，右子树的最小值比当前节点值大，并且左右子树都是搜索二叉树
	if leftMax < node.val && node.val < rightMin && leftIsBst && rightIsBst {
		return true, min, max
	} else {
		return false, min, max
	}
}

// isBBT 判断一个二叉树是否为平衡二叉树
// 对于任何一个子树，满足三个条件：
// 1、左子树与右子树都为平衡二叉树；
// 2、它是一颗空树或者它的左子树与右子树的高度差绝对值不超过1
func (node *binaryTreeNode) isBBT() {
	if isB, _ := isBbtProcess(node); isB {
		fmt.Println("BBT: true")
		return
	}
	fmt.Println("BBT: false")
}

func isBbtProcess(node *binaryTreeNode) (isB bool, height int) {
	if node == nil {
		return true, 0
	}
	leftIsB, leftHeight := isBbtProcess(node.left)
	rightIsB, rightHeight := isBbtProcess(node.right)
	if leftHeight >= rightHeight {
		height = leftHeight + 1
	} else {
		height = rightHeight + 1
	}

	if math.Abs(float64(leftHeight-rightHeight)) <= 1 && leftIsB && rightIsB {
		return true, height
	} else {
		return false, height
	}
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

// isFBT 判断一个二叉树是否为满二叉树
// 二叉树最大深度为k，如果其节点数为(2^k)-1，那此树即为满二叉树
func (node *binaryTreeNode) isFBT() {
	if node == nil {
		fmt.Println("FBT: true")
	}
	nodeCount := node.getNodeCount()
	maxDepth := node.getMaxDepth()
	if nodeCount == (1<<maxDepth)-1 {
		fmt.Println("FBT: true")
	} else {
		fmt.Println("FBT: false")
	}
}

// maxDiameter 二叉树最大直径（二叉树的直径指二叉树中任意两个节点之间最长路径的长度，即两个节点之间最大节点数减1）
func (node *binaryTreeNode) maxDiameter() {
	// maxDiameter 最大直径，全局变量
	var maxDiameter int
	// dfs 返回二叉树的最大深度-1
	var dfs func(node *binaryTreeNode) int
	dfs = func(node *binaryTreeNode) int {
		if node == nil {
			return -1
		}
		leftDepth := dfs(node.left) + 1
		rightDepth := dfs(node.right) + 1
		// 左右子树的最大深度之和大于maxDiameter就更新其值
		maxDiameter = max(maxDiameter, leftDepth+rightDepth)
		return max(leftDepth, rightDepth)
	}
	dfs(node)
	fmt.Printf("maxDiameter: %d\n", maxDiameter)
}

// maxPathSum 二叉树最大路径和（No.124）
func (node *binaryTreeNode) maxPathSum() {
	maxPathSum := math.MinInt
	var dfs func(node *binaryTreeNode) int
	dfs = func(node *binaryTreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := dfs(node.left)
		rightSum := dfs(node.right)
		maxPathSum = max(maxPathSum, leftSum+rightSum+node.val)
		// 返回当前子树的最大路径和，如果为负数，返回0，表示当前子树不加人路径
		return max(0, max(leftSum, rightSum)+node.val)
	}
	dfs(node)
	fmt.Printf("maxPathSum: %d\n", maxPathSum)
}

// rob 打家劫舍问题（No.337）
func (node *binaryTreeNode) rob() {
	var dfs func(node *binaryTreeNode) (int, int)
	// 返回的两个值：1、偷当前这一家能获得的最高金额；2、不偷当前这一家能获得的最高金额
	dfs = func(node *binaryTreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		leftRob, leftNoRob := dfs(node.left)
		rightRob, rightNoRob := dfs(node.right)
		// 由于不能偷紧挨着的两家，如果偷当前这一家，左孩子与右孩子这两家就不能偷
		rob := leftNoRob + rightNoRob + node.val
		// 如果不偷当前这一家，那左孩子与右孩子这两家可偷可不偷，按着能获得最高金额的方案来
		noRob := max(leftRob, leftNoRob) + max(rightRob, rightNoRob)
		return rob, noRob
	}
	maxRobCount := max(dfs(node))
	fmt.Printf("maxRobCount: %d\n", maxRobCount)
}

// maxUniValPath 二叉树最大同值路径（No.687）
// 最大同值路径必定为某一个节点的左最大同值路径与右最大同值路径之和
func (node *binaryTreeNode) maxUniValPath() {
	var maxUniValPath int
	var dfs func(node *binaryTreeNode) int
	dfs = func(node *binaryTreeNode) int {
		if node == nil {
			return 0
		}
		// leftVal与rightVal分别表示以当前节点为根节点左右子树的最大同值路径长
		leftVal := dfs(node.left)
		rightVal := dfs(node.right)
		// 存在左或者右孩子，并且孩子的值与当前节点的值不相同，那就无法构成同值路径，就将对应值置为0
		// 为何父结点和子结点值不同，就将对应部分的长度置为0，是因为值不同，已经无法在当前这个局部范围内构成相同值路径了，
		// 同时也不可能和更上一层的结点构成相同值路径了（因为不连续）
		if node.left != nil && node.left.val == node.val {
			leftVal = leftVal + 1
		} else {
			leftVal = 0
		}
		if node.right != nil && node.right.val == node.val {
			rightVal = rightVal + 1
		} else {
			rightVal = 0
		}
		// 最大同值路径必定为某一个节点的左最大同值路径与右最大同值路径之和
		maxUniValPath = max(maxUniValPath, leftVal+rightVal)
		// 对于当前节点来说，其返回的是左或者右的最大同值路径较大的一个，而不是两者之和，因为其父节点有可能也能与其构成同值路径
		return max(leftVal, rightVal)
	}
	dfs(node)
	fmt.Printf("maxUniValPath: %d\n", maxUniValPath)
}

// LCA 最低公共祖先。以下为递归优化解法，思路比较抽象，也可以使用map来记录每个节点的父节点来做。
// 只会存在以下两种情况：
// 1、最低公共祖先是c1或c2；
// 2、最低公共祖先不是是c1或c2，而是在c1、c2的上方。
func (node *binaryTreeNode) LCA(c1, c2 *binaryTreeNode) {
	var process func(node, c1, c2 *binaryTreeNode) *binaryTreeNode
	process = func(node, c1, c2 *binaryTreeNode) *binaryTreeNode {
		// 这个条件即能覆盖情况1
		// 先向下找到c1、c2,然后在上面的节点来判断
		if node == nil || node == c1 || node == c2 {
			return node
		}
		left := process(node.left, c1, c2)
		right := process(node.right, c1, c2)
		// 如果左子树与右子树返回值均不为空，那一定是左右分别发现了c1、c2，这个时候node就是他们的最低公共祖先
		if left != nil && right != nil {
			return node
		}
		if left != nil {
			return left
		} else {
			return right
		}
	}
	lowestPubAncestor := process(node, c1, c2)
	fmt.Printf("lowestPubAncestor: %d\n", lowestPubAncestor.val)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

	bt7 := &binaryTreeNode{4, nil, nil}
	bt8 := &binaryTreeNode{7, nil, nil}
	bt9 := &binaryTreeNode{13, nil, nil}
	bt4 := &binaryTreeNode{1, nil, nil}
	bt6 := &binaryTreeNode{14, bt9, nil}
	bt5 := &binaryTreeNode{6, bt7, bt8}
	bt3 := &binaryTreeNode{10, nil, bt6}
	bt2 := &binaryTreeNode{3, bt4, bt5}
	bt := &binaryTreeNode{8, bt2, bt3}

	bta := &binaryTreeNode{1, nil, nil}
	btb := &binaryTreeNode{1, nil, nil}
	btc := &binaryTreeNode{5, nil, nil}
	btd := &binaryTreeNode{5, nil, btc}
	bte := &binaryTreeNode{4, bta, btb}
	btf := &binaryTreeNode{5, bte, btd}

	bt.isBST()
	bt.isBSTRecursion()
	binaryTree.isBST()
	binaryTree.isBSTRecursion()

	bt.isCBT()
	bt.unRecursionMidPrint()

	maxDepth := bt.getMaxDepth()
	fmt.Printf("MaxDepth: %d\n", maxDepth)

	bt.isFBT()

	binaryTree.isBBT()
	bt.isBBT()

	binaryTree.maxDiameter()
	bt.maxDiameter()

	binaryTree.maxPathSum()
	bt.maxPathSum()

	binaryTree.rob()
	bt.rob()

	btf.maxUniValPath()

	bt.LCA(bt2, bt7)

	binaryTree.btPreMarshal()
}

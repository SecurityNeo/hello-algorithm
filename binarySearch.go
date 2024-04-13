package main

import (
	"fmt"
)

// binarySearchForOrdered 二分查找。非空升序数组，时间复杂度：O(log n)
func binarySearchForOrdered(arr []int, target int) {
	// 边界1： 第一个数即为待查找数
	if arr[0] == target {
		fmt.Printf("Found target,index %d\n", 0)
		return
	}
	lastIndex := len(arr) - 1
	// 边界2：最后一个数即为待查找数
	if arr[lastIndex] == target {
		fmt.Printf("Found target,index %d\n", lastIndex)
		return
	}

	lowIndex := 0
	highIndex := lastIndex

	for lowIndex <= highIndex {
		// midIndex := (lowIndex + highIndex) / 2
		// 防止 lowIndex + highIndex 溢出
		midIndex := lowIndex + ((highIndex - lowIndex) >> 1)
		if arr[midIndex] == target {
			fmt.Printf("Found target,index %d\n", midIndex)
			return
		} else if arr[midIndex] > target {
			highIndex = midIndex - 1
		} else if arr[midIndex] < target {
			lowIndex = midIndex + 1
		}
	}
	fmt.Printf("Not found target\n")
}

// extremumMin 二分查找。获取一个无序非空数组（数组中任意相邻两个元素不相等）的极小值(局部最小值)。时间复杂度：O(log n)
func extremumMin(arr []int) {
	// 边界1：数组就一个元素，直接返回
	if len(arr) == 1 {
		fmt.Printf("Got a minimum of the array，index： %d\n", 0)
		return
	}

	// 边界2：数组第一个元素比第二个元素小，直接返回
	if arr[0] < arr[1] {
		fmt.Printf("Got a minimum of the array，index： %d\n", 0)
		return
	}

	// 边界3：数组倒数第一个元素比倒数第二个元素小，直接返回
	if arr[len(arr)-1] < arr[len(arr)-2] {
		fmt.Printf("Got a minimum of the array，index： %d\n", len(arr)-1)
		return
	}

	leftIndex := 1
	rightIndex := len(arr) - 2

	// 在不满足边界2与3的情况下，leftIndex与rightIndex之间必定存在局部最小值，按照二分查找的思路进行后续查找逻辑。
	for leftIndex < rightIndex {
		// midIndex := (leftIndex + rightIndex) / 2
		// 防止 leftIndex + rightIndex 溢出
		midIndex := leftIndex + ((rightIndex - leftIndex) >> 1)
		// 局部最小值定义即为查找条件
		if arr[midIndex] < arr[midIndex-1] && arr[midIndex] < arr[midIndex+1] {
			fmt.Printf("Got a minimum of the array，index： %d\n", midIndex)
			return
		}

		/*如果不满足上面的条件，接着查找，arr[midIndex] > arr[midIndex-1]，说明左边上行状态，
		  左边肯定存在一个局部最小值，将右侧下标移到midIndex-1处 */
		if arr[midIndex] > arr[midIndex-1] {
			rightIndex = midIndex - 1
		} else {
			leftIndex = midIndex + 1
		}
	}

	// 程序到这里，说明leftIndex与rightIndex区间只有两个数了
	if arr[leftIndex] < arr[rightIndex] {
		fmt.Printf("Got a minimum of the array，index： %d\n", leftIndex)
		return
	} else {
		fmt.Printf("Got a minimum of the array，index： %d\n", rightIndex)
	}
}

// queryLeftIndex 二分查找。查找一个非空有序数组arr中大于等于target并且最靠左的位置
func queryLeftIndex(arr []int, target int) {
	// 边界1： 数组只有一个元素，且其值小于target
	if len(arr) == 1 {
		if arr[0] < target {
			fmt.Printf("Not found\n")
			// 边界2： 数组只有一个元素，且其值大于等于target
		} else {
			fmt.Printf("Got the number: %d, index: %d\n", arr[0], 0)
		}
		return
	}

	leftIndex := 0
	rightIndex := len(arr) - 1
	for leftIndex < rightIndex {
		// midIndex := (leftIndex + rightIndex) / 2
		// 防止 leftIndex + rightIndex 溢出
		midIndex := leftIndex + ((rightIndex - leftIndex) >> 1)
		if arr[midIndex] >= target {
			rightIndex = midIndex
		}
		if arr[midIndex] < target {
			leftIndex = midIndex + 1
		}
	}
	if arr[leftIndex] >= target {
		fmt.Printf("Got the number: %d, index: %d\n", arr[leftIndex], leftIndex)
	} else {
		fmt.Printf("Got the number: %d, index: %d\n", arr[rightIndex], rightIndex)
	}
}

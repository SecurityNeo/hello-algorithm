package main

import (
	"fmt"
)

// InsertSort 插入排序（每次将一个待排序的元素与已排序的元素进行逐一比较，直到找到合适的位置按大小插入）,
// 时间复杂度：O(n^2),适用于小规模数据和部分有序数据的排序需求。对于小规模数据和部分有序的数据，插入排序表现良好。
func InsertSort(arr []int) {
	if len(arr) > 1 {
		for i := 1; i < len(arr); i++ {
			for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
				swap(arr, j, j+1)
			}
		}
	}
	fmt.Println(arr)
}

// swap 数据交换
func swap(arr []int, i, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}

// mergeSort 归并排序+递归。时间复杂度O(N*logN)。
// 归并排序是将已有序的子序列合并为一个有序的完整序列。每次从两个数组中选择值较小的元素放到另一个辅助数组中。
func mergeSort(arr []int) {
	if len(arr) <= 1 {
		fmt.Println("Sorted, array: ", arr)
		return
	}

	// 此处调用一个递归函数
	mergeSortProcess(arr, 0, len(arr)-1)

	fmt.Println("Sorted, array: ", arr)
}

// mergeSortProcess 递归函数，实现将arr数组的L至R位置进行排序
func mergeSortProcess(arr []int, L, R int) {
	if L == R {
		return
	}
	// M为数组中点
	M := L + ((R - L) >> 1)
	// 实现L至M（数组左侧）排序
	mergeSortProcess(arr, L, M)
	// 实现M+1至R（数组右侧）排序
	mergeSortProcess(arr, M+1, R)
	// 将两个有序数组和并
	mergeArr(arr, L, M, R)
}

func mergeArr(arr []int, L, M, R int) {
	length := R - L + 1
	// 定义一个辅助数组，长度为待排序元素的个数
	helpArr := make([]int, length)
	// i为辅助数组待插入元素的位置
	i := 0
	// p1为arr数组中左侧待比较元素的位置
	p1 := L
	// p2为arr数组中右侧待比较元素的位置，M为数组中点，注意它不是整个arr数组的中点，而是arr数组中待排序那部分（L至R）的中点
	p2 := M + 1
	for p1 <= M && p2 <= R {
		if arr[p1] <= arr[p2] {
			helpArr[i] = arr[p1]
			p1++
			i++
		} else {
			helpArr[i] = arr[p2]
			p2++
			i++
		}
	}
	// 程序走到这里，一定是p1或者p2越界了，将左侧剩余元素放到辅助数组中，此时剩余元素本身是有序的，直接赋值，无需再比较
	for p1 <= M {
		helpArr[i] = arr[p1]
		i++
		p1++
	}
	for p2 <= R {
		helpArr[i] = arr[p2]
		i++
		p2++
	}
	// 将辅助数组元素全部拷贝至原始数组中
	for i := 0; i < len(helpArr); i++ {
		arr[L+i] = helpArr[i]
	}
}

// smallSum 利用归并排序思路解决求小和问题（在一个数组中，每一个元素左边比当前元素值小的元素值累加起来，叫做这个数组的小和）
func smallSum(arr []int) {
	if len(arr) < 2 {
		fmt.Println("Small Sum: 0")
		return
	}
	sum := smallSumProcess(arr, 0, len(arr)-1)
	fmt.Println("Array: ", arr)
	fmt.Printf("Small Sum: %d\n", sum)
}

func smallSumProcess(arr []int, L, R int) (sum int) {
	if L == R {
		return 0
	}
	M := L + ((R - L) >> 1)
	return smallSumProcess(arr, L, M) + smallSumProcess(arr, M+1, R) + smallSumMerge(arr, L, M, R)
}

func smallSumMerge(arr []int, L, M, R int) (sum int) {
	helpArrLength := R - L + 1
	helpArr := make([]int, helpArrLength)
	sum = 0
	i := 0
	p1 := L
	p2 := M + 1
	for p1 <= M && p2 <= R {
		if arr[p1] < arr[p2] {
			sum += (R - p2 + 1) * arr[p1]
			helpArr[i] = arr[p1]
			i++
			p1++
		} else {
			helpArr[i] = arr[p2]
			i++
			p2++
		}
	}
	for p1 <= M {
		helpArr[i] = arr[p1]
		i++
		p1++
	}
	for p2 <= R {
		helpArr[i] = arr[p2]
		i++
		p2++
	}
	for i := 0; i < len(helpArr); i++ {
		arr[L+i] = helpArr[i]
	}
	return sum
}

// inversePair 利用归并排序思路求逆序对问题。对于数组（所有数字不相同）中一个数A，如果在其右边的数比它小，那这两个数称为一个逆序对。
// 一个升序数组没有逆序对
func inversePair(arr []int) {
	if len(arr) <= 1 {
		fmt.Println("inverse pair count: 0")
		return
	}
	count := inversePairProcess(arr, 0, len(arr)-1)
	fmt.Printf("inverse pair count: %d\n", count)
}

func inversePairProcess(arr []int, L, R int) (count int) {
	if L == R {
		return 0
	}
	M := L + ((R - L) >> 1)
	count = inversePairProcess(arr, L, M) + inversePairProcess(arr, M+1, R) + inversePairMerge(arr, L, M, R)
	return count
}

func inversePairMerge(arr []int, L, M, R int) (count int) {
	arrLength := R - L + 1
	helpArr := make([]int, arrLength)

	p1 := L
	p2 := M + 1

	i := 0

	for p1 <= M && p2 <= R {
		if arr[p1] < arr[p2] {
			helpArr[i] = arr[p1]
			i++
			p1++
		} else {
			helpArr[i] = arr[p2]
			count += M - p1 + 1
			i++
			p2++
		}
	}

	for p1 <= M {
		helpArr[i] = arr[p1]
		p1++
		i++
	}
	for p2 <= R {
		helpArr[i] = arr[p2]
		p2++
		i++
	}
	return count
}

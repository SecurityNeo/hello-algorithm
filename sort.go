package main

import (
	"fmt"
	"math/rand"
	"time"
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

// quickSort1 快排前序。小于等于num的放左边，大于num的放右边。（快慢指针）
func quickSort1(arr []int, num int) {
	fmt.Println("Before sorted: ", arr)
	if len(arr) <= 1 {
		fmt.Printf("No need to sort,array length: %d\n", len(arr))
		return
	}

	i := -1
	j := 0

	for j < len(arr) {
		if arr[j] <= num {
			arr[j], arr[i+1] = arr[i+1], arr[j]
			i++
			j++

		} else {
			j++
		}
	}
	fmt.Println("quickSort1 result: ", arr)
}

// quickSort2 荷兰国旗问题（只做分区，不用保证整个数组有序）。
func quickSort2(arr []int, num int) {
	fmt.Println("Before sorted: ", arr)
	arrLength := len(arr)
	if arrLength <= 1 {
		fmt.Printf("No need to sort,array lenth: %d\n", arrLength)
		return
	}

	p1 := -1
	p2 := arrLength
	cur := 0

	for cur < p2 {
		if arr[cur] < num {
			arr[cur], arr[p1+1] = arr[p1+1], arr[cur]
			cur++
			p1++
		} else if arr[cur] == num {
			cur++
		} else {
			// 当当前值比基准值大的时候，标识当前值的指针不能自增，否则当arr[p2-1]的值也比基准值大的时候就乱了。
			arr[cur], arr[p2-1] = arr[p2-1], arr[cur]
			p2--
		}
	}
	fmt.Println("quickSort2 result: ", arr)
}

// quickSort3 快排（版本1）。思路：参照荷兰国旗问题思路，选取数组中的一个数作为基准元素，比基准元素小的放一个临时数组，
// 与基准元素相等的放一个临时数组， 比基准元素大的放另一个临时数组，然后将这三个临时数组合并，递归调用，即可以得到一个排序后的数组。
// 此种方法空间复杂度高。
// 注意：必须要设计一个与基准元素相等的临时数组，否则在遇到原始数组有多个相等的元素，并且这个元素在最后一次分区时被选为基准元素后，迭代将永远不会结束
func quickSort3(arr []int) {
	fmt.Println("Before sorted: ", arr)
	if len(arr) <= 1 {
		fmt.Printf("No need to sort,array lenth: %d\n", len(arr))
		return
	}
	result := quickSort3Process(arr)
	fmt.Println("quickSort3 result: ", result)
}

func quickSort3Process(arr []int) (Arr []int) {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	// 随机取数组中的一个作为基准数，防止特殊情况下（例如待排序的数组就是一个有序数组）复杂度退化为O(n^2)
	rand.Seed(time.Now().UnixNano())
	p := rand.Intn(length)
	pivot := arr[p]
	var lowArr, equal, highArr []int
	for _, v := range arr {
		if v < pivot {
			lowArr = append(lowArr, v)
		} else if v == pivot {
			equal = append(equal, v)
		} else {
			highArr = append(highArr, v)
		}
	}

	return append(quickSort3Process(lowArr), append(equal, quickSort3Process(highArr)...)...)
}

// quickSort4 快排（版本2）。参照荷兰国旗问题解题思路，先依据基准值给数组分为三个部分，再分别把小于基准值和大于基准值的部分继续迭代排序。
// 相比版本1，空间复杂度低一些。
func quickSort4(arr []int) {
	if len(arr) <= 1 {
		fmt.Printf("No need to sort,array lenth: %d\n", len(arr))
		return
	}
	quickSort4Process(arr, 0, len(arr)-1)
	fmt.Println("quickSort4 result: ", arr)
}

func quickSort4Process(arr []int, left, right int) {
	if left < right {
		rand.Seed(time.Now().UnixNano())
		// 基准数的位置必须在left与right之间
		p := rand.Intn(right-left) + left
		pivot := arr[p]
		p1 := left - 1
		p2 := right + 1
		cur := left
		for cur < p2 {
			if arr[cur] < pivot {
				arr[cur], arr[p1+1] = arr[p1+1], arr[cur]
				p1++
				cur++
			} else if arr[cur] == pivot {
				cur++
			} else {
				arr[cur], arr[p2-1] = arr[p2-1], arr[cur]
				p2--
			}
		}
		// 边界条件：如果p1为-1，那小于基准值的部分为空，就无需再迭代
		if p1 != -1 {
			quickSort4Process(arr, 0, p1)
		}
		// 边界条件：如果p2为right+1，那么大于基准值的部分为空，同样无需再迭代
		if p2 != right+1 {
			quickSort4Process(arr, p2, right)
		}
	}
}

// bubbleSort 冒泡排序。
func bubbleSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		fmt.Printf("No need to sort,array lenth: %d\n", len(arr))
		return
	}

	// 通过找到数组中处于无序部分的右边界来减少比较的范围，在一些场景下可以极大提高算法效率
	// 表示数组中处于无序部分的右边界
	sortBorder := length - 1
	// 表示每一轮次产生数据交换时的索引
	lastExchangeIndex := 0

	// 外层循环控制轮次
	for i := 0; i < length-1; i++ {
		// 数组是否有序的标识，如果在中间轮次已有序，那就直接结束。
		// 此标识初始值设为true，如果在某一轮次中，发生过数据位置交换，就表示此数组还不是有序状态，
		// 反之，这一轮次都没有出现数据交换机，那就表示数据已有序。
		isSorted := true
		// 内层循环控制比较的元素边界
		// 基础排序算法实现的版本里，j < length-i-1
		for j := 0; j < sortBorder; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSorted = false
				lastExchangeIndex = j
			}
		}
		// 将最后发生数据交换的索引赋值给sortBorder，这样在下一轮次循环里，待比较的数组元素由边界就是上一轮次最后交换数据的地方
		sortBorder = lastExchangeIndex
		// 如果已有序，则退出循环
		if isSorted {
			break
		}
	}
	fmt.Println("bubbleSort result: ", arr)
}

// selectSort 选择排序。时间复杂度O(n^2)，由于选择元素后会发生交换，有可能把前面的元素交换到后面，
// 所以选择排序是不稳定排序（如果a原本在b的前面，且a == b，排序之后a仍然在b的前面，则为稳定排序）。
func selectSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		fmt.Printf("No need to sort,array lenth: %d\n", len(arr))
		return
	}
	// 从无序区间不断挑选最小值与有序区间的最后一个元素交换
	for i := 0; i < length-1; i++ {
		// 假定无序区间第一个元素是此区间的最小值
		minIndex := i
		// 遍历无序区间，如果有比无序区间第一个元素更小的元素，将其索引赋值给minIndex
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 找出无序区间最小元素后，与第一个元素交换
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	fmt.Println("selectSort result: ", arr)
}

// heapInsert 实现对于一个大根堆arr，插入一个元素至堆尾，最终仍然为一个大根堆。
func heapInsert(arr []int, heapSize int) {
	// i 取最后一个节点索引
	i := heapSize
	// 对于一个完全二叉树，其中一个元素位置为 i，则其左儿子位置为 2*i+1，右儿子位置为 2*i+2，其父亲的位置为 (i-1)/2
	for arr[i] > arr[(i-1)/2] {
		arr[i], arr[(i-1)/2] = arr[(i-1)/2], arr[i]
		i = (i - 1) / 2
	}

}

// prepareHeap 构造一个大根堆
func prepareHeap(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		heapInsert(arr, i)
	}
}

// heapify 在一个大根堆中，去掉最大值（即arr[0]）,并要求剩余的元素仍然为一个大根堆
func heapify(arr []int, heapSize int) {
	// 先将第一个元素（大根堆中的最大值）与数组最后一个值交换
	arr[0], arr[heapSize] = arr[heapSize], arr[0]
	// 交换之后把最后一个元素（即之前数组中的最大值）排除在外
	lastIndex := heapSize - 1
	i := 0
	for 2*i+1 <= lastIndex { // 至少有一个孩子
		var largestIndex int
		if 2*i+2 <= lastIndex { // 存在右孩子
			if arr[i] >= arr[2*i+1] && arr[i] >= arr[2*i+2] {
				break
			} else if arr[2*i+1] > arr[i] && arr[2*i+1] > arr[2*i+2] {
				largestIndex = 2*i + 1
				arr[i], arr[2*i+1] = arr[2*i+1], arr[i]
				i = largestIndex
			} else {
				largestIndex = 2*i + 2
				arr[i], arr[2*i+2] = arr[2*i+2], arr[i]
				i = largestIndex
			}
		} else { // 只有左孩子
			if arr[i] >= arr[2*i+1] {
				break
			}
			arr[i], arr[2*i+1] = arr[2*i+1], arr[i]
			i = 2*i + 1
		}
	}
}

func heapSort(arr []int) {
	prepareHeap(arr)
	length := len(arr)
	for i := length - 1; i >= 0; i-- {
		heapify(arr, i)
	}
	fmt.Println("heapSort result:   ", arr)
}

// sortK 堆排序扩展。对于一个几乎已排好序的数组，数组中每个元素的位置跟最终排序完的位置偏移量最大不超过k
func sortK(arr []int, k int) {
	helpArr := prepareSmallHeap(arr, k+1)
	for i := 1; i < len(arr)-k; i++ {
		heapifySmall(helpArr, arr[i+k])
		arr[i] = helpArr[0]
	}
	// 将小根堆的元素做一次堆排序
	heapSort(helpArr)
	arr = append(arr[:len(arr)-k-1], helpArr...)
	fmt.Println("sortK result:      ", arr)
}

// prepareSmallHeap 构造一个小根堆
func prepareSmallHeap(arr []int, heapSize int) (heapArr []int) {
	for i := 0; i < heapSize; i++ {
		// 对于一个完全二叉树，其中一个元素位置为 i，则其左儿子位置为 2*i+1，右儿子位置为 2*i+2，其父亲的位置为 (i-1)/2
		for arr[i] < arr[(i-1)/2] {
			arr[i], arr[(i-1)/2] = arr[(i-1)/2], arr[i]
			i = (i - 1) / 2
		}
	}
	heapArr = append(heapArr, arr[:heapSize]...)
	return
}

// heapifySmall 替换小根堆中第一个节点后，保证其仍然为一个小根堆
func heapifySmall(helpArr []int, newNode int) {
	helpArr[0] = newNode
	length := len(helpArr)
	i := 0
	for 2*i+1 <= length-1 { // 至少有一个孩子
		if 2*i+2 <= length-1 { // 存在右孩子
			if helpArr[i] < helpArr[2*i+1] && helpArr[i] < helpArr[2*i+2] {
				break
			} else if helpArr[2*i+1] < helpArr[i] && helpArr[2*i+1] < helpArr[2*i+2] {
				helpArr[2*i+1], helpArr[i] = helpArr[i], helpArr[2*i+1]
				i = 2*i + 1
			} else {
				helpArr[2*i+2], helpArr[i] = helpArr[i], helpArr[2*i+2]
				i = 2*i + 2
			}
		} else { // 只有左孩子
			if helpArr[2*i+1] < helpArr[i] {
				helpArr[2*i+1], helpArr[i] = helpArr[i], helpArr[2*i+1]
				i = 2*i + 1
			} else {
				break
			}
		}
	}
}

// topK 堆排序扩展，获取数组中前k个最大或最小的元素。求前k个最大元素，构造长度为k的小根堆，反之构造大根堆。
// 然后用剩余的（n-k）个元素依次与堆顶元素比较，符合条件（求最大k个元素的话就看其是否比堆顶元素大），就交换并堆化。
func topK(arr []int, k int) {
	var topHeap []int
	// step1：先构造一个小根堆
	for i := 0; i < k; i++ {
		for arr[i] < arr[(i-1)/2] {
			arr[i], arr[(i-1)/2] = arr[(i-1)/2], arr[i]
			i = (i - 1) / 2
		}
	}
	topHeap = append(topHeap, arr[:k]...)

	// step2: 不断取原数组中剩余的元素与topHeap中的堆顶元素比较，如果比堆顶元素大，就用其替换掉堆顶元素，再参照heapify流程使topHeap为一个小根堆
	for i := k - 1; i < len(arr); i++ {
		if arr[i] > topHeap[0] {
			topHeap[0] = arr[i]
			smallestIndex := 0
			for 2*smallestIndex+1 <= k-1 { // 至少有一个孩子
				if 2*smallestIndex+2 <= k-1 { // 还有右孩子
					if topHeap[smallestIndex] < topHeap[2*smallestIndex+1] && topHeap[smallestIndex] < topHeap[2*smallestIndex+2] {
						break
					} else if topHeap[2*smallestIndex+1] < topHeap[smallestIndex] && topHeap[2*smallestIndex+1] < topHeap[2*smallestIndex+2] {
						topHeap[smallestIndex], topHeap[2*smallestIndex+1] = topHeap[2*smallestIndex+1], topHeap[smallestIndex]
						smallestIndex = 2*smallestIndex + 1
					} else {
						topHeap[smallestIndex], topHeap[2*smallestIndex+2] = topHeap[2*smallestIndex+2], topHeap[smallestIndex]
						smallestIndex = 2*smallestIndex + 2
					}
				} else { // 只有左孩子
					if topHeap[smallestIndex] < topHeap[2*smallestIndex+1] {
						break
					} else {
						topHeap[smallestIndex], topHeap[2*smallestIndex+1] = topHeap[2*smallestIndex+1], topHeap[smallestIndex]
						smallestIndex = 2*smallestIndex + 1
					}
				}
			}
		}
	}
	fmt.Println("topK result:    ", topHeap)
}

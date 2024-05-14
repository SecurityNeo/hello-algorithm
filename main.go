package main

func main() {

	Swap(100, 200)

	ExtractLastOne(52)

	arr := []int{1, 52, 52, 7, 1, 6, 6, 7, 10}
	ExtractOneTimes(arr)

	arr2 := []int{1, 52, 52, 7, 1, 6, 6, 7, 11, 13, 13, 12}
	ExtractTwoTimes(arr2)

	arr3 := []int{5, 3, 1, 8, 6, 9}
	InsertSort(arr3)

	arr4 := []int{1, 3, 5, 7, 10, 14, 17, 19}
	binarySearchForOrdered(arr4, 17)

	arr5 := []int{9, 8, 6, 7, 8, 9, 8, 9}
	extremumMin(arr5)

	arr6 := []int{1, 2, 2, 3, 4, 4, 4, 6, 9, 9, 9, 10}
	queryLeftIndex(arr6, 9)

	arr7 := []int{5, 3, 1, 8, 6, 9}
	mergeSort(arr7)

	arr8 := []int{5, 3, 1, 8, 6, 9}
	smallSum(arr8)

	arr9 := []int{5, 3, 1, 8, 6, 9}
	inversePair(arr9)

	arr10 := []int{5, 3, 1, 8, 6, 9}
	quickSort1(arr10, 7)

	arr11 := []int{2, 3, 1, 5, 4, 6, 3, 2, 7, 3, 2, 4, 1}
	quickSort2(arr11, 3)

	arr12 := []int{12, 87, 1, 66, 30, 126, 328, 12, 653, 67, 98, 3, 256, 5, 1, 1, 99, 109, 17, 70, 4}
	quickSort3(arr12)

	arr13 := []int{12, 87, 1, 66, 30, 126, 328, 12, 653, 67, 98, 3, 256, 5, 1, 1, 99, 109, 17, 70, 4}
	quickSort4(arr13)

	arr14 := []int{12, 87, 1, 66, 30, 126, 328, 12, 653, 67, 98, 3, 256, 5, 1, 1, 99, 109, 17, 70, 4}
	bubbleSort(arr14)

	arr15 := []int{12, 87, 1, 66, 30, 126, 328, 12, 653, 67, 98, 3, 256, 5, 1, 1, 99, 109, 17, 70, 4}
	selectSort(arr15)
}

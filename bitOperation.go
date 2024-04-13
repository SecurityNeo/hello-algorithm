package main

import "fmt"

/*
	异或（^）
		运算：当二进制位不同时，结果为1，否则为0 （不进位相加）
		特性：
			1、无序性（交换律与结合律）：a ^ b ^ c = a ^ c ^ b; a ^ b ^ c = a ^ (b ^ c)
			2、任何值与自身异或，结果为0，任何值与0异或，结果为其自身：0 ^ N = N; N ^ N = 0
			3、可移项行 ：a ^ b = c ^ d => a ^ c = b ^ d;
    按位与(&)
		运算：同时为1，结果为1，否则为0
	按位或(|)
		运算：有一个为1，结果为1，否则为0
*/

// Swap 存在两个整数a、b，a与b不指向同一个内存块，在不借助额外变量的情况下交换a与b的值
func Swap(a, b int) {
	fmt.Printf("before swap: a = %d,b = %d\n", a, b)
	// a = 甲；b = 已

	a = a ^ b
	// a = 甲 ^ 已；  b = 已

	b = a ^ b
	//  a = 甲 ^ 已;   b =  甲 ^ 已 ^ 已 = 甲

	a = a ^ b
	// a = 甲 ^ 已 ^ 甲 = 已

	fmt.Printf("after swap: a = %d,b = %d\n", a, b)
}

// ExtractLastOne 提取一个整数最右侧的1
func ExtractLastOne(a uint) {
	/*
			假定一个数a为52,
			     a:  1 1 0 1 0 0
			    ^a:  0 0 1 0 1 1
		      ^a+1:  0 0 1 1 0 0
		a & (^a+1):  0 0 0 1 0 0
	*/

	b := a & (^a + 1)
	fmt.Printf("%b\n", b)
}

// ExtractOneTimes 存在一个整型切片，其中仅有一个数出现了奇数次，获取这个数，要求时间复杂度O(n)，额外空间O(1)
func ExtractOneTimes(arr []int) {
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor = eor ^ arr[i]
	}
	fmt.Println(eor)
}

// ExtractTwoTimes 题目4：存在一个整型切片，其中仅有两个数出现了奇数次，获取这两个数，要求时间复杂度O(n)，额外空间O(1)
func ExtractTwoTimes(arr []int) {
	/*
		1、假设这两个数为a和b，定义一个额外变量eor，值为0，参考第三题，用eor与切片中所有的值做一次异或操作，这样eor=a^b;
		2、a与b是两个数，即说明a≠b，eor≠0，那eor必然有一个二进制位为1，,那a与b对应的二进制位必然不相等；
		3、假设，eor第三位为1，a第三位为1，b第三位为0，我们可以将这个切片分为两类，一类为第三位为1的（a在这一类），一类为第三位为0的（b在这一类）；
		4、再用一个额外变量onlyOne（值为0）与第一类的所有数异或，得到的就是a
	*/

	// 假设 这两个数为a、b
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor = eor ^ arr[i]
	}
	// 到这里 eor = a ^ b

	// 参考第2题（ExtractLastOne）,获取eor最后侧为1的数
	lastOne := eor & (^eor + 1)

	onlyOne := 0
	for i := 0; i < len(arr); i++ {
		// eor最后侧为1的数来与切片中所有的数进行与运算（&），如果这个位置都为1，则取一次异或
		if lastOne&arr[i] == 0 {
			onlyOne = onlyOne ^ arr[i]
		}
	}
	fmt.Printf("%d %d\n", onlyOne, eor^onlyOne)
}

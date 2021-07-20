// 在 Go 中，数组 是一个固定长度单一类型的数列（在典型的 Go 程序中，相对于数组而言，slice 使用的更多）
// 参考：https://blog.golang.org/slices-intro#TOC_3.

package main

import "fmt"

func main() {

	// 数组类型定义需要指定长度和元素类型，例如，该类型[4]int表示一个由四个整数组成的数组
	// 数组的大小是固定的，它的长度是其类型的一部分（[4]int并且[5]int是不同的、不兼容的类型）
	// 数组可以用通常的方式索引，所以表达式s[n]访问第 n 个元素，从零开始
	// 数组默认是零值的,对于 int 数组来说也就是 0

	// 声明数组(声明数组的语法要定义长度与类型)
	var test1 [4]int
	fmt.Println("test1 数组：", test1)
	// 数组不需要显式初始化；数组的零值是一个随时可用的数组，其元素本身已归零
	// a[2] == 0

	// 改
	test1[0] = 1
	fmt.Println("test1 数组赋值：", test1)
	// 查
	fmt.Println("test1 数组索引0的值为：", test1[0])

	// 声明并赋值数组
	// 方式一
	test2 := [2]string{"apple", "orange"}
	fmt.Println("test2 数组：", test2)
	// 方式二：您可以让编译器为您计算数组元素，在这两种情况下，类型b都是[2]string
	test3 := [...]string{"apple", "orange"}
	fmt.Println("test3 数组：", test3)

	// 数组的存储类型是单一的，但是你可以组合这些数组来构造多维的数据结构
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	// 注意，在使用 fmt.Println 来打印数组的时候，会使用[v1 v2 v3 ...] 的格式显示
	fmt.Println("数组组合:", twoD)
}

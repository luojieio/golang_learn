// 序列，即数组，数组的存储类型是单一的，但是你可以组合这些数据来构造多维的数据结构
// 在 Go 中，数组 是一个固定长度的数列
// 2个缺陷：（在典型的 Go 程序中，相对于数组而言，slice 使用的更多）
// - 长度固定
// - 单一类型的值

package main

import "fmt"

func main() {
	// 这里我们创建了一个数组 a 来存放刚好 5 个 int。元素的类型和长度都是数组类型的一部分。数组默认是零值的，对于 int 数组来说也就是 0
	// 声明数组的语法要定义长度与类型
	var a [5]int
	fmt.Println("emp:", a)

	// 数组修改值和获取值
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 使用内置函数 len 返回数组的长度
	fmt.Println("len:", len(a))

	// 使用这个语法在一行内初始化一个数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 数组的存储类型是单一的，但是你可以组合这些数据来构造多维的数据结构
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	// 注意，在使用 fmt.Println 来打印数组的时候，会使用[v1 v2 v3 ...] 的格式显示
	fmt.Println("2d:", twoD)
}

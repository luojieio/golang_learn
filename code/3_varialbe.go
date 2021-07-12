// 变量

package main

import "fmt"

func main() {
	// 在go中定义一个变量包括2个操作：声明变量、赋值变量，有4种方式
	// 方式1，声明与赋值分成2行
	var name1 string
	name1 = "luojie1"
	fmt.Println(name1)

	// 方式2，声明与赋值1行
	var name2 string = "luojie2"
	fmt.Println(name2)

	// 方式3，动态类型
	var name3 = "luojie3"
	fmt.Println(name3)

	// 方式4，精简模式
	name4 := "luojie4"
	fmt.Println(name4)

	// 一次性声明多个变量
	var b, c = "luojie1", "luojie2"
	fmt.Println(b, c)

	// 声明变量未赋值（声明变量且没有给出对应的初始值时，变量将会初始化为零值）
	var d int
	fmt.Println(d)

}

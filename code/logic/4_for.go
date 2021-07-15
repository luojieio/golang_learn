// for 循环
// for 是 Go 中唯一的循环结构。这里有 for 循环的三个基本使用方式。

package main

import "fmt"

func main() {

	// 方式1
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// 方式2，经典的初始化/条件/后续形式 for 循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 方式3，不带条件的 for 循环将一直执行，直到在循环体内使用了 break 或者 return 来跳出循环，实现while循环的功能
	for {
		fmt.Println("loop")
		break
	}
	// 或者这样
	for true {
		fmt.Println("loop")
		break
	}
}

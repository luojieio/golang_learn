// 逻辑判断
// if 和 else 分支结构在 Go 中当然是直接了当的了

package main

import "fmt"

func main() {

	// 注意，在 Go 中，你可以不适用圆括号，但是花括号是需要的。类比python中用:冒号展开，go中用花括号展开
	// Go 里没有三目运算符，所以即使你只需要基本的条件判断，你仍需要使用完整的 if 语句
	// 你可以不要 else 只用 if 语句
	if 8%4 == 0 {
		fmt.Println("8整除4")
	}

	// 可以有else
	if 7%2 == 0 {
		fmt.Println("7是偶数")
	} else {
		fmt.Println("7是奇数")
	}

	// 可以有else if ，在条件语句之前可以有一个语句；任何在这里声明的变量都可以在所有的条件分支中使用
	if num := 9; num < 0 {
		fmt.Println(num, "是负数")
	} else if num < 10 {
		fmt.Println(num, "有1个数")
	} else {
		fmt.Println(num, "有多个数")
	}
}

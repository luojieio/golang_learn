// 逻辑判断，分支结构
// switch ，方便的条件分支语句

package main

import (
	"fmt"
	"time"
)

func main() {

	// 一个基本的switch
	i := 2
	fmt.Println("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 在一个case语句中，你可以使用逗号来分隔多个表达式
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	// 我们很好的使用了可选的default 分支
	default:
		fmt.Println("It's a weekday")
	}

	// 不带表达式的 switch 是实现 if/else 逻辑的另一种方式。这里展示了 case 表达式是如何使用非常量的
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

}

// map 是 Go 内置关联数据类型（在一些其他的语言中称为哈希 或者字典 ）

package main

import "fmt"

func main() {
	// 声明map：要创建一个空 map，需要使用内建的 make函数
	// make(map[key-type]val-type)
	m := make(map[string]int)

	// 声明并赋值map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// map操作：
	// 1.增加与修改：赋值
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// 2.删除：内建的 delete 可以从一个 map 中移除键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	// 3.查看
	fmt.Println("map:", m["k1"])

	// 4.当对一个 map 调用内建的 len 时，返回的是键值对数目
	fmt.Println("len", len(m))

	// 5.当从一个 map 中取值时，可选的第二返回值指示这个键是在这个 map 中。这可以用来消除键不存在和键有零值，像 0 或者 "" 而产生的歧义
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

}

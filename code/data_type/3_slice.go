// Slice 是 Go 中一个关键的数据类型，是一个比数组更加强大的序列接口，切片类型是建立在 Go 数组类型之上的抽象，因此要理解切片我们必须先了解数组。
// 数组有其一席之地，但它们有点不灵活，因此您在 Go 代码中不会经常看到它们，然而，切片无处不在。它们建立在阵列上以提供强大的功能和便利

package main

import "fmt"

func main() {

	// 切片的类型规范是[]T，其中T是切片元素的类型。与数组类型不同，切片类型没有指定的长度
	// 切片的声明就像数组一样，只是省略了元素计数
	// 切片的零值是nil。对于 nil 切片，len和cap函数都将返回 0

	// 一、声明切片
	// 可以使用命名make的内置函数创建切片，make函数的语法为：func make([]T, len, cap)
	// 该make函数采用类型、长度和可选容量，其中 T 代表要创建的切片的元素类型，调用时，make分配一个数组并返回一个引用该数组的切片。
	// 切片的初始化零值
	// 写法1：
	var test1 []byte
	test1 = make([]byte, 5, 10)
	fmt.Println("test1:", test1)
	fmt.Println("len:", len(test1))
	fmt.Println("cap:", cap(test1))

	// 写法2：当容量参数被省略时，它默认为指定的长度。这是相同代码的更简洁版本：
	test2 := make([]byte, 5)
	fmt.Println("test2:", test2)

	// 二、声明并赋值切片
	test3 := []string{"a", "b", "c", "d"}
	fmt.Println("test3:", test3)

	// 三、切片操作
	// 切片也可以通过“切片”现有切片或数组来形成。切片是通过指定一个半开范围来完成的，其中两个索引以冒号分隔。例如，该表达式b[1:4]创建一个包含元素 1 到 3 的b切片（结果切片的索引将为 0 到 2）。
	test4 := []string{"g", "o", "l", "a", "a", "n", "g"}
	fmt.Println(test4[1:4])

	// 切片不会复制切片的数据。它创建一个指向原始数组的新切片值。这使得切片操作与操作数组索引一样有效。因此，修改重新切片的元素（不是切片本身）会修改原始切片的元素：
	test5 := test4[:]
	test5[0] = "new"
	fmt.Println("test4:", test4)
	fmt.Println("test5:", test5)

	// 四、增长切片
	// 切片不能超出其容量增长。尝试这样做会导致运行时恐慌，就像在切片或数组的边界之外进行索引时一样。同样，切片不能在零以下重新切片以访问数组中较早的元素
	// 要增加切片的容量，必须创建一个新的更大的切片并将原始切片的内容复制到其中，这种技术是其他语言的动态数组实现在幕后工作的方式
	// 1.追加
	test6 := []int{1, 2, 3, 4}
	test6 = append(test6, 7, 8, 9)
	fmt.Println(test6)

	// 2.拷贝,这里我们创建一个空的和 test6 有相同长度的 slice test7，并且将 test6 复制给 test7
	test7 := make([]int, len(test6))
	copy(test7, test6)
	fmt.Println(test7)

	// 五、切片的组合
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

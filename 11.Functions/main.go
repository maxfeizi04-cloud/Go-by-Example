package main

import (
	"fmt"
	"reflect"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

// 可变参数函数可以用任意数量的尾随参数调用
func sum(nums ...int) {
	fmt.Print(nums, " ")

	// 是对应类型的切片
	t := reflect.TypeOf(nums)
	fmt.Println(t)
	total := 0

	// 用 for range 取出所有值
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)

	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

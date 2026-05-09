package main

import (
	"fmt"
	"maps"
)

func main() {
	// 定义并初始化一个空 map string 为 key int 为 value
	// 要创建一个空 map，使用内置的 make : make(map[key-type]val-type)
	m := make(map[string]int)

	// 使用典型的 name[key] = val 语法设置键值对
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map =", m)

	// 使用 name[key] 可以获取某个键的值
	v1 := m["k1"]
	fmt.Println("v1 =", v1)

	// 如果键不存在，将返回该值类型的零值
	v3 := m["k3"]
	fmt.Println("v3 =", v3)

	// 内置函数 len 在 map 上调用时返回键值对的数量
	fmt.Println("len =", len(m))

	// 内置的 delete 可以从映射中删除键值对
	delete(m, "k2")
	fmt.Println("map =", m)

	// 要从映射中删除所有键值对，请使用 clear 内置函数
	clear(m)
	fmt.Println("map =", m)

	_, prs := m["k2"]
	fmt.Println("prs =", prs)

	n := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("n =", n)

	n2 := map[string]int{"a": 1, "b": 2, "c": 3}

	// maps.Equal 是 Go 1.21 在标准库 maps 包中提供的一个函数，用来比较两个 map 是否“键值完全相同”
	if maps.Equal(n, n2) {
		fmt.Println("n = n2")
	}
}

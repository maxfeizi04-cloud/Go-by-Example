package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c, "len:", len(c), "cap:", cap(c))

	sl1 := s[2:5]
	fmt.Println("sl1:", sl1)

	l := s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}
	twoD := make([][]int, 4) // 1. 创建外层切片，长度为3
	for i := range twoD {    // 2. 遍历外层每个元素
		innerLen := i + 1               // 3. 内层长度 = i+1（第0行长度1，第1行长度2，第2行长度3）
		twoD[i] = make([]int, innerLen) // 4. 为第i行创建内层切片
		for j := range innerLen {       // 5. 遍历内层每个位置
			twoD[i][j] = i + j // 6. 赋值：行索引 + 列索引
		}
	}
	fmt.Println("2d:", twoD)
	fmt.Println("2d[0]:", twoD[0])
	fmt.Println("2d[1]:", twoD[1])
	fmt.Println("2d[2]:", twoD[2])

	fmt.Println("2d[0][0]", twoD[0][0])
	fmt.Println("2d[1][1]", twoD[1][1])
	fmt.Println("2d[2][0]", twoD[2][2])
}

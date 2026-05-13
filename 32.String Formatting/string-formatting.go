package main

import (
	"fmt"
	"os"
)

// point 结构体，包含 x 和 y 两个整型字段
type point struct {
	x, y int
}

func main() {

	p := point{1, 2}

	// ========== 结构体格式化 ==========

	// %v  → 只输出值（默认格式）
	// 输出: struct1: {1 2}
	fmt.Printf("struct1: %v\n", p)

	// %+v → 输出字段名和值（调试结构体时最常用）
	// 输出: struct2: {x:1 y:2}
	fmt.Printf("struct2: %+v\n", p)

	// %#v → 输出 Go 语法表示形式（可直接复制为代码）
	// 输出: struct3: main.point{x:1, y:2}
	fmt.Printf("struct3: %#v\n", p)

	// ========== 类型 ==========

	// %T → 输出变量的类型
	// 输出: type: main.point
	fmt.Printf("type: %T\n", p)

	// ========== 布尔值 ==========

	// %t → 输出布尔值（true/false）
	// 输出: bool: true
	fmt.Printf("bool: %t\n", true)

	// ========== 整数 ==========

	// %d → 十进制整数
	// 输出: int: 123
	fmt.Printf("int: %d\n", 123)

	// %b → 二进制表示
	// 14 的二进制: 8+4+2 = 1110
	// 输出: bin: 1110
	fmt.Printf("bin: %b\n", 14)

	// %c → 输出该整数对应的 Unicode 字符
	// ASCII 33 = '!'
	// 输出: char: !
	fmt.Printf("char: %c\n", 33)

	// %x → 十六进制表示（小写）
	// 456 的十六进制: 0x1C8
	// 输出: hex: 1c8
	fmt.Printf("hex: %x\n", 456)

	// ========== 浮点数 ==========

	// %f → 默认精度（6位小数）的浮点数
	// 输出: float1: 78.900000
	fmt.Printf("float1: %f\n", 78.9)

	// %e → 科学计数法（小写 e）
	// 输出: float2: 1.234000e+08
	fmt.Printf("float2: %e\n", 123400000.0)

	// %E → 科学计数法（大写 E）
	// 输出: float3: 1.234000E+08
	fmt.Printf("float3: %E\n", 123400000.0)

	// ========== 字符串 ==========

	// %s → 原样输出字符串
	// 输出: str1: "string"
	fmt.Printf("str1: %s\n", "\"string\"")

	// %q → 带引号的 Go 语法字符串（转义特殊字符）
	// 输出: str2: "\"string\""
	fmt.Printf("str2: %q\n", "\"string\"")

	// %x → 字符串转为十六进制（每个字符对应2位hex）
	// 'h'=68='68' 'e'=65='65' 'x'=78='78' ...
	// 输出: str3: 6865782074686973
	fmt.Printf("str3: %x\n", "hex this")

	// ========== 指针 ==========

	// %p → 输出指针地址（十六进制）
	// 输出: pointer: 0xc0000b4000（地址值因运行环境而异）
	fmt.Printf("pointer: %p\n", &p)

	// ========== 宽度与对齐控制 ==========

	// %6d → 最小宽度6，右对齐（默认）
	// 输出: width1: |    12|   345|
	//          ↑4个空格+12   ↑3个空格+345
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// %6.2f → 最小宽度6，保留2位小数，右对齐
	// 输出: width2: |  1.20|  3.45|
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// %-6.2f → 最小宽度6，保留2位小数，左对齐（负号=左对齐）
	// 输出: width3: |1.20  |3.45  |
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// %6s → 最小宽度6，右对齐
	// 输出: width4: |   foo|     b|
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// %-6s → 最小宽度6，左对齐
	// 输出: width5: |foo   |b     |
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// ========== Sprintf / Fprintf ==========

	// Sprintf → 格式化后返回字符串，不打印
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)
	// 输出: sprintf: a string

	// Fprintf → 格式化后输出到指定 io.Writer
	// os.Stderr 是标准错误输出
	// 输出: io: an error
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}

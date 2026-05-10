package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// s 是一个 string ，其字面值表示泰语中的“hello”一词。Go 字符串字面量是 UTF-8 编码的文本。
	const s = "สวัสดี"

	// 由于字符串等同于 []byte ，这将产生存储在其中的原始字节的长度
	fmt.Println("len(s):", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()

	// 要计算字符串中有多少个字符，我们可以使用 utf8 包。请注意，
	// RuneCountInString 的运行时间取决于字符串的大小，因为它必须按顺序解码每个 UTF-8 字符。
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idex, runeVal := range s {
		fmt.Printf("%#U starts at %d\n", runeVal, idex)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeVal, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeVal, i)
		w = width

		examineRune(runeVal)
	}
}

func examineRune(r rune) {
	// 单引号括起来的值是 rune 字面量
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	/*	┌─────────────┬────────────────────────┬──────────────────┐
		│  函数        │  作用                   │  返回值           │
		├─────────────┼────────────────────────┼──────────────────┤
		│ Contains     │ 是否包含子串            │ bool             │
		│ Count        │ 统计子串出现次数         │ int              │
		│ HasPrefix    │ 是否有指定前缀           │ bool             │
		│ HasSuffix    │ 是否有指定后缀           │ bool             │
		│ Index        │ 子串首次出现的位置        │ int（-1=没找到）  │
		│ Join         │ 用分隔符拼接切片         │ string           │
		│ Repeat       │ 重复字符串 N 次         │ string           │
		│ Replace      │ 替换子串               │ string           │
		│ Split        │ 用分隔符分割字符串       │ []string         │
		│ ToLower      │ 全部转小写              │ string           │
		│ ToUpper      │ 全部转大写              │ string           │
		└─────────────┴────────────────────────┴──────────────────┘
	*/

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}

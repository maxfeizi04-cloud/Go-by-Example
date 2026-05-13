package main

import (
	"os"            // 用于标准输出
	"text/template" // Go 标准库的模板引擎
)

func main() {
	// ============================================================
	// 1. 基础模板：使用 {{.}} 代表当前对象（"点"就是传入的数据本身）
	// ============================================================
	t1 := template.New("t1")                // 创建一个名为 "t1" 的空模板
	t1, err := t1.Parse("Value is {{.}}\n") // 解析模板字符串，{{.}} 会被替换为传入的值
	if err != nil {
		panic(err) // 解析失败则 panic
	}

	// template.Must：如果 Parse 返回错误则 panic，否则返回模板指针
	// 这里等价于上面的 if err != nil 检查（重复解析同一个模板，实际场景选一种即可）
	t1 = template.Must(t1.Parse("Value is {{.}}\n"))

	// Execute：将数据 "some text" 填入模板，输出到 stdout
	// 输出: Value is some text
	t1.Execute(os.Stdout, "some text")

	// 传入整数 5，模板会自动调用 fmt.Sprint 将其转为字符串
	// 输出: Value is 5
	t1.Execute(os.Stdout, 5)

	// 传入切片时，{{.}} 会用 fmt.Sprint 输出整个切片
	// 输出: Value is [GO Rust C++ C#]
	t1.Execute(os.Stdout, []string{
		"GO",
		"Rust",
		"C++",
		"C#",
	})

	// ============================================================
	// 2. 辅助函数：简化模板的 创建 + 解析
	// ============================================================
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// ============================================================
	// 3. 访问结构体字段：{{.Name}} 取结构体的 Name 字段
	// ============================================================
	t2 := Create("t2", "Name: {{.Name}}\n")

	// 传入匿名结构体实例，模板通过 {{.Name}} 访问字段
	// 输出: Name: Jane Doe
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// 传入 map，模板同样可以通过 {{.Name}} 取 key 对应的 value
	// 输出: Name: Mickey Mouse
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// ============================================================
	// 4. 条件判断：{{if .}} ... {{else}} ... {{end}}
	//    - 非零值/非空字符串 → 执行 if 分支
	//    - 零值/空字符串     → 执行 else 分支
	//    - "-" 表示去掉模板中多余的空白（trim 动作）
	// ============================================================
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")

	// "not empty" 为非空字符串，走 if 分支
	// 输出: yes
	t3.Execute(os.Stdout, "not empty")

	// "" 为空字符串，为"假值"，走 else 分支
	// 输出: no
	t3.Execute(os.Stdout, "")

	// ============================================================
	// 5. 循环遍历：{{range .}} ... {{end}}
	//    - 遍历切片/数组/map 的每个元素
	//    - 在 range 块内部，{{.}} 代表当前遍历到的元素
	// ============================================================
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")

	// 遍历字符串切片，逐个输出元素
	// 输出: Range: Go Rust C++ C#
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}

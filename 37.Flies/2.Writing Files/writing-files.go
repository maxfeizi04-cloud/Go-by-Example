package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// check 是一个错误处理辅助函数，遇到错误直接 panic
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// ============================================================
	// 方式一：一次性写入整个文件（最简单）
	// ============================================================

	d1 := []byte("hello\ngo\n")                  // 准备要写入的字节切片
	path1 := filepath.Join(os.TempDir(), "dat1") // 拼接临时文件路径，如 /tmp/dat1
	err := os.WriteFile(path1, d1, 0644)         // 一次性写入文件，权限 0644（rw-r--r--）
	check(err)                                   // 如果文件已存在，会覆盖原内容

	// ============================================================
	// 方式二：通过文件句柄逐次写入（更灵活）
	// ============================================================

	path2 := filepath.Join(os.TempDir(), "dat2") // 第二个临时文件路径
	f, err := os.Create(path2)                   // 创建文件（已存在则清空），返回文件句柄
	check(err)

	defer f.Close() // 确保函数结束时关闭文件，释放资源

	// ---------- 写入字节切片 ----------

	d2 := []byte{115, 111, 109, 101, 10} // ASCII: s, o, m, e, \n → "some\n"
	n2, err := f.Write(d2)               // 将字节切片写入文件，返回写入的字节数
	check(err)
	fmt.Printf("wrote %d bytes\n", n2) // 输出: wrote 5 bytes

	// ---------- 写入字符串 ----------

	n3, err := f.WriteString("writes\n") // 直接写入字符串，无需手动转 []byte
	check(err)
	fmt.Printf("wrote %d bytes\n", n3) // 输出: wrote 7 bytes

	// ---------- 强制刷盘 ----------

	f.Sync() // 将内核缓冲区的数据强制写入磁盘
	// Write 和 WriteString 是直接写文件，但操作系统可能缓存
	// Sync 确保数据落盘，适合对数据安全要求高的场景（如日志、事务）

	// ---------- 带缓冲的写入 ----------

	w := bufio.NewWriter(f) // 在文件句柄外包装一层缓冲写入器
	// 默认缓冲区大小 4096 字节，数据先写入内存缓冲区
	// 累积到一定量后再一次性写入文件，减少系统调用次数，提升性能

	n4, err := w.WriteString("buffered\n") // 写入缓冲区（此时未必写到文件）
	check(err)
	fmt.Printf("wrote %d bytes\n", n4) // 输出: wrote 9 bytes

	w.Flush() // 将缓冲区中剩余的数据全部写入文件
	// 必须调用！否则缓冲区中残留的数据会丢失
}

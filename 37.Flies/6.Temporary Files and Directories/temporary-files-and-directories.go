package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// check 是一个错误处理辅助函数，如果传入的错误不为 nil，则触发 panic
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 创建一个临时文件
	// 第一个参数 "" 表示使用系统默认的临时目录（如 Linux 的 /tmp）
	// 第二个参数 "sample" 是文件名的前缀，系统会自动在后面追加随机字符串
	// 例如生成的文件名可能是: /tmp/sample123456789
	// 返回的 f 是一个 *os.File，可直接用于读写操作
	f, err := os.CreateTemp("", "sample")
	check(err)

	// 打印系统自动分配的临时文件完整路径
	fmt.Println("Temp file name:", f.Name())

	// defer 确保 main 函数结束时删除该临时文件，避免磁盘残留
	defer os.Remove(f.Name())

	// 向临时文件写入 4 个字节的数据（1, 2, 3, 4）
	// []byte{1, 2, 3, 4} 是原始字节，不是文本字符串
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// 创建一个临时目录
	// 第一个参数 "" 同样使用系统默认临时目录
	// 第二个参数 "sampledir" 是目录名的前缀
	// 生成的目录名可能是: /tmp/sampledir123456789
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	// defer 确保 main 函数结束时递归删除该临时目录及其所有内容
	defer os.RemoveAll(dname)

	// filepath.Join 将目录路径和文件名拼接成完整路径
	// 例如: /tmp/sampledir123456789/file1
	fname := filepath.Join(dname, "file1")

	// 将 2 个字节的数据（1, 2）写入临时目录中的 file1 文件
	// 0666 权限表示所有用户均可读写（实际权限还会受 umask 影响）
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}

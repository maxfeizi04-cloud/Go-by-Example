package main

import (
	"fmt"
	"io/fs"
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

	// 创建一个名为 "subdir" 的目录，权限为 0755（rwxr-xr-x）
	err := os.Mkdir("subdir", 0755)
	check(err)

	// defer 确保在 main 函数结束时删除 "subdir" 及其所有内容，用于清理资源
	defer os.RemoveAll("subdir")

	// 定义一个闭包函数，用于创建指定名称的空文件（内容为空字节切片）
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	// 在 subdir 目录下创建空文件 file1
	createEmptyFile("subdir/file1")

	// 递归创建多级目录 subdir/parent/child（类似 mkdir -p）
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	// 在对应目录下创建多个空文件
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// 读取 subdir/parent 目录下的所有条目（文件和子目录）
	c, err := os.ReadDir("subdir/parent")
	check(err)

	// 遍历并打印每个条目的名称和是否为目录
	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 切换当前工作目录到 subdir/parent/child
	err = os.Chdir("subdir/parent/child")
	check(err)

	// 读取当前目录（即 subdir/parent/child）下的所有条目
	// 这里用 "." 表示当前目录
	c, err = os.ReadDir(".")
	check(err)

	// 遍历并打印 child 目录下的条目信息
	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 通过相对路径向上回退三级目录，回到最初的项目根目录
	// ../../.. 从 child -> parent -> subdir -> 项目根目录
	err = os.Chdir("../../..")
	check(err)

	// 递归遍历整个 "subdir" 目录树
	// WalkDir 会对目录树中的每个文件和目录调用 visit 回调函数
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}

// visit 是 WalkDir 的回调函数，对遍历到的每个条目执行此函数
// 参数说明：
//   - path: 当前条目的路径（相对于 WalkDir 的起始目录）
//   - d:    当前条目的 fs.DirEntry 接口，可获取名称、是否为目录等信息
//   - err:  遍历过程中可能产生的错误
//
// 返回 nil 表示继续遍历，返回非 nil 错误则终止遍历
func visit(path string, d fs.DirEntry, err error) error {
	// 如果在访问该路径时发生了错误（如权限不足），直接返回错误以终止遍历
	if err != nil {
		return err
	}
	// 打印当前条目的路径和是否为目录
	fmt.Println(" ", path, d.IsDir())
	return nil
}

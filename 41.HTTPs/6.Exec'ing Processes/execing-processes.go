package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

/*
注意：syscall.Exec 是 Unix 特有的系统调用，在 Windows 上不可用。如果需要跨平台，
应使用 exec.Command。Go 的 syscall 包属于底层 API，一般应用开发推荐使用更安全的 os/exec 包。
*/
func main() {

	// ===== 第一步：查找命令的完整路径 =====

	// exec.LookPath 在系统的 PATH 环境变量中搜索指定命令
	// 例如在 Linux/macOS 上可能返回 "/bin/ls"
	// 如果命令不存在（如拼写错误或未安装），返回错误
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	// binary 的值类似: "/bin/ls" 或 "/usr/bin/ls"

	// ===== 第二步：准备参数 =====

	// args 是传递给新程序的参数切片
	// 约定：第一个元素（索引 0）必须是程序名称本身
	// 这与 C 语言的 argv[0] 约定一致
	args := []string{"ls", "-a", "-l", "-h"}

	// ===== 第三步：准备环境变量 =====

	// 获取当前进程的所有环境变量，传递给新进程
	// 如果想自定义环境变量，也可以手动构造：
	//   env := []string{"PATH=/usr/bin", "LANG=en_US.UTF-8"}
	env := os.Environ()

	// ===== 第四步：执行 exec =====

	// syscall.Exec 是一个特殊的系统调用（Unix/Linux/macOS）
	// 它会用新程序 **替换** 当前进程的内存映像
	//
	// 关键特性：
	//   - 当前进程的 PID 保持不变
	//   - 当前进程的代码、数据、堆栈全部被新程序替换
	//   - 成功后，exec 后面的代码 **永远不会执行**
	//   - 只有失败时才会返回 error
	//
	// 这与 exec.Command 不同：
	//   exec.Command → 创建子进程执行，父进程继续运行
	//   syscall.Exec  → 替换自身，没有父子关系
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		// 只有 exec 失败时才会走到这里
		panic(execErr)
	}

	// ⚠️ 这行代码永远不会被执行
	// 如果 syscall.Exec 成功，当前进程已经被 ls 替换了
	fmt.Println("这行永远不会打印")
}

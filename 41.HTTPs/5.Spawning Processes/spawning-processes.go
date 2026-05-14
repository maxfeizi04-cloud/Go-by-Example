package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// ===== 示例 1：执行简单命令，获取输出 =====

	// exec.Command 创建一个要执行的命令对象
	// 参数：命令名称 + 命令参数（这里只有命令名 "date"，无额外参数）
	// 注意：这里不会立即执行命令，只是创建命令描述
	dateCmd := exec.Command("date")

	// Output() 执行命令并等待完成，返回命令的标准输出（stdout）
	// 如果命令执行失败（如命令不存在、返回非零退出码），err 不为 nil
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	// dateOut 是 []byte 类型，转为 string 后打印
	fmt.Println(string(dateOut))

	// ===== 示例 2：处理命令执行错误 =====

	// 故意传入无效参数 "-x"，date 命令会执行失败
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		// errors.AsType 是 Go 的类型断言错误检查方式
		// 检查错误链中是否包含 *exec.Error 类型的错误
		// exec.Error 表示"命令执行本身出错"，如命令找不到
		if e, ok := errors.AsType[*exec.Error](err); ok {
			fmt.Println("failed executing:", e)

			// 检查错误链中是否包含 *exec.ExitError 类型的错误
			// exec.ExitError 表示"命令执行了，但返回了非零退出码"
		} else if e, ok := errors.AsType[*exec.ExitError](err); ok {
			// 获取命令的退出码（exit code）
			// 常见退出码：0=成功，1=一般错误，2=用法错误，127=命令未找到
			exitCode := e.ExitCode()
			fmt.Println("command exit rc =", exitCode)
		} else {
			panic(err)
		}
	}

	// ===== 示例 3：通过管道与命令交互 =====

	// 创建 grep 命令，用于搜索包含 "hello" 的行
	grepCmd := exec.Command("grep", "hello")

	// 获取命令的标准输入管道（*io.WriteCloser）
	// 通过这个管道可以向命令的 stdin 写入数据
	grepIn, _ := grepCmd.StdinPipe()

	// 获取命令的标准输出管道（*io.ReadCloser）
	// 通过这个管道可以从命令的 stdout 读取数据
	grepOut, _ := grepCmd.StdoutPipe()

	// Start() 启动命令，但不等待完成（非阻塞）
	// 命令开始运行后，我们可以同时向 stdin 写入、从 stdout 读取
	grepCmd.Start()

	// 向 grep 的 stdin 写入两行文本
	// grep 会筛选包含 "hello" 的行
	grepIn.Write([]byte("hello grep\ngoodbye grep"))

	// 关闭 stdin 管道，通知 grep 没有更多输入了
	// 不关闭的话 grep 会一直等待输入，不会输出结果
	grepIn.Close()

	// 从 grep 的 stdout 读取所有输出
	// "hello grep" 包含 "hello" → 会被保留
	// "goodbye grep" 不包含 "hello" → 会被过滤掉
	grepBytes, _ := io.ReadAll(grepOut)

	// Wait() 等待命令执行结束，回收资源
	// 必须在读取完 stdout 之后调用，否则可能死锁
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// ===== 示例 4：执行 shell 命令 =====

	// 有时需要执行包含管道、重定向等 shell 特性的命令
	// 不能直接 exec.Command("ls -a -l -h")（会被当作一个整体命令名）
	// 正确方式：用 bash -c "命令字符串"，让 bash 来解释执行
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")

	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

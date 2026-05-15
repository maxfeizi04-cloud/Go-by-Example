# Go by Example

通过实战示例学习 Go 语言，涵盖从基础语法到高级特性的完整学习路径。

## 项目简介

本项目是 [Go by Example](https://gobyexample.com/) 的学习实践代码库，以 `Go 1.26` 编写。每个目录对应一个 Go 语言特性，包含可运行的示例代码，适合 Go 初学者和进阶学习者参考。

## 目录结构

| 序号 | 主题 | 说明 |
|------|------|------|
| 01 | [Hello World](./01.Hello-World/hello-world.go) | 第一个 Go 程序 |
| 02 | [Values](./02.Values/values.go) | 基本数据类型 |
| 03 | [Variables](./03.Variables/variables.go) | 变量声明与赋值 |
| 04 | [Constants](./04.Constants/constants.go) | 常量定义 |
| 05 | [For](./05.For/for.go) | 循环语句 |
| 06 | [If/Else](./06.If-Else/if-else.go) | 条件分支 |
| 07 | [Switch](./07.Switch/switch.go) | 分支选择 |
| 08 | [Arrays](./08.Arrays/main.go) | 数组 |
| 09 | [Slices](./09.Slices/main.go) | 切片 |
| 10 | [Maps](./10.Maps/main.go) | 映射 |
| 11 | [Functions](./11.Functions/main.go) | 函数 |
| 12 | [Closures](./12.Closures/main.go) | 闭包 |
| 13 | [Recursion](./13.Recursion/main.go) | 递归 |
| 14 | [Range](./14.Range%20over%20Built-in%20Types/main.go) | 遍历内置类型 |
| 15 | [Pointers](./15.Pointers/main.go) | 指针 |
| 16 | [Strings and Runes](./16.Strings%20and%20Runes/main.go) | 字符串与字符 |
| 17 | [Structs](./17.Structs/main.go) | 结构体 |
| 18 | [Methods](./18.Methods/main.go) | 方法 |
| 19 | [Interfaces](./19.Interfaces/main.go) | 接口 |
| 20 | [Enums](./20.Enums/main.go) | 枚举 |
| 21 | [Struct Embedding](./21.Struct%20Embedding/main.go) | 结构体嵌入 |
| 22 | [Generics](./22.Generics/generics.go) | 泛型 |
| 23 | [Range over Iterators](./23.Range%20over%20Iterators/range-over-iterators.go) | 迭代器遍历 |
| 24 | [Errors](./24.Errors/errors.go) | 错误处理 |
| 25 | [Custom Errors](./25.Custom%20Errors/custom-errors.go) | 自定义错误 |
| 26 | [Goroutines & Channel](./26.Goroutines%20Channel/) | 协程与通道 |
| 27 | [Sorts](./27.Sorts/) | 排序 |
| 28 | [Panic](./28.Panic/panic.go) | panic 异常 |
| 29 | [Defer](./29.Defer/defer.go) | 延迟执行 |
| 30 | [Recover](./30.Recover/recover.go) | 异常恢复 |
| 31 | [String Functions](./31.String%20Functions/main.go) | 字符串函数 |
| 32 | [String Formatting](./32.String%20Formatting/string-formatting.go) | 字符串格式化 |
| 33 | [Templates](./33.Templates/) | 模板引擎 |
| 34 | [URL Parsing](./34.URL%20Parsing/url-parsing.go) | URL 解析 |
| 35 | [SHA256 Hashes](./35.SHA256%20Hashes/sha256-hashes.go) | SHA256 哈希 |
| 36 | [Base64 Encoding](./36.Base64%20Encoding/base64-encoding.go) | Base64 编码 |
| 37 | [Files](./37.Flies/) | 文件操作 |
| 38 | [Commands](./38.Commands/) | 命令行参数 |
| 39 | [Environment Variables](./39.Environment%20Variables/environment-variables.go) | 环境变量 |
| 40 | [Logging](./40.Logging/logging.go) | 日志 |
| 41 | [HTTP](./41.HTTPs/) | HTTP 客户端与服务端 |

## 运行方式

```bash
# 运行单个示例
go run "01.Hello-World/hello-world.go"

# 进入目录运行
cd "01.Hello-World" && go run hello-world.go
```

## 学习路线

1. **基础语法** (01-07)：程序结构、数据类型、控制流
2. **复合类型** (08-10)：数组、切片、映射
3. **函数式编程** (11-14)：函数、闭包、递归、遍历
4. **面向对象** (15-21)：指针、结构体、方法、接口、泛型
5. **错误处理** (24-30)：错误、异常、恢复
6. **标准库** (31-41)：字符串处理、模板、网络、文件、日志等

## 技术栈

- **语言**：Go 1.26
- **模块名**：`Go-by-Example`

---

持续更新中，每天进步一点点 🚀

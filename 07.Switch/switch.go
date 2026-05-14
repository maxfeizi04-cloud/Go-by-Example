package main

import (
	"fmt"
	"time"
)

func main() {

	// ===== 示例 1：基本 switch（按值匹配） =====

	// switch 最基础的用法：将变量 i 与每个 case 的值逐一比较
	// 匹配到哪个 case 就执行对应的代码块
	// Go 的 switch 不需要显式 break，匹配成功后自动跳出（不会穿透到下一个 case）
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two") // i == 2，匹配成功，执行这里
	case 3:
		fmt.Println("three")
	}

	// ===== 示例 2：多值 case =====

	// case 后面可以跟多个值，用逗号分隔
	// 只要匹配其中任意一个值，就执行该 case
	// time.Now().Weekday() 返回当前是星期几
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		// 星期六或星期天都会匹配
		fmt.Println("It's the weekend")
	default:
		// 其他星期几（周一到周五）
		fmt.Println("It's a weekday")
	}

	// ===== 示例 3：无条件 switch（替代 if-else 链） =====

	// switch 后面不跟表达式时，每个 case 可以写任意布尔条件
	// 等价于 if-else if-else，但更清晰易读
	// 从上到下依次判断，第一个为 true 的 case 被执行
	t := time.Now()
	switch {
	case t.Hour() < 12:
		// 当前小时 < 12（中午 12 点之前）
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// ===== 示例 4：类型 switch（type switch） =====

	// 定义一个匿名函数，参数类型为 interface{}（空接口，可接收任意类型）
	whatAmI := func(i interface{}) {
		// i.(type) 是类型断言的 switch 形式
		// 只能在 switch 中使用，用于判断接口变量的实际类型
		// t 被赋值为断言后的具体类型值
		switch t := i.(type) {
		case bool:
			// i 的实际类型是 bool
			fmt.Println("I'm a bool")
		case int:
			// i 的实际类型是 int
			fmt.Println("I'm an int")
		default:
			// 其他未知类型
			// %T 格式化动词打印变量的类型名
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)  // bool   → "I'm a bool"
	whatAmI(1)     // int    → "I'm an int"
	whatAmI("hey") // string → "Don't know type string"
}

package main

import (
	"fmt"
	"testing"
)

// TestIntMinBasic 是一个基本的单元测试
// 以 Test 开头的函数会被 go test 自动识别为测试用例
// 参数 *testing.T 提供了测试失败报告、日志等功能
func TestIntMinBasic(t *testing.T) {
	// 调用被测函数，获取结果
	ans := IntMin(2, -2)

	// 断言：如果结果不符合预期，通过 t.Errorf 报告错误
	// %d 是整数占位符，Errorf 格式化输出错误信息但不会终止测试
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// TestIntMinTableDriven 是表驱动测试（Table-Driven Test）
// 这是 Go 中最常用的测试模式：用结构体切片定义多组输入和期望输出
func TestIntMinTableDriven(t *testing.T) {
	// 定义测试用例表
	// 每个结构体包含一组输入 (a, b) 和期望输出 (want)
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},   // 0 和 1 的最小值为 0
		{1, 0, 0},   // 1 和 0 的最小值为 0
		{2, -2, -2}, // 2 和 -2 的最小值为 -2
		{0, -1, -1}, // 0 和 -1 的最小值为 -1
		{-1, 0, -1}, // -1 和 0 的最小值为 -1
	}

	// 遍历每组测试用例
	for _, tt := range tests {
		// 用输入值构造子测试名称，例如 "0,1"、"1,0"
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)

		// t.Run 创建一个子测试（Subtest）
		// 好处：
		//   1. 每组用例独立运行，某组失败不影响其他组
		//   2. 可通过名称单独运行，例如 go test -run TestIntMinTableDriven/2,-2
		//   3. 输出结果会按子测试分组显示，便于定位问题
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				// 报告错误：实际值与期望值不匹配
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// BenchmarkIntMin 是性能基准测试
// 以 Benchmark 开头的函数会被 go test -bench 自动识别
// 参数 *testing.B 提供基准测试所需的控制功能
func BenchmarkIntMin(b *testing.B) {
	// b.Loop() 是 Go 1.24 引入的基准测试循环方式
	// 框架会自动决定迭代次数以获得稳定的性能数据
	// 与旧版的 for i := 0; i < b.N; i++ 功能等价
	for b.Loop() {
		IntMin(1, 2)
	}
}

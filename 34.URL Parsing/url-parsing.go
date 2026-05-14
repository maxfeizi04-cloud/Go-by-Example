package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// 原始 URL 字符串，包含 URL 的所有组成部分
	//
	//  postgres://user:pass@host.com:5432/path?k=v#f
	//  ├──────┤ ├──────┤ ├──────────┤ ├───┤ ├──┤ ├┤
	//   Scheme   User     Host:Port   Path  Query Fragment
	//
	// URL 通用结构：
	//   scheme://[user:pass@]host[:port]/path[?query][#fragment]
	s := "postgres://user:pass@host.com:5432/path?k=v1#f"

	// ============================================================
	// 1. 解析 URL 字符串 → url.URL 结构体
	//    将完整 URL 拆解为各个组成部分
	// ============================================================
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// ============================================================
	// 2. Scheme：协议方案
	// ============================================================
	fmt.Println(u.Scheme)
	// postgres
	//    ↑ 也可以是 http、https、ftp、ssh 等

	// ============================================================
	// 3. User：用户认证信息
	//    u.User 本身是 *url.Userinfo 类型
	// ============================================================
	fmt.Println(u.User)
	// user:pass          完整的 用户名:密码
	fmt.Println(u.User.Username())
	// user               只取用户名
	p, _ := u.User.Password()
	fmt.Println(p)
	// pass               只取密码

	// ============================================================
	// 4. Host：主机名 + 端口
	//    u.Host 是 "host:port" 的合并字符串
	//    用 net.SplitHostPort 拆分
	// ============================================================
	fmt.Println(u.Host)
	// host.com:5432      完整的主机:端口
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	// host.com           主机名
	fmt.Println(port)
	// 5432               端口号

	// ============================================================
	// 5. Path：资源路径
	// ============================================================
	fmt.Println(u.Path)
	// /path

	// ============================================================
	// 6. Fragment：片段标识符（# 后面的部分）
	//    通常用于页面内锚点，不会发送到服务器
	// ============================================================
	fmt.Println(u.Fragment)
	// f

	// ============================================================
	// 7. RawQuery → 原始查询字符串（未解码）
	//    ParseQuery → 将查询字符串解析为 map
	// ============================================================
	fmt.Println(u.RawQuery)
	// k=v                原始查询参数
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	// map[k:[v]]         解析后的 map，value 是 []string（支持同名参数）
	fmt.Println(m["k"][0])
	// v                  取 key="k" 的第一个值
}

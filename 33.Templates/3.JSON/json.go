package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// ============================================================
// 结构体定义：用于演示 JSON 序列化/反序列化
// ============================================================

// response1：没有 JSON tag，字段名原样作为 JSON key
// 序列化结果：{"Page":1,"Fruits":["apple","peach","pear"]}
// 注意：key 首字母大写
type response1 struct {
	Page   int
	Fruits []string
}

// response2：有 JSON tag，可以自定义 JSON key 的名称
// 序列化结果：{"page":1,"fruits":["apple","peach","pear"]}
// 注意：key 全小写，由 tag 控制
type response2 struct {
	Page   int      `json:"page"`   // JSON key 为 "page"
	Fruits []string `json:"fruits"` // JSON key 为 "fruits"
}

func main() {

	// ============================================================
	// 1. 基本类型 → JSON
	//    json.Marshal 接收任意值，返回 ([]byte, error)
	// ============================================================

	// bool → "true"
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB)) // true

	// int → "1"
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB)) // 1

	// float → "2.34"
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB)) // 2.34

	// string → "\"gopher\""（JSON 中字符串必须带双引号）
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB)) // "gopher"

	// ============================================================
	// 2. 切片 → JSON 数组
	// ============================================================
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB)) // ["apple","peach","pear"]

	// ============================================================
	// 3. Map → JSON 对象
	//    map 的 key 必须是 string（JSON 对象的 key 只能是字符串）
	// ============================================================
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB)) // {"apple":5,"lettuce":7}

	// ============================================================
	// 4. 结构体 → JSON（无 tag）
	//    字段名首字母大写 → JSON key 也是大写
	//    未导出的字段（小写开头）不会被序列化
	// ============================================================
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	// {"Page":1,"Fruits":["apple","peach","pear"]}

	// ============================================================
	// 5. 结构体 → JSON（有 tag）
	//    json tag 指定了 JSON key 的名称
	//    生产环境推荐始终使用 tag，保证 JSON 格式稳定
	// ============================================================
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
	// {"page":1,"fruits":["apple","peach","pear"]}

	// ============================================================
	// 6. JSON → 通用 map（不知道结构时使用）
	//    map[string]interface{} 可以接收任意 JSON 对象
	// ============================================================
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// 声明一个 map，value 类型为 interface{}（可以存任意类型）
	var dat map[string]interface{}

	// json.Unmarshal：将 JSON 字节切片解析到目标变量
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat) // map[num:6.13 strs:[a b]]

	// ============================================================
	// 7. 从 map 中取值需要类型断言
	//    JSON 中的数字统一解析为 float64
	// ============================================================
	num := dat["num"].(float64) // 类型断言：interface{} → float64
	fmt.Println(num)            // 6.13

	// JSON 数组解析为 []interface{}
	strs := dat["strs"].([]interface{}) // 类型断言：interface{} → []interface{}
	str1 := strs[1].(string)            // 再断言：interface{} → string
	fmt.Println(str1)                   // a

	// ============================================================
	// 8. JSON → 具体结构体（知道结构时推荐使用）
	//    自动根据 json tag 匹配字段
	// ============================================================
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res) // JSON → 结构体
	fmt.Println(res)                  // {1 [apple peach]}
	fmt.Println(res.Fruits[0])        // apple（直接访问字段，无需断言）

	// ============================================================
	// 9. Encoder：流式编码，直接写入 io.Writer
	//    适合向 HTTP ResponseWriter、文件等直接输出 JSON
	//    自动在末尾添加换行符 \n
	// ============================================================
	enc := json.NewEncoder(os.Stdout) // 创建编码器，输出到标准输出
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d) // 直接输出：{"apple":5,"lettuce":7}\n

	// ============================================================
	// 10. Decoder：流式解码，从 io.Reader 读取 JSON
	//     适合解析 HTTP 请求体、文件等流式数据
	// ============================================================
	dec := json.NewDecoder(strings.NewReader(str)) // 从字符串创建解码器
	res1 := response2{}
	dec.Decode(&res1) // 读取并解析 JSON → 结构体
	fmt.Println(res1) // {1 [apple peach]}
}

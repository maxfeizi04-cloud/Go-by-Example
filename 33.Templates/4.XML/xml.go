package main

import (
	"encoding/xml"
	"fmt"
)

// ============================================================
// Plant 结构体：定义 XML 的映射关系
// ============================================================

type Plant struct {
	XMLName xml.Name `xml:"plant"`   // 指定 XML 根元素名为 <plant>
	Id      int      `xml:"id,attr"` // "attr" 表示这是 XML 属性，而非子元素
	Name    string   `xml:"name"`    // 映射为 <name> 子元素
	Origin  []string `xml:"origin"`  // 切片 → 多个 <origin> 子元素
}

// String() 方法：让 fmt.Println 直接输出可读格式
// 实现了 fmt.Stringer 接口，打印 Plant 时自动调用
func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {

	// ============================================================
	// 1. 创建 Plant 实例并赋值
	// ============================================================
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// ============================================================
	// 2. 结构体 → XML（序列化）
	//    MarshalIndent 生成格式化（带缩进）的 XML
	//    参数：(v any, prefix string, indent string)
	//      prefix = 每行前缀（这里用空格 " "）
	//      indent = 缩进字符（这里用两个空格 "  "）
	// ============================================================
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// 注意：
	//   - id="27" 是属性（因为 tag 中有 ",attr"）
	//   - <origin> 出现两次（因为 Origin 是 []string 切片）

	// ============================================================
	// 3. 手动添加 XML 声明头
	//    xml.Header = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
	//    MarshalIndent 默认不包含声明头，需要手动拼接
	// ============================================================
	fmt.Println(xml.Header + string(out))
	// 输出：
	// <?xml version="1.0" encoding="UTF-8"?>
	// <plant id="27">
	//   <name>Coffee</name>
	//   ...
	// </plant>

	// ============================================================
	// 4. XML → 结构体（反序列化）
	//    将 XML 字节流解析回 Plant 结构体
	// ============================================================
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
	// 输出：Plant id=27, name=Coffee, origin=[Ethiopia Brazil]
	// 这里调用了我们自定义的 String() 方法

	// ============================================================
	// 5. 嵌套 XML 结构：用路径表达式描述层级关系
	// ============================================================
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`            // 根元素 <nesting>
		Plants  []*Plant `xml:"parent>child>plant"` // 路径嵌套
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))

}

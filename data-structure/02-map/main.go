package main

import "fmt"

func main() {
	fmt.Println("=== Map（映射/字典）===\n")

	// 1. 创建 Map
	fmt.Println("1. 创建 Map:")
	m1 := make(map[string]int)
	m2 := map[string]int{"a": 1, "b": 2}
	fmt.Printf("   m1: %v\n", m1)
	fmt.Printf("   m2: %v\n", m2)

	// 2. 添加和修改
	fmt.Println("\n2. 添加和修改:")
	m1["apple"] = 5
	m1["banana"] = 3
	fmt.Printf("   添加后: %v\n", m1)
	m1["apple"] = 10
	fmt.Printf("   修改后: %v\n", m1)

	// 3. 读取（检查键是否存在）
	fmt.Println("\n3. 读取:")
	value, exists := m1["apple"]
	fmt.Printf("   apple: %d, 存在: %t\n", value, exists)
	value, exists = m1["orange"]
	fmt.Printf("   orange: %d, 存在: %t\n", value, exists)

	// 4. 删除
	fmt.Println("\n4. 删除:")
	delete(m1, "banana")
	fmt.Printf("   删除 banana 后: %v\n", m1)

	// 5. 遍历
	fmt.Println("\n5. 遍历:")
	m3 := map[string]int{"x": 1, "y": 2, "z": 3}
	for key, value := range m3 {
		fmt.Printf("   %s: %d\n", key, value)
	}

	// 6. Map 长度
	fmt.Println("\n6. Map 长度:")
	fmt.Printf("   长度: %d\n", len(m3))

	// 7. 嵌套 Map
	fmt.Println("\n7. 嵌套 Map:")
	nested := map[string]map[string]int{
		"fruits": {"apple": 5, "banana": 3},
		"veggies": {"carrot": 2, "potato": 4},
	}
	fmt.Printf("   嵌套 Map: %v\n", nested)
	fmt.Printf("   fruits.apple: %d\n", nested["fruits"]["apple"])

	// 8. Map 不是并发安全的
	fmt.Println("\n8. 注意事项:")
	fmt.Println("   - Map 不是并发安全的")
	fmt.Println("   - 遍历顺序是随机的")
	fmt.Println("   - 删除元素后不会缩容")
}

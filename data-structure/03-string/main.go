package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("=== String（字符串）===\n")

	// 1. 字符串基础
	fmt.Println("1. 字符串基础:")
	s := "Hello, 世界"
	fmt.Printf("   字符串: %s\n", s)
	fmt.Printf("   长度（字节）: %d\n", len(s))
	fmt.Printf("   长度（字符）: %d\n", utf8.RuneCountInString(s))

	// 2. 字符串是不可变的
	fmt.Println("\n2. 字符串不可变:")
	fmt.Println("   不能直接修改字符串")
	fmt.Println("   s[0] = 'h' // 编译错误")

	// 3. 字符串拼接
	fmt.Println("\n3. 字符串拼接:")
	s1 := "Hello"
	s2 := "World"
	s3 := s1 + " " + s2
	fmt.Printf("   + 操作符: %s\n", s3)

	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	fmt.Printf("   strings.Builder: %s\n", builder.String())

	// 4. 字符串切片
	fmt.Println("\n4. 字符串切片:")
	s4 := "Hello, World"
	fmt.Printf("   原字符串: %s\n", s4)
	fmt.Printf("   s[0:5]: %s\n", s4[0:5])
	fmt.Printf("   s[7:]: %s\n", s4[7:])

	// 5. 遍历字符串
	fmt.Println("\n5. 遍历字符串:")
	s5 := "Go语言"
	fmt.Println("   按字节遍历:")
	for i := 0; i < len(s5); i++ {
		fmt.Printf("     [%d] = %c (%d)\n", i, s5[i], s5[i])
	}

	fmt.Println("   按字符遍历:")
	for i, r := range s5 {
		fmt.Printf("     [%d] = %c (%U)\n", i, r, r)
	}

	// 6. 字符串转换
	fmt.Println("\n6. 字符串转换:")
	bytes := []byte("Hello")
	str := string(bytes)
	fmt.Printf("   []byte -> string: %s\n", str)
	fmt.Printf("   string -> []byte: %v\n", []byte(str))

	// 7. 常用字符串操作
	fmt.Println("\n7. 常用操作:")
	s6 := "  Hello, World  "
	fmt.Printf("   原字符串: '%s'\n", s6)
	fmt.Printf("   Trim: '%s'\n", strings.TrimSpace(s6))
	fmt.Printf("   ToUpper: '%s'\n", strings.ToUpper(s6))
	fmt.Printf("   ToLower: '%s'\n", strings.ToLower(s6))
	fmt.Printf("   Contains: %t\n", strings.Contains(s6, "World"))
	fmt.Printf("   Split: %v\n", strings.Split("a,b,c", ","))
	fmt.Printf("   Join: %s\n", strings.Join([]string{"a", "b", "c"}, "-"))
}

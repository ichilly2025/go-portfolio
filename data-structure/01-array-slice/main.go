package main

import "fmt"

func main() {
	fmt.Println("=== Array vs Slice ===\n")

	// 1. Array（数组）- 固定长度
	fmt.Println("1. Array（数组）:")
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	fmt.Printf("   数组: %v\n", arr)
	fmt.Printf("   长度: %d\n", len(arr))
	fmt.Printf("   类型: %T\n", arr)

	// 2. Slice（切片）- 动态长度
	fmt.Println("\n2. Slice（切片）:")
	slice := []int{1, 2, 3}
	fmt.Printf("   切片: %v\n", slice)
	fmt.Printf("   长度: %d, 容量: %d\n", len(slice), cap(slice))

	// 3. Slice 扩容
	fmt.Println("\n3. Slice 扩容:")
	slice = append(slice, 4, 5, 6)
	fmt.Printf("   扩容后: %v\n", slice)
	fmt.Printf("   长度: %d, 容量: %d\n", len(slice), cap(slice))

	// 4. Slice 切片操作
	fmt.Println("\n4. Slice 切片操作:")
	sub := slice[1:4]
	fmt.Printf("   原切片: %v\n", slice)
	fmt.Printf("   子切片 [1:4]: %v\n", sub)
	fmt.Printf("   共享底层数组\n")

	// 5. make 创建 Slice
	fmt.Println("\n5. make 创建 Slice:")
	s1 := make([]int, 5)      // 长度 5，容量 5
	s2 := make([]int, 5, 10)  // 长度 5，容量 10
	fmt.Printf("   s1: len=%d, cap=%d\n", len(s1), cap(s1))
	fmt.Printf("   s2: len=%d, cap=%d\n", len(s2), cap(s2))

	// 6. copy 复制 Slice
	fmt.Println("\n6. copy 复制 Slice:")
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Printf("   源: %v\n", src)
	fmt.Printf("   目标: %v\n", dst)
	dst[0] = 999
	fmt.Printf("   修改后 - 源: %v, 目标: %v\n", src, dst)

	// 7. 二维 Slice
	fmt.Println("\n7. 二维 Slice:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("   矩阵: %v\n", matrix)
	fmt.Printf("   matrix[1][2] = %d\n", matrix[1][2])
}

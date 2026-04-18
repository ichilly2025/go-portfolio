package main

import "fmt"

func main() {
	fmt.Println("=== Function（函数）===\n")

	// 1. 基本函数
	fmt.Println("1. 基本函数:")
	add := func(a, b int) int {
		return a + b
	}
	fmt.Printf("   add(3, 5) = %d\n", add(3, 5))

	// 2. 多返回值
	fmt.Println("\n2. 多返回值:")
	divide := func(a, b int) (int, int) {
		return a / b, a % b
	}
	quotient, remainder := divide(10, 3)
	fmt.Printf("   10 / 3 = %d 余 %d\n", quotient, remainder)

	// 3. 命名返回值
	fmt.Println("\n3. 命名返回值:")
	swap := func(a, b int) (x, y int) {
		x = b
		y = a
		return // 自动返回 x, y
	}
	x, y := swap(1, 2)
	fmt.Printf("   swap(1, 2) = %d, %d\n", x, y)

	// 4. 可变参数
	fmt.Println("\n4. 可变参数:")
	sum := func(nums ...int) int {
		total := 0
		for _, n := range nums {
			total += n
		}
		return total
	}
	fmt.Printf("   sum(1, 2, 3) = %d\n", sum(1, 2, 3))
	fmt.Printf("   sum(1, 2, 3, 4, 5) = %d\n", sum(1, 2, 3, 4, 5))

	// 5. 匿名函数
	fmt.Println("\n5. 匿名函数:")
	result := func(a, b int) int {
		return a * b
	}(3, 4)
	fmt.Printf("   匿名函数立即执行: %d\n", result)

	// 6. 闭包
	fmt.Println("\n6. 闭包:")
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	fmt.Printf("   第 1 次调用: %d\n", counter())
	fmt.Printf("   第 2 次调用: %d\n", counter())
	fmt.Printf("   第 3 次调用: %d\n", counter())

	// 7. 函数作为参数
	fmt.Println("\n7. 函数作为参数:")
	apply := func(f func(int, int) int, a, b int) int {
		return f(a, b)
	}

	multiply := func(a, b int) int {
		return a * b
	}

	fmt.Printf("   apply(add, 3, 4) = %d\n", apply(add, 3, 4))
	fmt.Printf("   apply(multiply, 3, 4) = %d\n", apply(multiply, 3, 4))

	// 8. 函数作为返回值
	fmt.Println("\n8. 函数作为返回值:")
	makeAdder := func(x int) func(int) int {
		return func(y int) int {
			return x + y
		}
	}

	add5 := makeAdder(5)
	add10 := makeAdder(10)
	fmt.Printf("   add5(3) = %d\n", add5(3))
	fmt.Printf("   add10(3) = %d\n", add10(3))

	// 9. defer
	fmt.Println("\n9. defer（延迟执行）:")
	deferDemo := func() {
		defer fmt.Println("   第三个执行")
		defer fmt.Println("   第二个执行")
		fmt.Println("   第一个执行")
	}
	deferDemo()

	// 10. 递归
	fmt.Println("\n10. 递归:")
	var factorial func(int) int
	factorial = func(n int) int {
		if n <= 1 {
			return 1
		}
		return n * factorial(n-1)
	}
	fmt.Printf("   5! = %d\n", factorial(5))
}

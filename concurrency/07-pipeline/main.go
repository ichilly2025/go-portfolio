package main

import "fmt"

// 阶段1：生成数字
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// 阶段2：平方
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// 阶段3：加倍
func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

func main() {
	fmt.Println("=== Pipeline (流水线) ===")
	fmt.Println("多阶段数据处理\n")

	// 构建流水线：生成 → 平方 → 加倍
	fmt.Println("流水线: 生成数字 → 平方 → 加倍\n")

	nums := generate(1, 2, 3, 4, 5)
	squared := square(nums)
	doubled := double(squared)

	// 输出结果
	fmt.Println("结果:")
	for result := range doubled {
		fmt.Printf("%d ", result)
	}
	fmt.Println("\n\n说明: 1² × 2 = 2, 2² × 2 = 8, 3² × 2 = 18, ...")
}

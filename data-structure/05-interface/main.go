package main

import "fmt"

// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 实现接口 - 矩形
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 实现接口 - 圆形
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// 使用接口
func PrintShapeInfo(s Shape) {
	fmt.Printf("   面积: %.2f\n", s.Area())
	fmt.Printf("   周长: %.2f\n", s.Perimeter())
}

// 空接口
func PrintAnything(v interface{}) {
	fmt.Printf("   值: %v, 类型: %T\n", v, v)
}

func main() {
	fmt.Println("=== Interface（接口）===\n")

	// 1. 接口基础
	fmt.Println("1. 接口基础:")
	var s Shape
	fmt.Printf("   零值: %v\n", s)

	// 2. 实现接口
	fmt.Println("\n2. 矩形:")
	rect := Rectangle{Width: 10, Height: 5}
	PrintShapeInfo(rect)

	fmt.Println("\n3. 圆形:")
	circle := Circle{Radius: 7}
	PrintShapeInfo(circle)

	// 4. 接口变量
	fmt.Println("\n4. 接口变量:")
	s = rect
	fmt.Printf("   s 是矩形: %+v\n", s)
	s = circle
	fmt.Printf("   s 是圆形: %+v\n", s)

	// 5. 类型断言
	fmt.Println("\n5. 类型断言:")
	s = rect
	if r, ok := s.(Rectangle); ok {
		fmt.Printf("   是矩形: %+v\n", r)
	}

	if c, ok := s.(Circle); ok {
		fmt.Printf("   是圆形: %+v\n", c)
	} else {
		fmt.Println("   不是圆形")
	}

	// 6. 类型选择
	fmt.Println("\n6. 类型选择:")
	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("   整数: %d\n", v)
		case string:
			fmt.Printf("   字符串: %s\n", v)
		case Rectangle:
			fmt.Printf("   矩形: %+v\n", v)
		default:
			fmt.Printf("   未知类型: %T\n", v)
		}
	}

	checkType(42)
	checkType("hello")
	checkType(rect)
	checkType(3.14)

	// 7. 空接口
	fmt.Println("\n7. 空接口（interface{}）:")
	PrintAnything(42)
	PrintAnything("hello")
	PrintAnything(rect)
	PrintAnything([]int{1, 2, 3})

	// 8. 接口组合
	fmt.Println("\n8. 接口组合:")
	type Reader interface {
		Read() string
	}
	type Writer interface {
		Write(string)
	}
	type ReadWriter interface {
		Reader
		Writer
	}
	fmt.Println("   ReadWriter 接口组合了 Reader 和 Writer")
}

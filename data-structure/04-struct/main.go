package main

import "fmt"

// 定义结构体
type Person struct {
	Name string
	Age  int
	City string
}

// 方法
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s", p.Name)
}

// 指针接收者方法（可以修改结构体）
func (p *Person) Birthday() {
	p.Age++
}

// 嵌套结构体
type Address struct {
	Street string
	City   string
}

type Employee struct {
	Person  // 匿名字段（嵌入）
	Address
	Salary int
}

func main() {
	fmt.Println("=== Struct（结构体）===\n")

	// 1. 创建结构体
	fmt.Println("1. 创建结构体:")
	p1 := Person{Name: "Alice", Age: 25, City: "Beijing"}
	p2 := Person{"Bob", 30, "Shanghai"} // 按顺序
	var p3 Person                        // 零值
	fmt.Printf("   p1: %+v\n", p1)
	fmt.Printf("   p2: %+v\n", p2)
	fmt.Printf("   p3: %+v\n", p3)

	// 2. 访问字段
	fmt.Println("\n2. 访问字段:")
	fmt.Printf("   p1.Name: %s\n", p1.Name)
	fmt.Printf("   p1.Age: %d\n", p1.Age)

	// 3. 修改字段
	fmt.Println("\n3. 修改字段:")
	p1.Age = 26
	fmt.Printf("   修改后: %+v\n", p1)

	// 4. 指针
	fmt.Println("\n4. 结构体指针:")
	p4 := &Person{Name: "Charlie", Age: 35, City: "Guangzhou"}
	fmt.Printf("   p4: %+v\n", p4)
	fmt.Printf("   (*p4).Name: %s\n", (*p4).Name)
	fmt.Printf("   p4.Name: %s (自动解引用)\n", p4.Name)

	// 5. 方法
	fmt.Println("\n5. 方法:")
	fmt.Printf("   %s\n", p1.Greet())
	fmt.Printf("   生日前: Age = %d\n", p1.Age)
	p1.Birthday()
	fmt.Printf("   生日后: Age = %d\n", p1.Age)

	// 6. 匿名结构体
	fmt.Println("\n6. 匿名结构体:")
	point := struct {
		X, Y int
	}{10, 20}
	fmt.Printf("   point: %+v\n", point)

	// 7. 嵌套结构体
	fmt.Println("\n7. 嵌套结构体:")
	emp := Employee{
		Person:  Person{Name: "David", Age: 28, City: "Shenzhen"},
		Address: Address{Street: "Main St", City: "Shenzhen"},
		Salary:  50000,
	}
	fmt.Printf("   员工: %+v\n", emp)
	fmt.Printf("   姓名: %s\n", emp.Name)   // 直接访问嵌入字段
	fmt.Printf("   街道: %s\n", emp.Street) // 直接访问嵌入字段

	// 8. 结构体比较
	fmt.Println("\n8. 结构体比较:")
	p5 := Person{Name: "Alice", Age: 25, City: "Beijing"}
	p6 := Person{Name: "Alice", Age: 25, City: "Beijing"}
	fmt.Printf("   p5 == p6: %t\n", p5 == p6)
}

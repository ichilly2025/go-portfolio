package main

import "fmt"

func main() {
	fmt.Println("=== Pointer（指针）===\n")

	// 1. 指针基础
	fmt.Println("1. 指针基础:")
	x := 10
	p := &x // 取地址
	fmt.Printf("   x 的值: %d\n", x)
	fmt.Printf("   x 的地址: %p\n", &x)
	fmt.Printf("   p 的值（地址）: %p\n", p)
	fmt.Printf("   p 指向的值: %d\n", *p) // 解引用

	// 2. 通过指针修改值
	fmt.Println("\n2. 通过指针修改值:")
	*p = 20
	fmt.Printf("   修改后 x: %d\n", x)
	fmt.Printf("   修改后 *p: %d\n", *p)

	// 3. 指针作为函数参数
	fmt.Println("\n3. 指针作为函数参数:")

	// 值传递
	increment := func(n int) {
		n++
		fmt.Printf("   函数内: %d\n", n)
	}
	y := 10
	increment(y)
	fmt.Printf("   函数外: %d (未改变)\n", y)

	// 指针传递
	incrementPtr := func(n *int) {
		*n++
		fmt.Printf("   函数内: %d\n", *n)
	}
	incrementPtr(&y)
	fmt.Printf("   函数外: %d (已改变)\n", y)

	// 4. 指针的零值
	fmt.Println("\n4. 指针的零值:")
	var ptr *int
	fmt.Printf("   零值指针: %v\n", ptr)
	if ptr == nil {
		fmt.Println("   指针为 nil")
	}

	// 5. new 函数
	fmt.Println("\n5. new 函数:")
	ptr2 := new(int)
	fmt.Printf("   new(int): %p, 值: %d\n", ptr2, *ptr2)
	*ptr2 = 100
	fmt.Printf("   赋值后: %p, 值: %d\n", ptr2, *ptr2)

	// 6. 结构体指针
	fmt.Println("\n6. 结构体指针:")
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Alice", Age: 25}
	personPtr := &person

	fmt.Printf("   person: %+v\n", person)
	fmt.Printf("   personPtr: %+v\n", personPtr)
	fmt.Printf("   (*personPtr).Name: %s\n", (*personPtr).Name)
	fmt.Printf("   personPtr.Name: %s (自动解引用)\n", personPtr.Name)

	// 7. 指针数组 vs 数组指针
	fmt.Println("\n7. 指针数组 vs 数组指针:")

	// 指针数组：数组的元素是指针
	a, b, c := 1, 2, 3
	ptrArray := [3]*int{&a, &b, &c}
	fmt.Printf("   指针数组: %v\n", ptrArray)
	fmt.Printf("   第一个元素指向的值: %d\n", *ptrArray[0])

	// 数组指针：指向数组的指针
	arr := [3]int{10, 20, 30}
	arrPtr := &arr
	fmt.Printf("   数组指针: %p\n", arrPtr)
	fmt.Printf("   数组指针指向的值: %v\n", *arrPtr)

	// 8. 注意事项
	fmt.Println("\n8. 注意事项:")
	fmt.Println("   - 不能对常量取地址")
	fmt.Println("   - 不能对表达式取地址")
	fmt.Println("   - 解引用 nil 指针会 panic")
	fmt.Println("   - Go 没有指针运算")
}

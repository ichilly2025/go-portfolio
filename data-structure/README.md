# Go 数据结构

Go 语言常用数据结构的示例和说明。

## 📁 目录结构

```
data-structure/
├── 01-array-slice/     # 数组和切片
├── 02-map/             # 映射（字典）
├── 03-string/          # 字符串
├── 04-struct/          # 结构体
├── 05-interface/       # 接口
├── 06-channel/         # 通道
├── 07-pointer/         # 指针
└── 08-function/        # 函数
```

## 🚀 快速开始

```bash
# 运行单个示例
cd data-structure/01-array-slice
go run main.go

# 运行所有示例
for dir in */; do
    echo "=== $dir ==="
    (cd "$dir" && go run main.go)
    echo ""
done
```

## 📚 数据结构说明

### 1. Array & Slice（数组和切片）

**Array（数组）**：
- 固定长度
- 值类型
- 长度是类型的一部分

**Slice（切片）**：
- 动态长度
- 引用类型
- 底层是数组

**示例**: `01-array-slice/main.go`

---

### 2. Map（映射）

**特点**：
- 键值对存储
- 无序
- 引用类型
- 不是并发安全的

**使用场景**：
- 字典查找
- 缓存
- 计数器

**示例**: `02-map/main.go`

---

### 3. String（字符串）

**特点**：
- 不可变
- UTF-8 编码
- 底层是 []byte

**常用操作**：
- 拼接、切片、遍历
- strings 包
- strconv 包

**示例**: `03-string/main.go`

---

### 4. Struct（结构体）

**特点**：
- 自定义类型
- 值类型
- 可以有方法

**使用场景**：
- 数据建模
- 面向对象编程
- 配置管理

**示例**: `04-struct/main.go`

---

### 5. Interface（接口）

**特点**：
- 定义行为
- 隐式实现
- 多态

**使用场景**：
- 抽象
- 依赖注入
- 插件系统

**示例**: `05-interface/main.go`


---

### 6. Channel（通道）

**特点**：
- 用于 goroutine 通信
- 类型安全
- 可以有缓冲

**使用场景**：
- 并发通信
- 同步
- 事件通知

**示例**: `06-channel/main.go`

---

### 7. Pointer（指针）

**特点**：
- 存储地址
- 可以修改原值
- 零值是 nil

**使用场景**：
- 避免拷贝大对象
- 修改函数参数
- 实现数据结构

**示例**: `07-pointer/main.go`

---

### 8. Function（函数）

**特点**：
- 一等公民
- 可以作为参数和返回值
- 支持闭包

**使用场景**：
- 回调函数
- 高阶函数
- 函数式编程

**示例**: `08-function/main.go`

---

## 🎯 数据结构对比

| 数据结构 | 类型 | 可变 | 有序 | 并发安全 | 使用场景 |
|---------|------|------|------|---------|---------|
| **Array** | 值 | 是 | 是 | 否 | 固定大小集合 |
| **Slice** | 引用 | 是 | 是 | 否 | 动态数组 |
| **Map** | 引用 | 是 | 否 | 否 | 键值对存储 |
| **String** | 值 | 否 | 是 | 是 | 文本处理 |
| **Struct** | 值 | 是 | - | 否 | 数据建模 |
| **Interface** | 引用 | - | - | - | 抽象和多态 |
| **Channel** | 引用 | - | 是 | 是 | 并发通信 |
| **Pointer** | 值 | - | - | 否 | 引用传递 |

---

## 💡 最佳实践

### 1. Slice 预分配

```
// ❌ 不预分配
s := []int{}

// ✅ 预分配
s := make([]int, 0, 100)
```

### 2. Map 预分配

```
// ❌ 不预分配
m := make(map[string]int)

// ✅ 预分配
m := make(map[string]int, 100)
```

### 3. 字符串拼接

```
// ❌ 使用 +
s := ""
for i := 0; i < 1000; i++ {
    s += "a"
}

// ✅ 使用 strings.Builder
var builder strings.Builder
for i := 0; i < 1000; i++ {
    builder.WriteString("a")
}
```

### 4. 小对象值传递

```
// ✅ 小对象值传递
func process(p Person) {}

// ❌ 小对象指针传递
func process(p *Person) {}
```

### 5. 大对象指针传递

```
// ❌ 大对象值传递
func process(data [1000000]byte) {}

// ✅ 大对象指针传递
func process(data *[1000000]byte) {}
```


---

## ⚠️ 常见陷阱

### 1. Slice 共享底层数组

```
原数组: [1, 2, 3, 4, 5]
子切片: [2, 3, 4]
修改子切片会影响原数组
```

### 2. Map 并发不安全

```
多个 goroutine 同时读写 map 会 panic
解决：使用 sync.Map 或加锁
```

### 3. 字符串不可变

```
不能直接修改字符串
需要转换为 []byte 或 []rune
```

### 4. Interface 的 nil

```
var i interface{} = nil  // nil 接口
var p *int = nil
var i interface{} = p    // 不是 nil 接口
```

### 5. 闭包捕获循环变量

```
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // 所有打印相同的值
    }()
}

// 正确做法
for i := 0; i < 5; i++ {
    i := i  // 创建新变量
    go func() {
        fmt.Println(i)
    }()
}
```

---

## 📊 性能对比

### Slice vs Array

```
Array:  固定大小，栈分配，快
Slice:  动态大小，可能堆分配，灵活
```

### Map vs Slice

```
Map:    O(1) 查找，无序
Slice:  O(n) 查找，有序
```

### String vs []byte

```
String: 不可变，适合只读
[]byte: 可变，适合修改
```

### Value vs Pointer

```
小对象（<64B）: 值传递更快
大对象（>64B）: 指针传递更快
```

---

## 🔍 选择指南

### 什么时候用 Array？

- 固定大小
- 需要值语义
- 性能关键

### 什么时候用 Slice？

- 动态大小
- 需要灵活性
- 大多数情况

### 什么时候用 Map？

- 键值对存储
- 快速查找
- 无序数据

### 什么时候用 Struct？

- 数据建模
- 组合多个字段
- 需要方法

### 什么时候用 Interface？

- 抽象行为
- 多态
- 依赖注入

### 什么时候用 Channel？

- Goroutine 通信
- 同步
- 事件驱动

### 什么时候用 Pointer？

- 避免拷贝大对象
- 需要修改原值
- 实现链表等数据结构

---

## 📖 学习资源

- [Go 语言规范](https://go.dev/ref/spec)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go 101](https://go101.org/)

---

## 🔗 相关示例

- [并发模型](../concurrency/) - 并发编程示例
- [内存管理](../memory/) - 内存管理示例
- [网络框架](../network/) - HTTP 框架对比

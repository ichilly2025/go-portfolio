# Go 并发模型示例

这个目录包含了 Go 语言 10 种常见并发模型的简单示例。

## 📁 目录结构

```
concurrency/
├── 01-goroutine-channel/    # Goroutine + Channel (CSP 模型)
├── 02-mutex/                 # Mutex (互斥锁)
├── 03-waitgroup/             # WaitGroup (等待组)
├── 04-context/               # Context (上下文控制)
├── 05-select/                # Select (多路复用)
├── 06-worker-pool/           # Worker Pool (工作池)
├── 07-pipeline/              # Pipeline (流水线)
├── 08-fan-out-fan-in/        # Fan-out/Fan-in (扇出/扇入)
├── 09-semaphore/             # Semaphore (信号量)
└── 10-once/                  # sync.Once (单次执行)
```

## 🚀 运行示例

### 运行单个示例

```bash
# 示例 1: Goroutine + Channel
cd concurrency/01-goroutine-channel
go run main.go

# 示例 2: Mutex
cd concurrency/02-mutex
go run main.go

# ... 以此类推
```

### 运行所有示例

```bash
# 在 concurrency 目录下
for dir in */; do
    echo "=== Running $dir ==="
    (cd "$dir" && go run main.go)
    echo ""
done
```

## 📚 并发模型说明

### 1. Goroutine + Channel (CSP 模型) ⭐⭐⭐

**核心思想**: "不要通过共享内存来通信，而要通过通信来共享内存"

**适用场景**:
- 生产者-消费者模式
- 任务分发
- 事件通知

**示例**: 生产者生产数据，消费者消费数据

---

### 2. Mutex (互斥锁)

**核心思想**: 共享内存 + 锁保护

**适用场景**:
- 简单的计数器
- 缓存更新
- 共享资源保护

**示例**: 多个 goroutine 并发增加计数器

---

### 3. WaitGroup (等待组)

**核心思想**: 等待多个 goroutine 完成

**适用场景**:
- 并发任务同步
- 批量处理
- 等待所有子任务完成

**示例**: 启动多个 worker，等待全部完成

---

### 4. Context (上下文控制)

**核心思想**: 超时控制和取消信号传播

**适用场景**:
- 超时控制
- 取消信号传播
- 请求级别的数据传递

**示例**: 3 秒后自动取消所有 worker

---

### 5. Select (多路复用)

**核心思想**: 监听多个 channel

**适用场景**:
- 超时控制
- 多个数据源
- 非阻塞操作

**示例**: 同时监听两个 channel 和超时

---

### 6. Worker Pool (工作池)

**核心思想**: 固定数量的 worker 处理任务队列

**适用场景**:
- 限制并发数
- 任务队列处理
- 资源池管理

**示例**: 3 个 worker 处理 9 个任务

---

### 7. Pipeline (流水线)

**核心思想**: 多阶段数据处理

**适用场景**:
- 数据处理流水线
- ETL（提取-转换-加载）
- 流式计算

**示例**: 生成数字 → 平方 → 加倍

---

### 8. Fan-out / Fan-in (扇出/扇入)

**核心思想**: 一个输入，多个处理器，一个输出

**适用场景**:
- 并行处理
- 结果聚合
- 负载均衡

**示例**: 多个 worker 并行计算平方，合并结果

---

### 9. Semaphore (信号量)

**核心思想**: 限制并发数量

**适用场景**:
- 限流
- 资源池
- 并发控制

**示例**: 最多 3 个任务同时执行

---

### 10. sync.Once (单次执行)

**核心思想**: 确保初始化代码只执行一次

**适用场景**:
- 单例模式
- 配置加载
- 资源初始化

**示例**: 多个 goroutine 获取单例，只初始化一次

---

## 🎯 选择建议

| 场景 | 推荐模型 |
|------|---------|
| 生产者-消费者 | Goroutine + Channel |
| 共享计数器 | Mutex |
| 等待多个任务 | WaitGroup |
| 超时控制 | Context |
| 多个数据源 | Select |
| 限制并发数 | Worker Pool / Semaphore |
| 数据流处理 | Pipeline |
| 并行计算 | Fan-out/Fan-in |
| 单例模式 | sync.Once |

## 📖 学习路径

### 初级（必须掌握）
1. Goroutine + Channel
2. WaitGroup
3. Mutex

### 中级（常用）
4. Context
5. Select
6. Worker Pool

### 高级（进阶）
7. Pipeline
8. Fan-out/Fan-in
9. Semaphore
10. sync.Once

## 🔗 相关资源

- [Go 并发编程实战](https://golang.org/doc/effective_go#concurrency)
- [Go 并发模式](https://go.dev/blog/pipelines)
- [Go 并发编程最佳实践](https://go.dev/blog/context)

## 💡 最佳实践

1. **优先使用 Channel**：Go 的哲学是通过通信来共享内存
2. **避免过度使用锁**：锁容易导致死锁和性能问题
3. **使用 Context 传递取消信号**：优雅地关闭 goroutine
4. **限制并发数**：使用 Worker Pool 或 Semaphore
5. **处理 goroutine 泄漏**：确保所有 goroutine 都能正常退出

## ⚠️ 常见陷阱

1. **Goroutine 泄漏**：忘记关闭 channel 或取消 context
2. **死锁**：循环等待或忘记释放锁
3. **竞态条件**：多个 goroutine 访问共享变量
4. **Channel 阻塞**：向已满的 channel 发送或从空 channel 接收
5. **过度并发**：创建太多 goroutine 导致性能下降

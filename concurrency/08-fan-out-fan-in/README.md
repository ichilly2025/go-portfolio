# Fan-out / Fan-in 模式

## 概念说明

### Fan-out（扇出）
**一个输入 → 多个处理器**

```
        输入 Channel
             ↓
    ┌────────┼────────┐
    ↓        ↓        ↓
 Worker1  Worker2  Worker3
    ↓        ↓        ↓
 Output1  Output2  Output3
```

**作用**：将任务分发给多个 worker 并行处理，提高吞吐量。

### Fan-in（扇入）
**多个输入 → 一个输出**

```
 Output1  Output2  Output3
    ↓        ↓        ↓
    └────────┼────────┘
             ↓
        合并 Channel
```

**作用**：将多个 worker 的结果合并到一个 channel。

### 完整流程

```
     输入
      ↓
  ┌───┴───┐
  ↓   ↓   ↓      Fan-out（分发）
  W1  W2  W3
  ↓   ↓   ↓
  └───┬───┘      Fan-in（合并）
      ↓
     输出
```

## 代码结构

### 1. fanOut 函数
```go
func fanOut(in <-chan int, numWorkers int) []<-chan int {
    channels := make([]<-chan int, numWorkers)
    for i := 0; i < numWorkers; i++ {
        channels[i] = square(i+1, in)  // 每个 worker 处理同一个输入
    }
    return channels
}
```

**关键点**：
- 多个 worker 从同一个输入 channel 读取数据
- 每个 worker 返回自己的输出 channel
- 返回所有 worker 的输出 channel 数组

### 2. fanIn 函数
```go
func fanIn(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup

    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for n := range c {
                out <- n  // 将所有输入合并到一个输出
            }
        }(ch)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}
```

**关键点**：
- 为每个输入 channel 启动一个 goroutine
- 所有 goroutine 将数据发送到同一个输出 channel
- 使用 WaitGroup 等待所有输入处理完成
- 最后关闭输出 channel

## 运行示例

```bash
go run main.go
```

输出：
```
=== Fan-out / Fan-in (扇出/扇入) ===
一个输入，多个处理器，一个输出

Fan-out: 将任务分发给 3 个 worker
Fan-in: 合并所有 worker 的结果

  Worker 1: 处理 1
  Worker 2: 处理 2
  Worker 3: 处理 3
结果:
1 4 9 16 25 36 49 64

说明: 多个 worker 并行处理，结果顺序可能不同
```

## 使用场景

### ✅ 适合的场景

1. **并行计算**
   ```
   图片处理：一张图片 → 多个 worker 处理不同区域 → 合并结果
   ```

2. **数据处理**
   ```
   日志分析：日志流 → 多个 worker 分析 → 合并统计结果
   ```

3. **爬虫**
   ```
   URL 队列 → 多个爬虫 worker → 合并抓取结果
   ```

4. **负载均衡**
   ```
   请求队列 → 多个处理器 → 合并响应
   ```

### ❌ 不适合的场景

1. **需要保证顺序**
   - Fan-out/Fan-in 不保证输出顺序
   - 如果需要顺序，使用 Pipeline

2. **任务之间有依赖**
   - 如果任务 B 依赖任务 A 的结果
   - 应该使用 Pipeline 而不是 Fan-out

3. **单个任务很快**
   - 如果任务执行时间很短（< 1ms）
   - 并行的开销可能大于收益

## 性能对比

### 串行处理
```
任务1 → 任务2 → 任务3 → 任务4 → 任务5 → 任务6
耗时: 6 × 100ms = 600ms
```

### Fan-out/Fan-in（3 个 worker）
```
任务1 → Worker1 ┐
任务2 → Worker2 ├→ 合并
任务3 → Worker3 ┘
任务4 → Worker1 ┐
任务5 → Worker2 ├→ 合并
任务6 → Worker3 ┘

耗时: 2 × 100ms = 200ms（快 3 倍）
```

## 注意事项

1. **Worker 数量**
   - 不是越多越好
   - 通常设置为 CPU 核心数
   - 过多会增加调度开销

2. **Channel 缓冲**
   - 输入 channel 可以有缓冲
   - 避免生产者阻塞

3. **错误处理**
   - Worker 出错时如何处理？
   - 是否需要取消其他 worker？
   - 考虑使用 Context

4. **资源清理**
   - 确保所有 channel 都被关闭
   - 避免 goroutine 泄漏

## 扩展阅读

- [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
- [Advanced Go Concurrency Patterns](https://go.dev/blog/io2013-talk-concurrency)

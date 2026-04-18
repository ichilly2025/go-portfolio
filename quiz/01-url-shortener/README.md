# 题目 1：URL 短链接服务

## 难度：中等

## 题目描述

实现一个 URL 短链接服务，类似于 bit.ly 或 tinyurl.com。

## 功能要求

### 1. 基本功能
- 将长 URL 转换为短 URL
- 通过短 URL 还原长 URL
- 统计每个短 URL 的访问次数

### 2. 技术要求
- 使用 Map 存储 URL 映射关系
- 使用 Struct 定义数据结构
- 使用 Mutex 保证并发安全
- 短 URL 长度为 6 个字符（使用 base62 编码）

### 3. 接口定义

```go
type URLShortener interface {
    // 生成短链接
    Shorten(longURL string) string
    
    // 还原长链接
    Expand(shortURL string) (string, bool)
    
    // 获取访问次数
    GetStats(shortURL string) int
}
```

## 输入输出示例

### 示例 1：基本功能
```
输入：
  Shorten("https://www.example.com/very/long/url/path")
输出：
  "abc123"

输入：
  Expand("abc123")
输出：
  "https://www.example.com/very/long/url/path", true
```

### 示例 2：访问统计
```
输入：
  Expand("abc123")  // 第 1 次访问
  Expand("abc123")  // 第 2 次访问
  GetStats("abc123")
输出：
  2
```

### 示例 3：不存在的短链接
```
输入：
  Expand("xyz999")
输出：
  "", false
```

## 测试用例

程序应该通过以下测试：

1. **基本功能测试**
   - 生成短链接
   - 还原长链接
   - 相同的长链接应该返回相同的短链接

2. **并发安全测试**
   - 多个 goroutine 同时生成短链接
   - 多个 goroutine 同时访问短链接
   - 访问计数应该准确

3. **边界测试**
   - 空字符串
   - 非常长的 URL
   - 不存在的短链接

## 提示

1. **短链接生成**：
   - 可以使用计数器 + base62 编码
   - base62 = [0-9a-zA-Z]

2. **数据结构**：
   ```go
   type URLData struct {
       LongURL  string
       ShortURL string
       Visits   int
   }
   ```

3. **并发安全**：
   - 使用 sync.Mutex 或 sync.RWMutex
   - 读多写少场景适合用 RWMutex

4. **性能优化**：
   - 使用 Map 预分配容量
   - 考虑使用双向映射（长→短，短→长）

## 评分标准

- 功能完整性（40%）
- 并发安全（30%）
- 代码质量（20%）
- 测试覆盖（10%）

## 扩展挑战（可选）

1. 添加过期时间功能
2. 添加自定义短链接功能
3. 添加 URL 验证功能
4. 实现持久化存储
5. 添加 HTTP 服务接口

## 时间限制

建议完成时间：30-45 分钟

# 编码完成
```
go run *.go
```
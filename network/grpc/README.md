# gRPC 简单示例

这是一个最简单的 gRPC 示例，实现了一个计算器服务，提供加法功能。

## 文件说明

- `calculator.proto` - Protocol Buffers 定义文件
- `server.go` - gRPC 服务端
- `client.go` - gRPC 客户端
- `pb/` - 自动生成的代码目录

## 快速开始

### 1. 生成 protobuf 代码

```bash
make proto
```

### 2. 启动服务端（终端 1）

```bash
make server
# 或
go run server.go
```

输出：
```
gRPC 服务器启动在 :50051
```

### 3. 运行客户端（终端 2）

```bash
make client
# 或
go run client.go
```

输出：
```
10 + 20 = 30
```

服务端会显示：
```
收到请求: 10 + 20 = 30
```

## 代码说明

### Protocol Buffers 定义

```protobuf
service Calculator {
    rpc Add(AddRequest) returns (AddResponse);
}
```

定义了一个 `Calculator` 服务，包含一个 `Add` 方法。

### 服务端实现

```go
type server struct {
    pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
    result := req.A + req.B
    return &pb.AddResponse{Result: result}, nil
}
```

实现 `Add` 方法，接收请求并返回结果。

### 客户端调用

```go
client := pb.NewCalculatorClient(conn)
response, err := client.Add(ctx, &pb.AddRequest{A: 10, B: 20})
```

创建客户端并调用 `Add` 方法。

## 常用命令

```bash
# 生成 protobuf 代码
make proto

# 启动服务端
make server

# 运行客户端
make client

# 清理生成的文件
make clean
```

## 下一步

可以尝试：
1. 添加更多方法（减法、乘法、除法）
2. 实现流式传输
3. 添加错误处理
4. 添加拦截器（中间件）

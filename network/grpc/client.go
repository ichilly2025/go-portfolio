package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ichilly2025/go-portfolio/network/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := pb.NewCalculatorClient(conn)

	// 调用 RPC 方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 测试加法
	response, err := client.Add(ctx, &pb.AddRequest{A: 10, B: 20})
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	log.Printf("10 + 20 = %d", response.Result)
}

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ichilly2025/go-portfolio/network/grpc/pb"

	"google.golang.org/grpc"
)

// server 实现 Calculator 服务
type server struct {
	pb.UnimplementedCalculatorServer
}

// Add 实现加法
func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	result := req.A + req.B
	log.Printf("收到请求: %d + %d = %d", req.A, req.B, result)
	return &pb.AddResponse{Result: result}, nil
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})

	log.Println("gRPC 服务器启动在 :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}

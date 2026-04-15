.PHONY: proto server client clean

# 生成 protobuf 代码
proto:
	@echo "生成 protobuf 代码..."
	@mkdir -p pb
	protoc --go_out=pb --go_opt=paths=source_relative \
	       --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	       calculator.proto
	@echo "✅ 生成完成！"

# 运行服务端
server:
	@echo "启动 gRPC 服务器..."
	go run server.go

# 运行客户端
client:
	@echo "运行 gRPC 客户端..."
	go run client.go

# 清理生成的文件
clean:
	@echo "清理生成的文件..."
	rm -rf pb/*.pb.go
	@echo "✅ 清理完成！"

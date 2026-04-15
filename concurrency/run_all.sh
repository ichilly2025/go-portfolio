#!/bin/bash

echo "=========================================="
echo "Go 并发模型示例 - 运行所有示例"
echo "=========================================="
echo ""

# 定义所有示例目录
examples=(
    "01-goroutine-channel"
    "02-mutex"
    "03-waitgroup"
    "04-context"
    "05-select"
    "06-worker-pool"
    "07-pipeline"
    "08-fan-out-fan-in"
    "09-semaphore"
    "10-once"
)

# 运行每个示例
for example in "${examples[@]}"; do
    echo "=========================================="
    echo "运行: $example"
    echo "=========================================="
    (cd "$example" && go run main.go)
    echo ""
    echo "按 Enter 继续下一个示例..."
    read
    echo ""
done

echo "=========================================="
echo "所有示例运行完成！"
echo "=========================================="

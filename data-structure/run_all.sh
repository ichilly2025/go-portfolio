#!/bin/bash

echo "=========================================="
echo "Go 数据结构示例 - 运行所有示例"
echo "=========================================="
echo ""

# 定义所有示例目录
examples=(
    "01-array-slice"
    "02-map"
    "03-string"
    "04-struct"
    "05-interface"
    "06-channel"
    "07-pointer"
    "08-function"
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

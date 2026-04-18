package main

import (
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Encode 将 string 编码为 Base62
func Encode(str string) string {
	n := stringToUint64(str)
	if n == 0 {
		return string(alphabet[0])
	}
	res := ""
	for n > 0 {
		res = string(alphabet[n%62]) + res
		n /= 62
	}
	return res
}

// Decode 将 Base62 字符串解码为 uint64
func Decode(s string) uint64 {
	var res uint64
	for _, char := range s {
		res = res*62 + uint64(strings.Index(alphabet, string(char)))
	}
	return res
}

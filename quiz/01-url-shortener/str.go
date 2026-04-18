package main

import (
    "crypto/md5"
    "encoding/binary"
)

func stringToUint64(s string) uint64 {
    hash := md5.Sum([]byte(s))
    return binary.BigEndian.Uint64(hash[:8])
}
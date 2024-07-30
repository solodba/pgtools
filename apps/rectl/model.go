package rectl

import (
	"fmt"
	"strconv"
)

// 多事务结构体
type Mxid struct {
	// 下一个多事务ID
	NextMxid string
	// 最旧多事务ID
	OldestMxid string
}

// 十六进制转换成十进制
func HexToDec(num string) uint64 {
	hexNum, _ := strconv.ParseUint(num, 16, 64)
	return hexNum
}

// 十进制转换成十六进制
func DecToHexWal(num uint64) string {
	dec := fmt.Sprintf("%x", num)
	if len(dec) < 2 {
		return "0000000" + dec
	} else {
		return "000000" + dec
	}
}

func DecToHexMxid(num uint64) string {
	return fmt.Sprintf("%x", num)
}

// Mxid构造函数
func NewMxid() *Mxid {
	return &Mxid{}
}

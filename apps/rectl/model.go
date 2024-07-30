package rectl

// 多事务结构体
type Mxid struct {
	// 下一个多事务ID
	NextMxid string
	// 最旧多事务ID
	OldestMxid string
}

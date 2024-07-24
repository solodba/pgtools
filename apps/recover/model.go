package recover

// TxInfo结构体
type TxInfo struct {
	// 事务id
	TxId string `json:"txid"`
}

// TxInfo结构体构造函数
func NewTxInfo() *TxInfo {
	return &TxInfo{}
}

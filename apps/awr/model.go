package awr

// 系统信息结构体
type SystemInfo struct {
	// 主机名
	Hostname string `json:"hostname"`
	// 运行时间
	RunTime string `json:"run_time"`
	// CPU信息
	Cpu string `json:"cpu"`
	// 平均负载
	LoadAvg string `json:"load_avg"`
	// 内存信息
	Memory string `json:"memory"`
	// 交换分区信息
	Swap string `json:"swap"`
}

// 系统信息结构体构造函数
func NewSystemInfo() *SystemInfo {
	return &SystemInfo{}
}

// AwrData结构体
type AwrData struct {
	SystemInfo *SystemInfo `json:"system_info"`
}

// AwrData结构体初始化函数
func NewAwrData() *AwrData {
	return &AwrData{
		SystemInfo: NewSystemInfo(),
	}
}
